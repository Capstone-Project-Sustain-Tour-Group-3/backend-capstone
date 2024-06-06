package entities

import (
	"github.com/google/uuid"
	"time"
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
	VisitStart    time.Time
	VisitEnd      time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
