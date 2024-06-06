package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalizationCategory struct {
	Id                    uuid.UUID           `gorm:"primaryKey;not null"`
	UserPersonalizationId uuid.UUID           `gorm:"type:varchar(191);index;uniqueIndex:uniq_pers_category;not null"`
	UserPersonalization   UserPersonalization `gorm:"foreignKey:UserPersonalizationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CategoryId            uuid.UUID           `gorm:"type:varchar(191);index;uniqueIndex:uniq_pers_category;not null"`
	Category              Category            `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
