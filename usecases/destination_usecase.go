package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type IDestinationUsecase interface {
	SearchDestinations(page, limit int, searchQuery, sortQuery, filterQuery string) (string, *int64, *[]dto.SearchDestination, error)
	DetailDestination(id uuid.UUID) (*dto.DetailDestinationResponse, error)
}

type DestinationUsecase struct {
	destinationRepo repositories.IDestinationRepository
}

func NewDestinationUsecase(destinationRepo repositories.IDestinationRepository) *DestinationUsecase {
	return &DestinationUsecase{destinationRepo}
}

func (uc *DestinationUsecase) SearchDestinations(page, limit int, searchQuery, sortQuery, filterQuery string) (string, *int64, *[]dto.SearchDestination, error) {
	categoryName, total, destinations, err := uc.destinationRepo.FindAll(page, limit, searchQuery, sortQuery, filterQuery)
	if err != nil {
		return categoryName, nil, nil, &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data destinasi"}
	}
	response := dto.ToSearchDestinationsResponse(&destinations)
	return categoryName, total, response, nil
}

func (uc *DestinationUsecase) DetailDestination(id uuid.UUID) (*dto.DetailDestinationResponse, error) {
	destination, err := uc.destinationRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	similarDestinations, err := uc.destinationRepo.FindByCategoryId(destination.CategoryId)
	if err != nil {
		return nil, err
	}

	response := dto.ToDetailDestinationResponse(destination, &similarDestinations)

	return response, nil
}
