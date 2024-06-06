package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Facility struct {
	Id        uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Url       string         `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
