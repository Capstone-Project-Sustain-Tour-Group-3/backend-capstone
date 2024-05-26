package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Destination struct {
	Id                   uuid.UUID
	//DestinationAddressId DestinationAddress
	Name                 string
	Description          string
	OpenTime             string
	ClosedTime           string
	EntryPrice           float64
	Longitude            float64
	Latitude             float64
	VisitCount           int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
}
