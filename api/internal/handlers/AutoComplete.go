package handlers

import (
	// "encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	// "github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/gin-gonic/gin"
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


func (h *AutoCompleteHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {

	baseURL := "http://192.168.1.227:4000/v1/autocomplete"


    params := url.Values{}
    params.Add("boundary.circle.lat", "64")
    params.Add("boundary.circle.lon", "17")
    params.Add("boundary.circle.radius", "1000")
    params.Add("format", "json")
    params.Add("text", c.Query("text"))

    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    resp, err := http.Get(fullURL)

    if err != nil {
        h.logger.Error("Error reading response: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        h.logger.Error("Error reading response: %v", err)
    }
	h.logger.Info(string(body))

    // Unmarshal the JSON response into a slice of Location
    // var locations []geojson.GeojsonEncoder
    // if err := json.Unmarshal(body, &locations); err != nil {
    //     h.logger.Error("Error reading response: %v", err)
    // }

	// for _, location := range locations {
		// Print the JSON string of the feature
		// h.logger.Info(string(location))
	// }
}

