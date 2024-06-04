package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalization struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	User       User
	ProvinceId string
	Province   Province
	CategoryId uuid.UUID
	Category   Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
