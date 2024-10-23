package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"

	 

	"github.com/spatial-go/geoos/geoencoding/geojson"
)

type AddRouteHandler struct {
	logger         slog.Logger
	database 	   *sql.DB
}

type AddRouteHandlerParams  struct {
	Logger         slog.Logger
	Database 	   *sql.DB
}


func NewAddRouteHandler(params AddRouteHandlerParams) *AddRouteHandler {
	return &AddRouteHandler{
		logger: params.Logger,	
		database: params.Database,
	}
}


func (h *AddRouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	baseURL := "http: 


    params := url.Values{}
    params.Add("boundary.circle.lat", "64")
    params.Add("boundary.circle.lon", "17")
    params.Add("boundary.circle.radius", "1000")
    params.Add("format", "json")
    params.Add("text", r.PathValue("text"))

    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	h.logger.Info("Server started", slog.String("port", fullURL), slog.String("env", r.PathValue("text")))

    resp, err := http.Get(fullURL)
    if err != nil {
        log.Fatalf("Error fetching data", err)
        fmt.Printf("Error fetching data", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
         
    }

     
    var locations []geojson.Feature
    if err := json.Unmarshal(body, &locations); err != nil {
         
    }

     
	for _, location := range locations {
		fmt.Print(location.Properties["name"])
	}
}

