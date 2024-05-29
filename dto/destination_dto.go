package dto

import (
	"capstone/entities"
	"github.com/google/uuid"
)

type SearchDestinationsResponse struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"nama"`
	Url      *string   `json:"url_media"`
	Province string    `json:"provinsi"`
	City     string    `json:"kota"`
}

func ToSearchDestinationsResponse(destinations *[]entities.Destination) *[]SearchDestinationsResponse {
	responses := make([]SearchDestinationsResponse, len(*destinations))

	for idx, destination := range *destinations {
		var url *string

		if len(destination.DestinationMedias) > 0 {
			url = &destination.DestinationMedias[0].Url
		}

		responses[idx] = SearchDestinationsResponse{
			Id:       destination.Id,
			Name:     destination.Name,
			Url:      url,
			Province: destination.DestinationAddress.Province.Name,
			City:     destination.DestinationAddress.City,
		}
	}

	return &responses
}

type DetailDestinationResponse struct {
	Destination *entities.Destination `json:"destination"`
}
