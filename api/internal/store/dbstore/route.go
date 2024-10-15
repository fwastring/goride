package dbstore

import (
	// "goride/internal/store/db"
	"goride/internal/store"
	"goride/internal/store/types"

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


func (s *RouteStore) CreateRoute(addressFrom string, addressTo string, geometry types.Geometry4326) error {
    route := store.Route{
		StartAddress: addressFrom,
		EndAddress: addressTo,
        Geometry: geometry,
    }

	return s.db.Create(&route).Error
}

func (s *RouteStore) UpdateRoute(route store.Route) error {
	return s.db.Save(&route).Error
}

func (s *RouteStore) GetRoute(id uint) (store.Route, error) {
    var route store.Route
    err := s.db.Where("id = ?", id).First(&route).Error

    return route, err  // Returns both, route and err (nil if no error)
}

func (s *RouteStore) GetAllRoutes() ([]store.Route, error) {
    var routes []store.Route
    result := s.db.Find(&routes)

	if result.Error != nil {
		return nil, result.Error
	}

    return routes, nil
}
