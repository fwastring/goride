package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"

	"goride/store/dbstore"

	"github.com/gin-gonic/gin"
)

type DeleteRouteHandler struct {
	logger         slog.Logger
	database 	   *sql.DB
	routeStore	   dbstore.RouteStore
}

type DeleteRouteHandlerParams struct {
	Logger         slog.Logger
	Database 	   *sql.DB
	RouteStore	   dbstore.RouteStore
}


func NewDeleteRouteHandler(params DeleteRouteHandlerParams) *DeleteRouteHandler {
	return &DeleteRouteHandler{
		logger: params.Logger,	
		database: params.Database,
		routeStore: params.RouteStore,
	}
}

func (h *DeleteRouteHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	 
	var DeleteObject struct {
		ID string `json:"id"`
	}

	if err := c.ShouldBindJSON(&DeleteObject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	id, err := strconv.ParseUint(DeleteObject.ID, 10, 0)

	err = h.routeStore.DeleteRoute(uint(id))
	if err != nil {
        h.logger.Error("Error deleting route", err)
	}

     
    c.JSON(http.StatusOK, "Deleted route")
}
