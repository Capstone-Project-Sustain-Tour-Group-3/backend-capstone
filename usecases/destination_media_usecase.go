package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/externals/cloudinary"
	"capstone/repositories"

	"github.com/google/uuid"
)

type IDestinationMediaUsecase interface {
	Create(request dto.CreateDestinationMediaRequest) error
	FindAll(page, limit int, searchQuery string) (*int64, []dto.GetDetailDestinationMediaResponse, error)
	FindById(id uuid.UUID) (*dto.GetDetailDestinationMediaResponse, error)
	Update(id uuid.UUID, request dto.UpdateDestinationMediaRequest) error
	Delete(id uuid.UUID) error
	UploadImage(request dto.UploadDestinationMediaRequest) (string, error)
	UpdateImage(id uuid.UUID, request dto.UpdateImageDestinationMediaRequest) (string, error)
}

type DestinationMediaUsecase struct {
	destinationMediaRepo repositories.IDestinationMediaRepository
	destinationRepo      repositories.IDestinationRepository
	cloudinaryClient     cloudinary.ICloudinaryClient
}

func NewDestinationMediaUsecase(destinationMediaRepo repositories.IDestinationMediaRepository, destinationRepo repositories.IDestinationRepository, cloudinaryClient cloudinary.ICloudinaryClient) *DestinationMediaUsecase {
	return &DestinationMediaUsecase{
		destinationMediaRepo: destinationMediaRepo,
		destinationRepo:      destinationRepo,
		cloudinaryClient:     cloudinaryClient,
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

func (uc *DestinationMediaUsecase) FindAll(page, limit int, searchQuery string) (*int64, []dto.GetDetailDestinationMediaResponse, error) {
	total, destinationMedias, err := uc.destinationMediaRepo.FindAll(page, limit, searchQuery)
	if err != nil {
		return nil, nil, err
	}

	response := dto.ToGetAllDestinationMediaResponse(destinationMedias)

	return total, response, nil
}

func (uc *DestinationMediaUsecase) FindById(id uuid.UUID) (*dto.GetDetailDestinationMediaResponse, error) {
	destinationMedia, err := uc.destinationMediaRepo.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "Konten media destinasi tidak ditemukan"}
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
	tx := uc.destinationMediaRepo.BeginTx()
	if err = uc.destinationMediaRepo.Delete(destinationMedia, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err = uc.cloudinaryClient.DeleteImage(destinationMedia.Url); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menghapus media destinasi"}
	}

	tx.Commit()

	return nil
}

func (uc *DestinationMediaUsecase) UploadImage(request dto.UploadDestinationMediaRequest) (string, error) {
	url, err := uc.cloudinaryClient.UploadImage(request.File, "destination_media")
	if err != nil {
		return "", err
	}

	destinationMedia := dto.ToDestinationMedia(request.DestinationId, "image", url, request.Title)

	if err = uc.destinationMediaRepo.Create(destinationMedia, nil); err != nil {
		return "", err
	}

	return url, nil
}

func (uc *DestinationMediaUsecase) UpdateImage(id uuid.UUID, request dto.UpdateImageDestinationMediaRequest) (string, error) {
	destinationMedia, err := uc.destinationMediaRepo.FindById(id)
	if err != nil {
		return "", &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}

	var oldUrl string = destinationMedia.Url

	url, err := uc.cloudinaryClient.UploadImage(request.File, "destination_media")
	if err != nil {
		return "", err
	}

	destinationMedia.Url = url
	destinationMedia.Title = request.Title

	if err = uc.destinationMediaRepo.Update(destinationMedia); err != nil {
		return "", err
	}

	if err = uc.cloudinaryClient.DeleteImage(oldUrl); err != nil {
		return "", &errorHandlers.InternalServerError{Message: "Gagal menghapus media destinasi"}
	}

	return url, nil
}
