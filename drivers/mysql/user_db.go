package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id              uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Email           string         `gorm:"type:varchar(255);not null" json:"email"`
	Password        string         `gorm:"type:varchar(255);not null" json:"password"`
	Username        string         `gorm:"type:varchar(255);not null" json:"username"`
	Fullname        string         `gorm:"type:varchar(255);not null" json:"fullname"`
	Bio             string         `gorm:"type:varchar(255)" json:"bio"`
	PhoneNumber     string         `gorm:"type:varchar(255);not null" json:"phone_number"`
	ProfileImageUrl string         `gorm:"type:varchar(255)" json:"profile_image_url"`
	Gender          string         `gorm:"type:varchar(255)" json:"gender"`
	Country         string         `gorm:"type:varchar(255)" json:"country"`
	City            string         `gorm:"type:varchar(255)" json:"city"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
