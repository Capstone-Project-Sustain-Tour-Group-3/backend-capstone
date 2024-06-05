package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalizationProvince struct {
	Id                    uuid.UUID           `gorm:"primaryKey;not null"`
	UserPersonalizationId uuid.UUID           `gorm:"type:varchar(191);index;uniqueIndex:uniq_pers_prov;not null"`
	UserPersonalization   UserPersonalization `gorm:"foreignKey:UserPersonalizationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProvinceId            string              `gorm:"type:char(2);index;uniqueIndex:uniq_pers_prov;not null"`
	Province              Province            `gorm:"foreignKey:ProvinceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
