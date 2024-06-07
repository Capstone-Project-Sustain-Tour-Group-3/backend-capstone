package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Admin struct {
	Id              uuid.UUID             `gorm:"primaryKey;not null" json:"id"`
	Username        string                `gorm:"type:varchar(255);not null;uniqueIndex:us_active"`
	Password        string                `gorm:"type:varchar(255);not null"`
	ProfileImageURL string                `gorm:"type:varchar(255)"`
	Role            string                `gorm:"type:enum('admin', 'super admin');not null"`
	RefreshToken    string                `gorm:"type:text"`
	CreatedAt       time.Time             `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt       time.Time             `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt       gorm.DeletedAt        `gorm:"index"`
	DeleteMilli     soft_delete.DeletedAt `gorm:"softDelete:milli;uniqueIndex:us_active"`
}
