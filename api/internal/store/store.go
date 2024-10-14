package store

import (
	"goride/internal/store/types"

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
    ID       uint         `gorm:"primaryKey"`
    Geometry types.Geometry4326 `gorm:"type:geometry"`  // Geometry field with PostGIS
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
