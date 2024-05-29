package usecases

import (
	"capstone/dto"
	"capstone/repositories"
	"github.com/google/uuid"
)

type IDestinationUsecase interface {
	SearchDestinations() (*[]dto.SearchDestinationsResponse, error)
	DetailDestination(id uuid.UUID) (*dto.DetailDestinationResponse, error)
}

type DestinationUsecase struct {
	destinationRepo repositories.IDestinationRepository
}

func NewDestinationUsecase(destinationRepo repositories.IDestinationRepository) *DestinationUsecase {
	return &DestinationUsecase{destinationRepo}
}

func (uc *DestinationUsecase) SearchDestinations() (*[]dto.SearchDestinationsResponse, error) {
	destinations, err := uc.destinationRepo.FindAll()
	if err != nil {
		return nil, err
	}
	response := dto.ToSearchDestinationsResponse(&destinations)
	return response, nil
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
