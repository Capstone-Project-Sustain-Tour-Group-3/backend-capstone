package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Destination struct {
	Id          uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Category    Category       `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"` //nolint:lll
	CategoryId  uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"province_id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	OpenTime    string         `gorm:"type:varchar(255);not null" json:"open_time"`
	CloseTime   string         `gorm:"type:varchar(255);not null" json:"close_time"`
	EntryPrice  float64        `gorm:"type:decimal(10,2);not null" json:"entry_price"`
	Longitude   float64        `gorm:"type:decimal(9,6);not null" json:"longitude"`
	Latitude    float64        `gorm:"type:decimal(9,6);not null" json:"latitude"`
	VisitCount  int            `gorm:"type:int;not null" json:"visit_count"`
	CreatedAt   time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
