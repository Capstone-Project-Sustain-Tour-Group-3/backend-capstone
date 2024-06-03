package entities

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	Id         string
	ProvinceId string
	Province   Province
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
