package mysql

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	Id         string         `gorm:"primaryKey;type:char(4);not null" json:"id"`
	Province   Province       `gorm:"foreignKey:ProvinceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"province"` //nolint:lll
	ProvinceId string         `gorm:"type:char(2);index;not null" json:"province_id"`
	Name       string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
