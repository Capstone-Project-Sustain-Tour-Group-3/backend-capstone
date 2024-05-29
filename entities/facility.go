package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Facility struct {
	Id        uuid.UUID
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
