package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationAddress struct {
	Id            uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Destination   Destination    `gorm:"foreignKey:DestinationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"destination"` //nolint:lll
	DestinationId uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"destination_id"`
	Province      Province       `gorm:"foreignKey:ProvinceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"province"` //nolint:lll
	ProvinceId    uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"province_id"`
	City          string         `gorm:"type:varchar(255);not null" json:"city"`
	Regency       string         `gorm:"type:varchar(255);not null" json:"regency"`
	StreetName    string         `gorm:"type:varchar(255);not null" json:"street_name"`
	PostalCode    string         `gorm:"type:varchar(255);not null" json:"postal_code"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
