package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Route struct {
	Id             uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	UserId         uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"user_id"`
	User           User           `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CityId         uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"city_id"`
	City           City           `gorm:"foreignKey:CityId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"city"`
	Name           string         `gorm:"type:varchar(255);not null" json:"name"`
	StartLongitude float64        `gorm:"type:decimal(9,6);not null" json:"start_longitude"`
	StartLatitude  float64        `gorm:"type:decimal(9,6);not null" json:"start_latitude"`
	Price          float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt      time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt      time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
