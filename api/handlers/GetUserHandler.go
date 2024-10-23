package handlers

import (
	"database/sql"
	"goride/store/dbstore"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserHandler struct {
	logger         slog.Logger
	database 	   *sql.DB
	userStore	   dbstore.UserStore
}

type GetUserHandlerParams struct {
	Logger         slog.Logger
	Database 	   *sql.DB
	UserStore	   dbstore.UserStore
}

func NewGetUserHandler(params GetUserHandlerParams) *GetUserHandler {
	return &GetUserHandler{
		logger: params.Logger,	
		database: params.Database,
		userStore: params.UserStore,
	}
}

func (h *GetUserHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	 
	 
  
}

