package entities

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	Id        string
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
