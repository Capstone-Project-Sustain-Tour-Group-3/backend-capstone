package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type IDestinationUsecase interface {
	SearchDestinations(page, limit int, searchQuery, sortQuery string) (*int64, *[]dto.SearchDestinationsResponse, error)
	DetailDestination(id uuid.UUID) (*dto.DetailDestinationResponse, error)
}

type DestinationUsecase struct {
	destinationRepo repositories.IDestinationRepository
}

func NewDestinationUsecase(destinationRepo repositories.IDestinationRepository) *DestinationUsecase {
	return &DestinationUsecase{destinationRepo}
}

func (uc *DestinationUsecase) SearchDestinations(page, limit int, searchQuery, sortQuery string) (*int64, *[]dto.SearchDestinationsResponse, error) {
	total, destinations, err := uc.destinationRepo.FindAll(page, limit, searchQuery, sortQuery)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data destinasi"}
	}
	response := dto.ToSearchDestinationsResponse(&destinations)
	return total, response, nil
}

func (uc *DestinationUsecase) DetailDestination(id uuid.UUID) (*dto.DetailDestinationResponse, error) {
	destination, err := uc.destinationRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := &dto.DetailDestinationResponse{
		Destination: destination,
	}

	return response, nil
}
