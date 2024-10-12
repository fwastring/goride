package main

import (
	"fmt"
	// "goride/internal/config"
	"goride/internal/handlers"
	"log/slog"
	"os"

	database "goride/internal/store/db"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

var Environment = "development"

func main() {
	// Initialize zap logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := gin.New()

	// Load the config
	// cfg := config.MustLoadConfig()

	// Initialize database
	db := database.InitDatabase()
	fmt.Print(db)

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


