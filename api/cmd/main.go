package main

import (
	"database/sql"
	"goride/handlers"
	"goride/store/dbstore"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var Environment = "development"

func main() {
	 
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := gin.New()

	connStr := "user=myuser dbname=mydb password=mypassword host=localhost port=5432 sslmode=disable"

	 
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to connect to database", err)
	}

	db.Exec(`
		CREATE EXTENSION IF NOT EXISTS postgis;
		CREATE EXTENSION IF NOT EXISTS citext;
	`)
	 
    createTableSQL, err := os.ReadFile("./scripts/schema.sql")
	if err != nil {
		logger.Error("Failed to read SQL schema", err)
	}

     
    _, err = db.Exec(string(createTableSQL))
    if err != nil {
        logger.Error("Failed to create table", err)
    }

	routeStore := dbstore.NewRouteStore(dbstore.NewRouteStoreParams{
		DB: db,
	})

	userStore := dbstore.NewUserStore(dbstore.NewUserStoreParams{
		DB: db,
	})

	sessionStore := dbstore.NewSessionStore(dbstore.NewSessionStoreParams{
		DB: db,
	})


    if err != nil {
		logger.Error("Failed to insert to database", err)
    }

	 
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	 
	r.GET("/", func(c *gin.Context) {
		handlers.NewAutoCompleteHandler(handlers.AutoCompleteHandlerParams{
			Logger: *logger,
		}).ServeHTTP(c, c.Writer, c.Request)
	})


	userRoute := r.Group("/user")

	userRoute.GET("", func(c *gin.Context) {
		handlers.NewGetUserHandler(handlers.GetUserHandlerParams{
			Logger: *logger,
			Database: db,
			UserStore: *userStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	userRoute.POST("/login", func(c *gin.Context) {
		handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
			Logger: *logger,
			SessionStore: *sessionStore,
			UserStore: *userStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	tripRoute := r.Group("/route")

	tripRoute.GET("/all", func(c *gin.Context) {
		handlers.NewGetAllRoutesHandler(handlers.GetAllRoutesHandlerParams{
			Logger: *logger,
			Database: db,
			RouteStore: *routeStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	tripRoute.GET("/:id", func(c *gin.Context) {
		handlers.NewGetRouteByIDHandler(handlers.GetRouteByIDHandlerParams{
			Logger: *logger,
			Database: db,
			RouteStore: *routeStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	tripRoute.POST("/create", func(c *gin.Context) {
		handlers.NewCreateRouteHandler(handlers.CreateRouteHandlerParams{
			Logger: *logger,
			Database: db,
			RouteStore: *routeStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	tripRoute.POST("/search", func(c *gin.Context) {
		handlers.NewSearchRouteHandler(handlers.SearchRouteHandlerParams{
			Logger: *logger,
			Database: db,
			RouteStore: *routeStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	tripRoute.POST("/join", func(c *gin.Context) {
		handlers.NewJoinRouterHandler(handlers.JoinRouterHandlerParams{
			Logger: *logger,
			Database: db,
			RouteStore: *routeStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	tripRoute.POST("/delete", func(c *gin.Context) {
		handlers.NewDeleteRouteHandler(handlers.DeleteRouteHandlerParams{
			Logger: *logger,
			Database: db,
			RouteStore: *routeStore,
		}).ServeHTTP(c, c.Writer, c.Request)
	})

	r.Run()
}
