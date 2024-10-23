package handlers

import (
	b64 "encoding/base64"
	"fmt"
	"goride/hash"
	"goride/store"
	"goride/store/dbstore"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostLoginHandler struct {
	logger			  slog.Logger
	userStore         dbstore.UserStore
	sessionStore      dbstore.SessionStore
	passwordhash      hash.PasswordHash
}

type PostLoginHandlerParams struct {
	Logger			  slog.Logger
	UserStore         dbstore.UserStore
	SessionStore      dbstore.SessionStore
	PasswordHash      hash.PasswordHash
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		logger: 		   params.Logger,
		userStore:         params.UserStore,
		sessionStore:      params.SessionStore,
		passwordhash:      params.PasswordHash,
	}
}

func (h *PostLoginHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {

	var login struct {
		Username string `json:"username"`
		Password   string `json:"password"`
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := h.userStore.GetUser(login.Username)

	if err != nil {
		h.logger.Error("error",err)
		return
	}

	passwordIsValid, err := h.passwordhash.ComparePasswordAndHash(login.Password, user.Password)

	if err != nil || !passwordIsValid {
		h.logger.Error("error",err)
		return
	}

	session, err := h.sessionStore.CreateSession(&store.Session{
		UserID: user.ID,
	})

	if err != nil {
		h.logger.Error("error",err)
		return
	}

	userID := user.ID
	sessionID := session.SessionID

	cookieValue := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", sessionID, userID)))

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     "session",
		Value:    cookieValue,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)

    c.JSON(http.StatusOK, "Logged in!")
}

