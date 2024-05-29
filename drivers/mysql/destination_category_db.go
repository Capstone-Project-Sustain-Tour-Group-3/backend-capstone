package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationCategory struct {
	Id            uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Destination   Destination    `gorm:"foreignKey:DestinationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"destination"` //nolint:lll
	DestinationId uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"destination_id"`
	Category      Category       `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"` //nolint:lll
	CategoryId    uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"category_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
