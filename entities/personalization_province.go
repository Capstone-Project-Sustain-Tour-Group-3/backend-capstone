package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalizationProvince struct {
	Id                    uuid.UUID
	UserPersonalizationId uuid.UUID
	ProvinceId            string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
