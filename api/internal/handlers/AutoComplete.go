package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"

	// "strconv"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	// "gorm.io/gorm/logger"
)

type AutoCompleteHandler struct {
	logger         slog.Logger
}

type AutoCompleteHandlerParams struct {
	Logger         slog.Logger
}


func NewAutoCompleteHandler(params AutoCompleteHandlerParams) *AutoCompleteHandler {
	return &AutoCompleteHandler{
		logger: params.Logger,	
	}
}


func (h *AutoCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	baseURL := "http://192.168.1.227:4000/v1/autocomplete"


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
        log.Fatalf("Error fetching data: %v", err)
        fmt.Printf("Error fetching data: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // logger.Fatalf("Error reading response: %v", err)
    }

    // Unmarshal the JSON response into a slice of Location
    var locations []geojson.Feature
    if err := json.Unmarshal(body, &locations); err != nil {
        // log.Fatalf("Error unmarshaling JSON: %v", err)
    }

    // Print the location data
	for _, location := range locations {
		fmt.Print(location.Properties["name"])
	}
}

