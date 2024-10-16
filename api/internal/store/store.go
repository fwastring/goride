package store

import (
	"goride/internal/store/types"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"gorm.io/gorm"
)

// type osrmresponse struct {
// 	route []routeresponse `json:"routes"`
// }
//
// type routeresponse struct {
//     ID       uint         
//     Geometry Geometry
// }
//
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}


type Route struct {
    ID       uint         `gorm:"primaryKey"`
	StartAddress string
	EndAddress string
    Geometry types.Geometry4326 `gorm:"type:geometry"`  // Geometry field with PostGIS
}

type Point struct {
	Latitude float64
	Longitude float64
}

type Location struct {
    DisplayName string  `json:"display_name"`
    Lat         string  `json:"lat"`
    Lon         string  `json:"lon"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type RouteStore interface {
	CreateRoute(coordinates geojson.Geometry) error
	GetRoute(id uint) (*Route, error)
}
