package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalization struct {
	Id         uuid.UUID      `gorm:"primaryKey;not null"`
	UserId     uuid.UUID      `gorm:"type:varchar(191);index;priority:12;not null"`
	User       User           `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProvinceId uuid.UUID      `gorm:"type:varchar(191);index;not null"`
	Province   Province       `gorm:"foreignKey:ProvinceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CategoryId uuid.UUID      `gorm:"type:varchar(191);index;not null"`
	Category   Category       `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time      `gorm:"created_at"`
	UpdatedAt  time.Time      `gorm:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"deleted_at;index"`
}
