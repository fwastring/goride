package dbstore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goride/store"
	"goride/store/types"

	_ "github.com/lib/pq"  
)

type RouteStore struct {
    db *sql.DB  
}

type NewRouteStoreParams struct {
    DB *sql.DB  
}

func NewRouteStore(params NewRouteStoreParams) *RouteStore {
    return &RouteStore{
        db: params.DB,
    }
}

func (s *RouteStore) CreateRoute(addressFrom string, addressTo string, geometry store.Geometry) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO routes(start_address, end_address, geometry) VALUES ($1, $2, ST_GeomFromGeoJSON($3))")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	 
	geoJSON, err := json.Marshal(geometry)
	if err != nil {
		return 0, err
	}


	_, err = stmt.Exec(addressFrom,addressTo, string(geoJSON))
	if err != nil {
		return 0, err
	}
	if err != nil {
		return 0, err
	}

    query := `
        SELECT LAST_INSERT_ID()
    `

	var id int64

	row := s.db.QueryRow(query, id)
 
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

    return id, nil
}

 
func (s *RouteStore) UpdateRoute(id uint, geometry store.Geometry) error {
    query := `
        UPDATE routes
        SET geometry = ST_GeomFromGeoJSON($1)
        WHERE route_id = $2
    `
	 
	geoJSON, err := json.Marshal(geometry)
	if err != nil {
		return fmt.Errorf("failed to marshal geometry to GeoJSON", err)
	}
    _, err = s.db.Exec(query, string(geoJSON), id)
    if err != nil {
        return fmt.Errorf("failed to update route: %w", err)
    }
    return nil
}

 
func (s *RouteStore) GetRoute(id uint) (store.Route, error) {
    var route store.Route
    query := `
        SELECT route_id, start_address, end_address, ST_ASGeoJSON(geometry)
        FROM routes
        WHERE route_id = $1
    `
    row := s.db.QueryRow(query, id)
	var geomJSON string
 
	if err := row.Scan(&route.ID,&route.StartAddress,&route.EndAddress, &geomJSON); err != nil {
		return store.Route{}, fmt.Errorf("error scanning row", err)
	}

	if err := json.Unmarshal([]byte(geomJSON), &route.Geometry); err != nil {
		return store.Route{}, fmt.Errorf("error unmarshaling GeoJSON", err)
	}
     
     
     
     
     
     
    return route, nil
}

 
func (s *RouteStore) GetAllRoutes() ([]store.Route, error) {
	const query = `
        SELECT route_id, start_address, end_address, ST_ASGeoJSON(geometry)
		FROM routes;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying routes", err)
	}
	defer rows.Close()

	var routes []store.Route

	for rows.Next() {
		var route store.Route
		var geomJSON string

		if err := rows.Scan(&route.ID,&route.StartAddress,&route.EndAddress, &geomJSON); err != nil {
			return nil, fmt.Errorf("error scanning row", err)
		}

		if err := json.Unmarshal([]byte(geomJSON), &route.Geometry); err != nil {
			return nil, fmt.Errorf("error unmarshaling GeoJSON", err)
		}

		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows", err)
	}

	return routes, nil
}

 
func (s *RouteStore) DeleteRoute(id uint) error {
    query := `
        DELETE FROM routes
        WHERE route_id = $1
    `
    _, err := s.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete route: %w", err)
    }
    return nil
}

 
func (s *RouteStore) GetRoutesOnPoints(pointFrom types.Point, pointTo types.Point, threshold int) ([]int, error) {
    query := `
        SELECT route_id
        FROM routes
        WHERE (
            ST_DWithin(
                ST_Transform(ST_SetSRID(ST_MakePoint($1, $2), 4326), 3857),
                ST_Transform(geometry, 3857),
				$5
            )
            AND
            ST_DWithin(
                ST_Transform(ST_SetSRID(ST_MakePoint($3, $4), 4326), 3857),
                ST_Transform(geometry, 3857),
				$5
            )
        )
    `

    rows, err := s.db.Query(query, pointFrom.Latitude, pointFrom.Longitude, pointTo.Latitude, pointTo.Longitude, threshold)
    if err != nil {
        return nil, fmt.Errorf("failed to get routes on points: %w", err)
    }
    defer rows.Close()

    var routeIDs []int
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return nil, fmt.Errorf("failed to scan route ID: %w", err)
        }
        routeIDs = append(routeIDs, id)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("failed during rows iteration: %w", err)
    }

    return routeIDs, nil
}

