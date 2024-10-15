package handlers

import (
	"goride/internal/store/dbstore"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetAllRoutesHandler struct {
	logger         slog.Logger
	database 	   gorm.DB
}

type GetAllRoutesHandlerParams struct {
	Logger         slog.Logger
	Database 	   gorm.DB
}

func NewGetAllRoutesHandler(params GetAllRoutesHandlerParams) *GetAllRoutesHandler {
	return &GetAllRoutesHandler{
		logger: params.Logger,	
		database: params.Database,
	}
}

func (h *GetAllRoutesHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	routeStore := dbstore.NewRouteStore(dbstore.NewRouteStoreParams{DB: &h.database})

	routes, err := routeStore.GetAllRoutes()
	if err != nil {
        h.logger.Error("Error reading response: %v", err)
	}

    // Send response as JSON
    c.JSON(http.StatusOK, routes)
}

