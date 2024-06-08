package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationMedia struct {
	Id            uuid.UUID
	DestinationId uuid.UUID
	Destination   Destination
	Url           string
	Type          string
	Title         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
