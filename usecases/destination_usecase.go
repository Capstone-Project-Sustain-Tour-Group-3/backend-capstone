package usecases

import (
	"errors"
	"fmt"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/externals/cloudinary"
	"capstone/repositories"

	"github.com/google/uuid"
)

type IDestinationUsecase interface {
	SearchDestinations(page, limit int, searchQuery, sortQuery, filterQuery string) (string, *int64, *[]dto.SearchDestination, error)
	DetailDestination(id uuid.UUID) (*dto.DetailDestinationResponse, error)
	CreateDestination(destinationReq *dto.CreateDestinationRequest) error
}

type DestinationUsecase struct {
	destinationRepo         repositories.IDestinationRepository
	destinationFacilityRepo repositories.IDestinationFacilityRepository
	destinationMediaRepo    repositories.IDestinationMediaRepository
	destinationAddressRepo  repositories.IDestinationAddressRepository
	categoryRepo            repositories.ICategoryRepository
	cloudinaryClient        cloudinary.ICloudinaryClient
}

func NewDestinationUsecase(
	destinationRepo repositories.IDestinationRepository,
	destinationFacilityRepo repositories.IDestinationFacilityRepository,
	destinationMediaRepo repositories.IDestinationMediaRepository,
	destinationAddressRepo repositories.IDestinationAddressRepository,
	categoryRepo repositories.ICategoryRepository,
	cloudinaryClient cloudinary.ICloudinaryClient,
) *DestinationUsecase {
	return &DestinationUsecase{
		destinationRepo,
		destinationFacilityRepo,
		destinationMediaRepo,
		destinationAddressRepo,
		categoryRepo,
		cloudinaryClient,
	}
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

func (uc *DestinationUsecase) CreateDestination(destinationReq *dto.CreateDestinationRequest) error {
	category, err := uc.categoryRepo.FindById(destinationReq.DestinationInfo.CategoryId)
	if err != nil {
		return errors.New("category not found")
	}

	if category == nil {
		return errors.New("category not found")
	}

	destination := &entities.Destination{
		Id:          uuid.New(),
		CategoryId:  destinationReq.DestinationInfo.CategoryId,
		Name:        destinationReq.DestinationInfo.Name,
		Description: destinationReq.DestinationInfo.Description,
		OpenTime:    destinationReq.DestinationInfo.OpenTime,
		CloseTime:   destinationReq.DestinationInfo.CloseTime,
		EntryPrice:  destinationReq.DestinationInfo.EntryPrice,
		Latitude:    destinationReq.DestinationInfo.Latitude,
		Longitude:   destinationReq.DestinationInfo.Longitude,
		VisitCount:  0,
	}

	if err = uc.destinationRepo.Create(destination); err != nil {
		return fmt.Errorf("error when create destination: %w", err)
	}

	destinationFacilities := dto.ToDestinationFacilities(destination.Id, destinationReq.DestinationInfo.FacilityIds)

	if err = uc.destinationFacilityRepo.Create(destinationFacilities); err != nil {
		return fmt.Errorf("error when create destination facility: %w", err)
	}

	for _, image := range destinationReq.DestinationImages {
		urlMedia, errFile := uc.cloudinaryClient.UploadImage(image.File)

		if errFile != nil {
			return fmt.Errorf("error when upload image to cloud: %w", errFile)
		}

		destinationMedia := dto.ToDestinationMedia(destination.Id, "image", urlMedia, image.Title)

		if err = uc.destinationMediaRepo.Create(destinationMedia); err != nil {
			return fmt.Errorf("error when create destination media: %w", err)
		}
	}

	destinationAddress := dto.ToDestinationAddress(destination.Id, destinationReq.DestinationAddress)

	if err = uc.destinationAddressRepo.Create(destinationAddress); err != nil {
		return fmt.Errorf("error when create destination address: %w", err)
	}

	return nil
}
