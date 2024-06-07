package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Admin struct {
	Id              uuid.UUID
	Username        string
	Password        string
	ProfileImageURL *string
	Role            string
	RefreshToken    string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	DeleteMilli     soft_delete.DeletedAt
}
