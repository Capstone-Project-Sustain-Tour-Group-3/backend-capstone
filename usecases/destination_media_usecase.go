package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type IDestinationMediaUsecase interface {
	Create(request dto.CreateDestinationMediaRequest) error
	FindAll() ([]entities.DestinationMedia, error)
	FindById(id uuid.UUID) (*dto.GetDetailDestinationMediaResponse, error)
	Update(id uuid.UUID, request dto.UpdateDestinationMediaRequest) error
	Delete(id uuid.UUID) error
}

type DestinationMediaUsecase struct {
	destinationMediaRepo repositories.IDestinationMediaRepository
	destinationRepo      repositories.IDestinationRepository
}

func NewDestinationMediaUsecase(destinationMediaRepo repositories.IDestinationMediaRepository, destinationRepo repositories.IDestinationRepository) *DestinationMediaUsecase {
	return &DestinationMediaUsecase{
		destinationMediaRepo: destinationMediaRepo,
		destinationRepo:      destinationRepo,
	}
}

func (uc *DestinationMediaUsecase) Create(request dto.CreateDestinationMediaRequest) error {
	if _, err := uc.destinationRepo.FindById(request.DestinationId); err != nil {
		return &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}

	destinationMedia := entities.DestinationMedia{
		Id:            uuid.New(),
		DestinationId: request.DestinationId,
		Url:           request.Url,
		Type:          request.Type,
		Title:         request.Title,
	}

	if err := uc.destinationMediaRepo.Create(&destinationMedia, nil); err != nil {
		return err
	}

	return nil
}

func (uc *DestinationMediaUsecase) FindAll() ([]entities.DestinationMedia, error) {
	destinationMedias, err := uc.destinationMediaRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return destinationMedias, nil
}

func (uc *DestinationMediaUsecase) FindById(id uuid.UUID) (*dto.GetDetailDestinationMediaResponse, error) {
	destinationMedia, err := uc.destinationMediaRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := dto.ToGetDetailDestinationMediaResponse(*destinationMedia)

	return &response, nil
}

func (uc *DestinationMediaUsecase) Update(id uuid.UUID, request dto.UpdateDestinationMediaRequest) error {
	destinationMedia, err := uc.destinationMediaRepo.FindById(id)
	if err != nil {
		return &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}

	destinationMedia.DestinationId = request.DestinationId
	destinationMedia.Url = request.Url
	destinationMedia.Type = request.Type
	destinationMedia.Title = request.Title

	if err = uc.destinationMediaRepo.Update(destinationMedia); err != nil {
		return err
	}

	return nil
}

func (uc *DestinationMediaUsecase) Delete(id uuid.UUID) error {
	destinationMedia, err := uc.destinationMediaRepo.FindById(id)
	if err != nil {
		return &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}
	if err = uc.destinationMediaRepo.Delete(destinationMedia); err != nil {
		return err
	}
	return nil
}
