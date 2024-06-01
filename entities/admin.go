package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Id              uuid.UUID
	Email           string
	Username        string
	Password        string
	ProfileImageURL string
	Role            string
	RefreshToken    string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
