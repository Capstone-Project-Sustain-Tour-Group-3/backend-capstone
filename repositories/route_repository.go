package repositories

import (
	"strings"

	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RouteRepository interface {
	FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error)
	FindById(id uuid.UUID) (*entities.Route, error)
	Delete(route *entities.Route) error

	FindVisitedByUserSubquery(id uuid.UUID) *gorm.DB
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) *routeRepository {
	return &routeRepository{db: db}
}

func (r *routeRepository) FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error) {
	var routes []entities.Route
	var total int64
	offset := (page - 1) * limit
	db := r.db.Model(&entities.Route{})
	if searchQuery != "" {
		db = db.Joins("JOIN cities ON cities.id = routes.city_id").
			Where("LOWER(cities.name) LIKE ?", "%"+strings.ToLower(searchQuery)+"%")
	}

	if err := db.Debug().
		Offset(offset).Limit(limit).
		Preload("User").
		Preload("City").
		Preload("RouteDetail").
		Preload("RouteDetail.Destination").
		Find(&routes).
		Error; err != nil {
		return nil, nil, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return &routes, &total, nil
}

func (r *routeRepository) FindById(id uuid.UUID) (*entities.Route, error) {
	var route entities.Route
	if err := r.db.Debug().
		Preload("User").
		Preload("City").
		Preload("RouteDetail").
		Preload("RouteDetail.Destination").
		Where("id = ?", id).
		First(&route).
		Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *routeRepository) Delete(route *entities.Route) error {
	if err := r.db.Delete(&route).Error; err != nil {
		return err
	}
	return nil
}

func (r *routeRepository) FindVisitedByUserSubquery(id uuid.UUID) *gorm.DB {
	return r.db.Model(&entities.Route{}).
		Distinct("route_details.destination_id").
		InnerJoins("RouteDetail").
		InnerJoins("RouteDetail.Destination").
		Where("user_id = ?", id)
}
