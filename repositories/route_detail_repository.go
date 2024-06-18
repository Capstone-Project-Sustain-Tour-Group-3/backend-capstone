package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RouteDetailRepository interface {
	Create(routeDetail *entities.RouteDetail) error
	DeleteMany(routeDetail *[]entities.RouteDetail) error
}

type routeDetailRepository struct {
	db *gorm.DB
}

func NewRouteDetailRepository(db *gorm.DB) *routeDetailRepository {
	return &routeDetailRepository{db}
}

func (r *routeDetailRepository) Create(routeDetail *entities.RouteDetail) error {
	if err := r.db.Create(&routeDetail).Error; err != nil {
		return err
	}
	return nil
}

func (r *routeDetailRepository) DeleteMany(routeDetail *[]entities.RouteDetail) error {
	if len(*routeDetail) == 0 {
		return nil
	}

	var ids []uuid.UUID
	for _, routedetail := range *routeDetail {
		ids = append(ids, routedetail.Id)
	}

	if err := r.db.Where("id IN ?", ids).Delete(&entities.RouteDetail{}).Error; err != nil {
		return err
	}

	return nil
}
