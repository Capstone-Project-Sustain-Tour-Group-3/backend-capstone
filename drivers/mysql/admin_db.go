package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Id              uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Email           string         `gorm:"type:varchar(255);not null;unique"`
	Username        string         `gorm:"type:varchar(255);not null;unique"`
	Password        string         `gorm:"type:varchar(255);not null"`
	ProfileImageURL string         `gorm:"type:varchar(255)"`
	Role            string         `gorm:"type:enum('admin', 'super admin');not null"`
	RefreshToken    string         `gorm:"type:varchar(255)"`
	CreatedAt       time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt       time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt       gorm.DeletedAt `gorm:"type:timestamp"`
}
