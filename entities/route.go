package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Route struct {
	Id             uuid.UUID
	UserId         uuid.UUID
	CityId         string
	City           City
	User           User
	RouteDetail    []RouteDetail
	Name           string
	StartLongitude float64
	StartLatitude  float64
	Price          float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
