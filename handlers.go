package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)
// GetRouteByID retrieves a route by its ID from the database
func GetRouteByIDHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Parse the incoming JSON request body
        var id struct {
            ID int `json:"id"`
        }

        if err := c.ShouldBindJSON(&id); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        var route Route
        var geomJSON string

        // Query the route by ID
        query := "SELECT id, ST_AsGeoJSON(geometry) FROM routes WHERE id = $1"
        err := db.QueryRow(query, id.ID).Scan(&route.ID, &geomJSON)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching route: %v", err)})
            return
        }

        // Unmarshal the GeoJSON string into the Geometry struct
        if err := json.Unmarshal([]byte(geomJSON), &route.Geometry); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error parsing geometry: %v", err)})
            return
        }

        // Send response as JSON
        c.JSON(http.StatusOK, route)
    }
}


func AddRouteHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Parse the incoming JSON request
        var addresses struct {
            From string `json:"from"`
            To   string `json:"to"`
        }

        if err := c.ShouldBindJSON(&addresses); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        fromPoint := getLocation(addresses.From)
        toPoint := getLocation(addresses.To)

        // Add the route to the database
        if err := addRoute(db, fromPoint, toPoint); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error adding route: %v", err)})
            return
        }

        // Send success response
        c.JSON(http.StatusCreated, gin.H{"message": "Route added successfully"})
    }
}


func GetAllRoutesHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get all routes from the database
        allRoutes, err := getAllRoutes(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching routes: %v", err)})
            return
        }

        // Return JSON response with all routes
        response := RoutesResponse{
            Routes: allRoutes,
        }

        c.JSON(http.StatusOK, response)
    }
}


func DeleteRouteHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Delete route from the database
        onRoute, err := deleteRoute(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting route: %v", err)})
            return
        }

        // Return JSON response indicating if the address was on the route
        response := AddressCheckResponse{
            OnRoute: onRoute,
        }

        c.JSON(http.StatusOK, response)
    }
}


// Route handler for /api/on-route using PostGIS
func OnRouteHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Parse the incoming JSON request
        var request AddressCheckRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            // Return a bad request error if the JSON is invalid
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        // Check if the address is on any route using PostGIS
        routeResult, err := isOnRoute(db, getLocation(request.From), getLocation(request.To))
        if err != nil {
            // Return an internal server error if something went wrong
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error checking route: %v", err)})
            return
        }

        // Prepare the response
        response := AddressCheckResponse{
            OnRoute: routeResult.OnRoute,
            RouteID: routeResult.RouteID,
        }

        // Send the response as JSON
        c.JSON(http.StatusOK, response)
    }
}

