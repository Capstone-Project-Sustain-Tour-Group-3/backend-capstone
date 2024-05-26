package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Province struct {
	Id                 uuid.UUID
	DestinationAddress DestinationAddress
	Name               string
	Url                string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}
