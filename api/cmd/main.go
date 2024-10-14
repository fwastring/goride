package main

import (
	"goride/internal/handlers"
	"goride/internal/store"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// "github.com/paulsmith/gogeos/geos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Environment = "development"

func main() {
	// Initialize zap logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := gin.New()

	// Load the config
	// cfg := config.MustLoadConfig()
	dbURL := "postgres://myuser:mypassword@localhost:5432/mydb"

	// Initialize database
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	// point, err := geos.NewPoint(geos.NewCoord(40.7128, -74.0060)) // Example coordinates for New York
	db.AutoMigrate(&store.Route{})


    var route store.Route
    err = db.Where("id = ?", 3).First(&route).Error

    if err != nil {
		logger.Error("Failed to insert to database: %v", err)
    }

	// Set up middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Define the routes
	r.GET("/", func(c *gin.Context) {
		handlers.NewAutoCompleteHandler(handlers.AutoCompleteHandlerParams{
			Logger: *logger,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	r.GET("/route", func(c *gin.Context) {
		handlers.NewGetRouteByIDHandler(handlers.GetRouteByIDHandlerParams{
			Logger: *logger,
			Database: *db,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	r.POST("/route", func(c *gin.Context) {
		handlers.NewCreateRouteHandler(handlers.CreateRouteHandlerParams{
			Logger: *logger,
			Database: *db,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	r.Run()

}


