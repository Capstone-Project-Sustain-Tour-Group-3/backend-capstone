package mysql

import (
	"time"

	"gorm.io/gorm"
)

type Subdistrict struct {
	Id        string         `gorm:"primaryKey;type:char(6);not null" json:"id"`
	City      City           `gorm:"foreignKey:CityId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"city"` //nolint:lll
	CityId    string         `gorm:"type:char(4);index;not null" json:"city_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
