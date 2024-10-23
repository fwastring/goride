package handlers

import (
	"database/sql"
	"goride/store/dbstore"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllRoutesHandler struct {
	logger         slog.Logger
	database 	   *sql.DB
	routeStore	   dbstore.RouteStore
}

type GetAllRoutesHandlerParams struct {
	Logger         slog.Logger
	Database 	   *sql.DB
	RouteStore	   dbstore.RouteStore
}

func NewGetAllRoutesHandler(params GetAllRoutesHandlerParams) *GetAllRoutesHandler {
	return &GetAllRoutesHandler{
		logger: params.Logger,	
		database: params.Database,
		routeStore: params.RouteStore,
	}
}

func (h *GetAllRoutesHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	 

	routes, err := h.routeStore.GetAllRoutes()
	if err != nil {
        h.logger.Error("Error reading response", err)
	}

     
    c.JSON(http.StatusOK, routes)
}

