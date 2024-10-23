package handlers

import (
	"database/sql"
	"goride/store/dbstore"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetRouteByIDHandler struct {
	logger         slog.Logger
	database 	   *sql.DB
	routeStore	   dbstore.RouteStore
}

type GetRouteByIDHandlerParams struct {
	Logger         slog.Logger
	Database 	   *sql.DB
	RouteStore	   dbstore.RouteStore
}

func NewGetRouteByIDHandler(params GetRouteByIDHandlerParams) *GetRouteByIDHandler {
	return &GetRouteByIDHandler{
		logger: params.Logger,	
		database: params.Database,
		routeStore: params.RouteStore,
	}
}

func (h *GetRouteByIDHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	 
	idInt, _ := strconv.Atoi(c.Param("id"))
	id := uint(idInt)


	route, err := h.routeStore.GetRoute(id)
	if err != nil {
        h.logger.Error("Error reading response", err)
	}

     
    c.JSON(http.StatusOK, route)
}

