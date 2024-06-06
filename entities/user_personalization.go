package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalization struct {
	Id                        uuid.UUID
	UserId                    uuid.UUID
	User                      User
	PersonalizationProvinces  []PersonalizationProvince
	PersonalizationCategories []PersonalizationCategory
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 gorm.DeletedAt
}
