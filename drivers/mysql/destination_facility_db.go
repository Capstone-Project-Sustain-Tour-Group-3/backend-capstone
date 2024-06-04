package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationFacility struct {
	Id            uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Destination   Destination    `gorm:"foreignKey:DestinationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"destination"` //nolint:lll
	DestinationId uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"destination_id"`
	Facility      Facility       `gorm:"foreignKey:FacilityId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"facility"` //nolint:lll
	FacilityId    uuid.UUID      `gorm:"type:varchar(191);index;not null" json:"facility_id"`
	CreatedAt     time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
