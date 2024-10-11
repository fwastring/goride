package main

import (
	// "log"
	"context"
	"errors"

	"time"
	"goride/internal/config"
	"goride/internal/hash/passwordhash"
	"goride/internal/handlers"
	"log/slog"
	"os"
	"net/http"
	"os/signal"
	"syscall"



	database "goride/internal/store/db"
	"goride/internal/store/dbstore"


	_ "github.com/lib/pq"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Environment = "development"



func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()


	cfg := config.MustLoadConfig()

	db := database.MustOpen(cfg.DatabaseName)
	passwordhash := passwordhash.NewHPasswordHash()

	userStore := dbstore.NewUserStore(
		dbstore.NewUserStoreParams{
			DB:           db,
			PasswordHash: passwordhash,
		},
	)

	sessionStore := dbstore.NewSessionStore(
		dbstore.NewSessionStoreParams{
			DB: db,
		},
	)



	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
		)

		// r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)

		r.Post("/", handlers.NewAutoCompleteHandler(handlers.AutoCompleteHandlerParams{
			Logger: *logger,
		},
		).ServeHTTP)

		r.Post("/register", handlers.NewPostRegisterHandler(handlers.PostRegisterHandlerParams{
			UserStore: userStore,
		}).ServeHTTP)

		r.Post("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
			UserStore:         userStore,
			SessionStore:      sessionStore,
			PasswordHash:      passwordhash,
			SessionCookieName: cfg.SessionCookieName,
		}).ServeHTTP)
		// 	SessionCookieName: cfg.SessionCookieName,
		// }).ServeHTTP)
	})

		killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", cfg.Port), slog.String("env", Environment))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")



	// router.POST("/api/add-route", AddRouteHandler(db))
 //    router.POST("/api/on-route", OnRouteHandler(db))
 //    router.GET("/api/routes", GetAllRoutesHandler(db))
 //    router.DELETE("/api/routes/delete", DeleteRouteHandler(db))
 //    router.POST("/api/route", GetRouteByIDHandler(db)) // Use param for ID
	//
 //    log.Println("Server is running on port 8080...")
 //    if err := router.Run(":8080"); err != nil {
 //        log.Fatal(err)
 //    }
}

type OsrmResponse struct {
	Route []Route `json:"routes"`
}

type Route struct {
	ID string `json:"id"`
	Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
	Type        string      `json:"type"`
}

type Point struct {
	Longitude float64
	Latitude  float64
}

type Location struct {
    DisplayName string  `json:"display_name"`
    Lat         string  `json:"lat"`
    Lon         string  `json:"lon"`
}


type RouteResult struct {
    RouteID int  `json:"route_id"`
    OnRoute bool `json:"on_route"`
}

type AddressCheckRequest struct {
	From string `json:"from"`
	To string `json:"to"`
}

type AddressCheckResponse struct {
	OnRoute bool `json:"on_route"`
	RouteID int `json:"id"`
}

type RoutesResponse struct {
	Routes []Route `json:"routes"`
}
