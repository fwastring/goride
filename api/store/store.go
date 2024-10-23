package store

import (
	"time"

	"github.com/spatial-go/geoos/geoencoding/geojson"
)

 
 
 
 
 
 
 
 
 

type User struct {
	ID       uint   `json:"id"`
	Username    string `json:"username"`
	Password string `json:"-"`
}

type Session struct {
	ID        uint   `json:"id"`
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	LastSeen  time.Time `json:"last_seen"`
	User 	  User	 `json:"user"`
}

type UserStore interface {
	CreateUser(username string, password string) error
	GetUser(username string) (*User, error)
}

type SessionStore interface {
	CreateSession(session *Session) (*Session, error)
	GetUserFromSession(sessionID string, userID string) (*User, error)
}


type Route struct {
	ID       uint `json:"id"`
	StartAddress string `json:"from"`
	EndAddress string `json:"to"`
	Geometry Geometry `json:"geometry"`
}

type OsrmResponse struct {
	Route []Route `json:"routes"`
}

 
type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
	Type        string      `json:"type"`
}


type Location struct {
    DisplayName string  `json:"display_name"`
    Lat         string  `json:"lat"`
    Lon         string  `json:"lon"`
}

type RouteResult struct {
	RouteID int
	OnRoute bool

}

type RouteStore interface {
	CreateRoute(coordinates geojson.Geometry) error
	GetRoute(id uint) (*Route, error)
}
