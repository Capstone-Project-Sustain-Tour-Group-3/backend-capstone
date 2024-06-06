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
	ProvinceId    string         `gorm:"type:char(2);index;not null" json:"province_id"`
	City          City           `gorm:"foreignKey:CityId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"city"` //nolint:lll
	CityId        string         `gorm:"type:char(4);index;not null" json:"city_id"`
	Subdistrict   Subdistrict    `gorm:"foreignKey:SubdistrictId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"subdistrict"` //nolint:lll
	SubdistrictId string         `gorm:"type:char(6);index;not null" json:"subdistrict_id"`
	StreetName    string         `gorm:"type:varchar(255);not null" json:"street_name"`
	PostalCode    string         `gorm:"type:varchar(255);not null" json:"postal_code"`
	CreatedAt     time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
