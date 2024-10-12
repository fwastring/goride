package store

import (
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Route struct {
	gorm.Model
	ID		 uint 	`gorm:"primaryKey" json:"id"`
	Geometry Geometry `gorm:"type:geometry"`
}

type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
	Type        string      `json:"type"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type RouteStore interface {
	CreateRoute(coordinates geojson.Geometry) error
	GetRoute(id uint) (*Route, error)
}
