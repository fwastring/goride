package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)


func main() {
	router := gin.Default()

    // Enable CORS with default options
    router.Use(cors.Default())
	db := initDatabase()
	defer db.Close()

	router.POST("/api/add-route", AddRouteHandler(db))
    router.POST("/api/on-route", OnRouteHandler(db))
    router.GET("/api/routes", GetAllRoutesHandler(db))
    router.DELETE("/api/routes/delete", DeleteRouteHandler(db))
    router.POST("/api/route", GetRouteByIDHandler(db)) // Use param for ID

	// Start the server
    log.Println("Server is running on port 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}

type OsrmResponse struct {
	Route []Route `json:"routes"`
}

// Route represents a route with geometry.
type Route struct {
	ID string `json:"id"`
	Geometry `json:"geometry"`
}

// Geometry represents the geometric structure.
type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
	Type        string      `json:"type"`
}

// Point represents a geographical point with longitude and latitude.
type Point struct {
	Longitude float64
	Latitude  float64
}

// Structure to unmarshal the response JSON
type Location struct {
    DisplayName string  `json:"display_name"`
    Lat         string  `json:"lat"`
    Lon         string  `json:"lon"`
}


type RouteResult struct {
    RouteID int  `json:"route_id"`
    OnRoute bool `json:"on_route"`
}


// Struct for receiving the input JSON
type AddressCheckRequest struct {
	From string `json:"from"`
	To string `json:"to"`
}

// Struct for response
type AddressCheckResponse struct {
	OnRoute bool `json:"on_route"`
	RouteID int `json:"id"`
}

// Struct for response
type RoutesResponse struct {
	Routes []Route `json:"routes"`
}
