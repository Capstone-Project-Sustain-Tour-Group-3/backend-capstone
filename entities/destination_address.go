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
	Province      Province
	City          string
	Regency       string
	StreetName    string
	PostalCode    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
