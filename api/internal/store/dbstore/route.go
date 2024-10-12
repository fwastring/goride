package dbstore

import (
	"goride/internal/types"

	// "github.com/spatial-go/geoos/geoencoding/geojson"
	"gorm.io/gorm"
)

type RouteStore struct {
	db           *gorm.DB
}

type NewRouteStoreParams struct {
	DB           *gorm.DB
}

func NewRouteStore(params NewRouteStoreParams) *RouteStore {
	return &RouteStore{
		db:           params.DB,
	}
}

func (s *RouteStore) CreateRoute(geometry types.Geometry) error {
	return s.db.Create(&types.Route{
		Geometry: geometry,
	}).Error
}

func (s *RouteStore) GetRoute(id uint) (*types.Route, error) {

	var route types.Route
	err := s.db.Where("id = ?", id).First(&route).Error

	if err != nil {
		return nil, err
	}
	return &route, err
}
