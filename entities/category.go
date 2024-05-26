package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationCategory struct {
	Id                 uuid.UUID
	DestinationId      uuid.UUID
	CategoryId         uuid.UUID
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}
