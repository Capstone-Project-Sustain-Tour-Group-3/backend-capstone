package entities

import (
	"time"

	"github.com/google/uuid"
)

type RouteDetail struct {
	Id            uuid.UUID
	DestinationId uuid.UUID
	Destination   Destination
	RouteId       uuid.UUID
	Route         Route
	Longitude     float64
	Latitude      float64
	Duration      int
	Order         int
	VisitStart    []uint8
	VisitEnd      []uint8
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
