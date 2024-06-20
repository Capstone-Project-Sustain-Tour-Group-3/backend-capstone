package dto

import (
	"mime/multipart"

	"capstone/entities"

	"github.com/google/uuid"
)

type UploadDestinationMediaRequest struct {
	DestinationId uuid.UUID      `json:"destination_id" form:"destination_id" validate:"required"`
	File          multipart.File `json:"file" form:"file" validate:"required"`
	Title         string         `json:"title" form:"title" validate:"required"`
}

type UpdateImageDestinationMediaRequest struct {
	File  multipart.File `json:"file" form:"file" validate:"required"`
	Title string         `json:"title" form:"title" validate:"required"`
}

type CreateDestinationMediaRequest struct {
	DestinationId uuid.UUID `json:"destination_id" validate:"required"`
	Url           string    `json:"url" validate:"required,fileext=mp4 mov"`
	Type          string    `json:"type" validate:"required,oneof=image video"`
	Title         string    `json:"title" validate:"required"`
}

type Destination struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateDestinationMediaRequest struct {
	DestinationId uuid.UUID `json:"destination_id" validate:"required"`
	Url           string    `json:"url" validate:"required,fileext=mp4 mov"`
	Type          string    `json:"type" validate:"required,oneof=image video"`
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

func ToGetAllDestinationMediaResponse(destinationMedias []entities.DestinationMedia) []GetDetailDestinationMediaResponse {
	var response []GetDetailDestinationMediaResponse
	for _, destinationMedia := range destinationMedias {
		response = append(response, ToGetDetailDestinationMediaResponse(destinationMedia))
	}
	return response
}

type GetDetailDestinationMediaResponse struct {
	Id            uuid.UUID   `json:"id"`
	DestinationId uuid.UUID   `json:"destination_id"`
	Url           string      `json:"url"`
	Type          string      `json:"type"`
	Title         string      `json:"title"`
	Destination   Destination `json:"destination"`
}
