package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"goride/internal/store"
	"goride/internal/store/dbstore"
	"goride/internal/store/types"

	"github.com/gin-gonic/gin"
	"github.com/paulsmith/gogeos/geos"
	"gorm.io/gorm"
	"github.com/twpayne/go-geom/encoding/wkt"
	gogeom "github.com/twpayne/go-geom"
)

type CreateRouteHandler struct {
	logger         slog.Logger
	database 	   gorm.DB
}

type CreateRouteHandlerParams struct {
	Logger         slog.Logger
	Database 	   gorm.DB
}


func NewCreateRouteHandler(params CreateRouteHandlerParams) *CreateRouteHandler {
	return &CreateRouteHandler{
		logger: params.Logger,	
		database: params.Database,
	}
}

func (h *CreateRouteHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var addresses struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	if err := c.ShouldBindJSON(&addresses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	points := []store.Point{getLocation(addresses.From, h.logger), getLocation(addresses.To, h.logger)}

	response, err := getRouteFromOSRM(points, h.logger)

	if err != nil {
        h.logger.Error("Error reading response: %v", err)
	}
	// Add the route to the database

	routeStore := dbstore.NewRouteStore(dbstore.NewRouteStoreParams{DB: &h.database})

	// Convert the Geometry struct to WKT
	geometry := *response.Geometry
    wkt, err := geometryToWKT(geometry)

    if err != nil {
        h.logger.Error("Error converting to WKT: %v", err)
    }

    // Create a *geos.Geometry object from the WKT string
    geosGeom, err := geos.FromWKT(wkt)
    if err != nil {
        h.logger.Error("Error converting from WKT: %v", err)
    }

    // Wrap the *geos.Geometry in your types.Geometry4326 struct
    geometry4326 := types.Geometry4326{Geometry: geosGeom}

	err = routeStore.CreateRoute(addresses.From, addresses.To, geometry4326)
	if err != nil {
		h.logger.Error("Error error adding route to db: %v", err)
	}

    // Send response as JSON
    c.JSON(http.StatusOK, "Route created")
}

// Function to convert a Geometry struct to WKT
func geometryToWKT(geom geos.Geometry) (string, error) {
    // Build WKT string for a LineString
    wkt := "LINESTRING("
	coords, err := geom.Coords()
	if err != nil {

	}
    for i, coord := range coords {
        if i > 0 {
            wkt += ", "
        }
        wkt += fmt.Sprintf("%f %f", coord.X, coord.Y)
    }
    wkt += ")"
    return wkt, nil
}


func getRouteFromOSRM(points []store.Point, logger slog.Logger) (types.Geometry4326, error) {
	blank := types.Geometry4326{}
	if len(points) < 2 {
        return blank, fmt.Errorf("at least two points are required")
    }

    // Initialize an empty slice to hold coordinate pairs
    var coordinates []string

    // Iterate over the points and create "lon,lat" strings
    for _, point := range points {
        coord := fmt.Sprintf("%.4f,%.4f", point.Longitude, point.Latitude)
        coordinates = append(coordinates, coord)
    }

    // Join the coordinates slice into a single string, with ";" as separator
    coordinatesStr := strings.Join(coordinates, ";")

    // Build the final URL
    url := fmt.Sprintf("http://localhost:5000/route/v1/driving/%s?overview=full&geometries=geojson", coordinatesStr)

	// Query the OSRM API
	resp, err := http.Get(url)
	if err != nil {
		return blank, err
	}
	defer resp.Body.Close()

	// Check if we received a successful status code
	if resp.StatusCode != http.StatusOK {
		return blank, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return blank, err
	}

	// Parse the OSRM response to extract GeoJSON geometry
	var osrmResp struct {
		Routes []struct {
			Geometry struct {
				Coordinates [][]float64 `json:"coordinates"`
				Type        string      `json:"type"`
			} `json:"geometry"`
		} `json:"routes"`
	}

	if err := json.Unmarshal(body, &osrmResp); err != nil {
		return blank, err
	}

	// Check if any routes were returned
	if len(osrmResp.Routes) == 0 {
		return blank, fmt.Errorf("no routes found in OSRM response")
	}

	// Extract the geometry of the first route (which is a LineString)
	routeGeometry := osrmResp.Routes[0].Geometry

	// Create a go-geom LineString and set its coordinates
	lineString := gogeom.NewLineString(gogeom.XY)
	var coords []gogeom.Coord
	for _, coord := range routeGeometry.Coordinates {
		coords = append(coords, gogeom.Coord{coord[0], coord[1]})
	}

	// Set the coordinates of the LineString
	if lineString, err = lineString.SetCoords(coords); err != nil {
		return blank, fmt.Errorf("error setting LineString coordinates: %w", err)
	}

	// Marshal the go-geom LineString to WKT (Well-Known Text)
	wktStr, err := wkt.Marshal(lineString)
	if err != nil {
		return blank, fmt.Errorf("error marshaling geometry to WKT: %w", err)
	}

	// Convert WKT to geos.Geometry
	geom, err := geos.FromWKT(wktStr)
	if err != nil {
		return blank, fmt.Errorf("error converting WKT to geos.Geometry: %w", err)
	}

	// Return the Geometry4326 object
	return types.Geometry4326{Geometry: geom}, nil
}



func getLocation(address string, logger slog.Logger) store.Point {
    baseURL := "https://nominatim.openstreetmap.org/search"

    params := url.Values{}
    params.Add("q", address)
    params.Add("format", "json")

    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    resp, err := http.Get(fullURL)
    if err != nil {
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
		logger.Error("Error reading body: %s \n", err)
    }

    // Unmarshal the JSON response into a slice of Location
    var locations []store.Location
    if err := json.Unmarshal(body, &locations); err != nil {
		logger.Error("Error reading body: %s \n", err)
    }

    // Print the location data
	lon, err := strconv.ParseFloat(locations[0].Lon, 64)
	lat, err := strconv.ParseFloat(locations[0].Lat, 64)
	point := store.Point{Longitude: lon, Latitude: lat}
	return point
}

