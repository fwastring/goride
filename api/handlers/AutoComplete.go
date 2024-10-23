package handlers

import (
	 
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	 
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

	baseURL := "http: 


    params := url.Values{}
    params.Add("boundary.circle.lat", "64")
    params.Add("boundary.circle.lon", "17")
    params.Add("boundary.circle.radius", "1000")
    params.Add("format", "json")
    params.Add("text", c.Query("text"))

    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    resp, err := http.Get(fullURL)

    if err != nil {
        h.logger.Error("Error reading response", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        h.logger.Error("Error reading response", err)
    }
	h.logger.Info(string(body))

     
     
     
     
     

	 
		 
		 
	 
}

