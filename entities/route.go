package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Route struct {
	Id             uuid.UUID
	UserId         uuid.UUID
	CityId         uuid.UUID
	City           City
	User           User
	Name           string
	StartLongitude float64
	StartLatitude  float64
	Price          float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
