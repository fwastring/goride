package handlers

import (
	"goride/internal/store/dbstore"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetRouteByIDHandler struct {
	logger         slog.Logger
	database 	   gorm.DB
}

type GetRouteByIDHandlerParams struct {
	Logger         slog.Logger
	Database 	   gorm.DB
}

func NewGetRouteByIDHandler(params GetRouteByIDHandlerParams) *GetRouteByIDHandler {
	return &GetRouteByIDHandler{
		logger: params.Logger,	
		database: params.Database,
	}
}

func (h *GetRouteByIDHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	idInt, _ := strconv.Atoi(c.Query("id"))
	id := uint(idInt)

	routeStore := dbstore.NewRouteStore(dbstore.NewRouteStoreParams{DB: &h.database})

	route, err := routeStore.GetRoute(id)
	if err != nil {
        h.logger.Error("Error reading response: %v", err)
	}

    // Send response as JSON
    c.JSON(http.StatusOK, route)
}

