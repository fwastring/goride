package handlers

import (
	 
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"goride/store/dbstore"
	"goride/store/types"

	"github.com/gin-gonic/gin"
)

type SearchRouteHandler struct {
	logger      slog.Logger
	database    *sql.DB
	routeStore  dbstore.RouteStore
}

type SearchRouteHandlerParams struct {
	Logger      slog.Logger
	Database    *sql.DB
	RouteStore  dbstore.RouteStore
}

func NewSearchRouteHandler(params SearchRouteHandlerParams) *SearchRouteHandler {
	return &SearchRouteHandler{
		logger:     params.Logger,
		database:   params.Database,
		routeStore: params.RouteStore,
	}
}

func (h *SearchRouteHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	 
	var addresses struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	if err := c.ShouldBindJSON(&addresses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	points := []types.Point{getLocation(addresses.From, h.logger), getLocation(addresses.To, h.logger)}
	fmt.Print(points[0])
	fmt.Print(points[1])

	 
	routeResults, err := h.routeStore.GetRoutesOnPoints(points[0], points[1], 10000)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding routes"})
		return
	}

	c.JSON(http.StatusOK, routeResults)
}

