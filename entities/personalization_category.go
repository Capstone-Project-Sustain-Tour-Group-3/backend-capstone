package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalizationCategory struct {
	Id                    uuid.UUID
	UserPersonalizationId uuid.UUID
	CategoryId            uuid.UUID
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
