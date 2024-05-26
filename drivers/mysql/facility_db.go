package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	Id          uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Url         string         `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
