package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id              uuid.UUID
	Email           string
	Password        string
	Username        string
	Fullname        string
	Bio             string
	PhoneNumber     string
	ProfileImageUrl string
	Gender          string
	City            string
	Province        string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
