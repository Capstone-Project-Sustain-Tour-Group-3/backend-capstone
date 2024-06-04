package entities

import (
	"time"

	"gorm.io/gorm"
)

type Subdistrict struct {
	Id        string
	CityId    string
	City      City
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
