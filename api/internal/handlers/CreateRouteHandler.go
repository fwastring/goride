package handlers

import (
	"encoding/json"
	"fmt"
	"goride/internal/store/dbstore"
	"goride/internal/types"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	// "github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
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

	points := []types.Point{getLocation(addresses.From, h.logger), getLocation(addresses.To, h.logger)}

	response, err := getRouteFromOSRM(points, h.logger)
	if err != nil {
        h.logger.Error("Error reading response: %v", err)
	}
	// Add the route to the database

	routeStore := dbstore.NewRouteStore(dbstore.NewRouteStoreParams{DB: &h.database})
	fmt.Print(routeStore)

	for _, route := range response.Route {
		h.logger.Info(route.Type)
	}

	// err = routeStore.CreateRoute(response.Route[0].Geometry)
	// if err != nil {
 //        h.logger.Error("Error creating route: %v", err)
	// }
	//
    // Send response as JSON
    c.JSON(http.StatusOK, "Route created")
}

func getRouteFromOSRM(points []types.Point, logger slog.Logger) (types.OsrmResponse, error) {
	blank := types.OsrmResponse{}
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

    logger.Info("Generated OSRM request URL", "url", url)

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

	var osrmResp types.OsrmResponse
	if err := json.Unmarshal(body, &osrmResp); err != nil {
		return blank, err
	}

	// Return the geometry of the first route
	if len(osrmResp.Route) > 0 {
		return osrmResp, nil
	}

	return blank, fmt.Errorf("no routes found")
}

func getLocation(address string, logger slog.Logger) types.Point {
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
    var locations []types.Location
    if err := json.Unmarshal(body, &locations); err != nil {
		logger.Error("Error reading body: %s \n", err)
    }

    // Print the location data
	lon, err := strconv.ParseFloat(locations[0].Lon, 64)
	lat, err := strconv.ParseFloat(locations[0].Lat, 64)
	point := types.Point{Longitude: lon, Latitude: lat}
	return point
}

