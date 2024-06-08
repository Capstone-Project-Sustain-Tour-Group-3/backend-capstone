package dto

import (
	"capstone/entities"
	"github.com/google/uuid"
)

type CreateDestinationMediaRequest struct {
	DestinationId uuid.UUID `json:"destination_id" validate:"required"`
	Url           string    `json:"url" validate:"required"`
	Type          string    `json:"type" validate:"required"`
	Title         string    `json:"title" validate:"required"`
}

type Destination struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateDestinationMediaRequest struct {
	DestinationId uuid.UUID `json:"destination_id" validate:"required"`
	Url           string    `json:"url" validate:"required"`
	Type          string    `json:"type" validate:"required"`
	Title         string    `json:"title" validate:"required"`
}

func ToGetDetailDestinationMediaResponse(destinationMedia entities.DestinationMedia) GetDetailDestinationMediaResponse {
	return GetDetailDestinationMediaResponse{
		Id:            destinationMedia.Id,
		DestinationId: destinationMedia.DestinationId,
		Url:           destinationMedia.Url,
		Type:          destinationMedia.Type,
		Title:         destinationMedia.Title,
		Destination: Destination{
			Id:   destinationMedia.Destination.Id,
			Name: destinationMedia.Destination.Name,
		},
	}
}

type GetDetailDestinationMediaResponse struct {
	Id            uuid.UUID `json:"id"`
	DestinationId uuid.UUID `json:"destination_id"`
	Url           string    `json:"url"`
	Type          string    `json:"type"`
	Title         string    `json:"title"`
	Destination   Destination
}
