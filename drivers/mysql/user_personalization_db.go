package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalization struct {
	Id        uuid.UUID      `gorm:"primaryKey;not null"`
	UserId    uuid.UUID      `gorm:"type:varchar(191);index;unique;not null"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at;index"`
}
