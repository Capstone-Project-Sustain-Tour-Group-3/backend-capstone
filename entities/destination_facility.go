package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationFacility struct {
	Id            uuid.UUID
	DestinationId uuid.UUID
	FacilityId    uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
