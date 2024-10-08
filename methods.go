package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next(w, r)
	}
}

func initDatabase() *sql.DB {

	connStr := "user=myuser dbname=mydb password=mypassword host=postgis-db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	if err != nil {
		log.Fatalf("Failed to create PostGIS extension: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS routes (
			id SERIAL PRIMARY KEY,
			geometry GEOMETRY(LineString, 4326)
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	return db
}

func getAllRoutes(db *sql.DB) ([]Route, error) {
	const query = `
		SELECT id, ST_AsGeoJSON(geometry) 
		FROM routes;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying routes: %v", err)
	}
	defer rows.Close()

	var routes []Route

	for rows.Next() {
		var route Route
		var geomJSON string

		if err := rows.Scan(&route.ID, &geomJSON); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		if err := json.Unmarshal([]byte(geomJSON), &route.Geometry); err != nil {
			return nil, fmt.Errorf("error unmarshaling GeoJSON: %v", err)
		}

		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %v", err)
	}

	return routes, nil
}

func deleteRoute(db *sql.DB) (bool, error) {
	const query = `
		DELETE FROM routes;
	`
	result, err := db.Exec(query)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func isOnRoute(db *sql.DB, addressFrom Point, addressTo Point) (RouteResult, error) {
	fmt.Print(addressFrom.Latitude, addressFrom.Longitude)
    const query = `
        SELECT r.id
        FROM routes r
        WHERE (
            ST_DWithin(
                ST_Transform(ST_SetSRID(ST_MakePoint($1, $2), 4326), 3857),  -- Transform from WGS 84 to Web Mercator
                ST_Transform(r.geometry, 3857),
                1500
            )
            AND
            ST_DWithin(
                ST_Transform(ST_SetSRID(ST_MakePoint($3, $4), 4326), 3857),
                ST_Transform(r.geometry, 3857),
                1500
            )
        )
        LIMIT 1;
    `

    var routeID int
    err := db.QueryRow(query, addressFrom.Longitude, addressFrom.Latitude, addressTo.Longitude, addressTo.Latitude).Scan(&routeID)

    if err != nil {
        if err == sql.ErrNoRows {
            return RouteResult{RouteID: 0, OnRoute: false}, nil
        }
        return RouteResult{}, err
    }

    return RouteResult{RouteID: routeID, OnRoute: true}, nil
}

func getLocation(address string) Point {
    baseURL := "https://nominatim.openstreetmap.org/search"

    params := url.Values{}
    params.Add("q", address)
    params.Add("format", "json")

    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    resp, err := http.Get(fullURL)
    if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response: %v", err)
    }

    // Unmarshal the JSON response into a slice of Location
    var locations []Location
    if err := json.Unmarshal(body, &locations); err != nil {
        log.Fatalf("Error unmarshaling JSON: %v", err)
    }

    // Print the location data
	lon, err := strconv.ParseFloat(locations[0].Lon, 64)
	lat, err := strconv.ParseFloat(locations[0].Lat, 64)
	point := Point{lon, lat}
	return point
}

// addRoute retrieves a route from the OSRM and adds it to the database.
func addRoute(db *sql.DB, from Point, to Point) error {
	// Get route from OSRM
	osrmResponse, err := getRouteFromOSRM(from, to)
	if err != nil {
		return fmt.Errorf("failed to get route from OSRM: %v", err)
	}

	// Prepare the SQL insert statement
	stmt, err := db.Prepare("INSERT INTO routes(geometry) VALUES (ST_GeomFromGeoJSON($1))")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %v", err)
	}
	defer stmt.Close()

	// Insert each route's geometry into the database
	for _, route := range osrmResponse.Route {
		geoJSON, err := json.Marshal(route.Geometry)
		if err != nil {
			return fmt.Errorf("failed to marshal geometry to GeoJSON: %v", err)
		}

		_, err = stmt.Exec(string(geoJSON))
		if err != nil {
			return fmt.Errorf("failed to insert route geometry: %v", err)
		}
	}

	return nil
}

// Function to get a route by ID from the database
func getRouteByID(db *sql.DB, id string) (Route, error) {
    var route Route
    var geomJSON string

    // Query the route by ID
    err := db.QueryRow("SELECT ST_AsGeoJSON(geometry) FROM routes WHERE id = $1", id).Scan(&geomJSON)
    if err != nil {
        return route, err
    }

    // Unmarshal the GeoJSON string into the Geometry struct
    if err := json.Unmarshal([]byte(geomJSON), &route.Geometry); err != nil {
        return route, err
    }

    return route, nil
}


// findMatchingRoutes finds routes that match the given pickup and dropoff points.
func findMatchingRoutes(db *sql.DB, pickupLat, pickupLng, dropoffLat, dropoffLng float64, distanceThreshold float64) ([]Route, error) {
	query := `
        SELECT id, ST_AsGeoJSON(geometry) 
        FROM routes 
        WHERE ST_DWithin(geometry::geography, 
                         ST_MakePoint($1, $2)::geography, $3)
          AND ST_DWithin(geometry::geography, 
                         ST_MakePoint($4, $5)::geography, $3);`

	rows, err := db.Query(query, pickupLng, pickupLat, distanceThreshold, dropoffLng, dropoffLat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matchedRoutes []Route
	for rows.Next() {
		var id int
		var geomJSON string
		if err := rows.Scan(&id, &geomJSON); err != nil {
			return nil, err
		}

		// Unmarshal the GeoJSON string into the Geometry struct
		var geometry Geometry
		if err := json.Unmarshal([]byte(geomJSON), &geometry); err != nil {
			return nil, err
		}

		// Append the route to the matchedRoutes slice
		matchedRoutes = append(matchedRoutes, Route{Geometry: geometry})
	}

	return matchedRoutes, rows.Err()
}


// getRouteFromOSRM fetches the route from the OSRM API.
func getRouteFromOSRM(pointA Point, pointB Point) (OsrmResponse, error) {
	url := fmt.Sprintf("http://osrm-backend:5000/route/v1/driving/%.4f,%.4f;%.4f,%.4f?overview=full&geometries=geojson", pointA.Longitude, pointA.Latitude, pointB.Longitude, pointB.Latitude)
	blank := OsrmResponse{}

	resp, err := http.Get(url)
	if err != nil {
		return blank, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return blank, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return blank, err
	}

	var osrmResp OsrmResponse
	if err := json.Unmarshal(body, &osrmResp); err != nil {
		return blank, err
	}

	// Return the geometry of the first route
	if len(osrmResp.Route) > 0 {
		return osrmResp, nil
	}

	return blank, fmt.Errorf("no routes found")
}

