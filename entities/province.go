package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationAddress struct {
	Id            uuid.UUID
	DestinationId uuid.UUID
	ProvinceId    uuid.UUID
	City          string
	Subdistrict   string
	StreetName    string
	PostalCode    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
