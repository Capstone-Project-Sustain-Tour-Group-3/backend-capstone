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
	GetAllDestinations(page, limit int, searchQuery string) (*int64, *[]dto.GetAllDestination, error)
	GetDestinationById(id uuid.UUID) (*dto.GetByIdDestinationResponse, error)
	UpdateDestination(id uuid.UUID, destinationReq *dto.UpdateDestinationRequest) error
	DeleteDestination(id uuid.UUID) error
	IncrementVisitCount(id uuid.UUID) error
	GetCitiesWithDestinations() (*[]dto.Cities, error)
	GetDestinationByCityId(id string) (*[]dto.DestinationsByCity, error)
}

type DestinationUsecase struct {
	destinationRepo         repositories.IDestinationRepository
	destinationFacilityRepo repositories.IDestinationFacilityRepository
	destinationMediaRepo    repositories.IDestinationMediaRepository
	destinationAddressRepo  repositories.IDestinationAddressRepository
	categoryRepo            repositories.ICategoryRepository
	facilityRepo            repositories.IFacilityRepository
	provinceRepo            repositories.IProvinceRepository
	cityRepo                repositories.ICityRepository
	subdistrictRepo         repositories.ISubdistrictRepository
	cloudinaryClient        cloudinary.ICloudinaryClient
}

func NewDestinationUsecase(
	destinationRepo repositories.IDestinationRepository,
	destinationFacilityRepo repositories.IDestinationFacilityRepository,
	destinationMediaRepo repositories.IDestinationMediaRepository,
	destinationAddressRepo repositories.IDestinationAddressRepository,
	categoryRepo repositories.ICategoryRepository,
	facilityRepo repositories.IFacilityRepository,
	provinceRepo repositories.IProvinceRepository,
	cityRepo repositories.ICityRepository,
	subdistrictRepo repositories.ISubdistrictRepository,
	cloudinaryClient cloudinary.ICloudinaryClient,
) *DestinationUsecase {
	return &DestinationUsecase{
		destinationRepo,
		destinationFacilityRepo,
		destinationMediaRepo,
		destinationAddressRepo,
		categoryRepo,
		facilityRepo,
		provinceRepo,
		cityRepo,
		subdistrictRepo,
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
		return nil, &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}

	similarDestinations, err := uc.destinationRepo.FindByCategoryId(destination.CategoryId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi serupa"}
	}

	response := dto.ToDetailDestinationResponse(destination, &similarDestinations)

	return response, nil
}

func (uc *DestinationUsecase) CreateDestination(destinationReq *dto.CreateDestinationRequest) error {
	category, err := uc.categoryRepo.FindById(destinationReq.CategoryId)
	if err != nil || category == nil {
		return errors.New("category not found")
	}

	destination := &entities.Destination{
		Id:          uuid.New(),
		CategoryId:  destinationReq.CategoryId,
		Name:        destinationReq.Name,
		Description: destinationReq.Description,
		OpenTime:    destinationReq.OpenTime,
		CloseTime:   destinationReq.CloseTime,
		EntryPrice:  destinationReq.EntryPrice,
		Latitude:    destinationReq.Latitude,
		Longitude:   destinationReq.Longitude,
		VisitCount:  0,
	}

	txDestination := uc.destinationRepo.BeginTx()

	if err = uc.destinationRepo.Create(destination, txDestination); err != nil {
		txDestination.Rollback()
		return fmt.Errorf("error when create destination: %w", err)
	}

	destinationFacilities := dto.ToDestinationFacilities(destination.Id, destinationReq.FacilityIds)

	if err = uc.destinationFacilityRepo.Create(destinationFacilities, txDestination); err != nil {
		txDestination.Rollback()
		return fmt.Errorf("error when create destination facility: %w", err)
	}

	destinationAddress := dto.ToDestinationAddress(destination.Id, destinationReq.DestinationAddress)

	// check province
	province, err := uc.provinceRepo.FindById(destinationAddress.ProvinceId)

	if err != nil || province == nil {
		txDestination.Rollback()
		return errors.New("province not found")
	}

	// check city
	city, err := uc.cityRepo.FindById(destinationAddress.CityId)

	if err != nil || city == nil {
		txDestination.Rollback()
		return errors.New("city not found")
	}

	// check subdistrict
	subdistrict, err := uc.subdistrictRepo.FindById(destinationAddress.SubdistrictId)

	if err != nil || subdistrict == nil {
		txDestination.Rollback()
		return errors.New("subdistrict not found")
	}

	if err = uc.destinationAddressRepo.Create(destinationAddress, txDestination); err != nil {
		txDestination.Rollback()
		return fmt.Errorf("error when create destination address: %w", err)
	}

	txDestination.Commit()

	return nil
}

func (uc *DestinationUsecase) UpdateDestination(id uuid.UUID, destinationReq *dto.UpdateDestinationRequest) error {
	destination, err := uc.destinationRepo.FindById(id)
	if err != nil {
		return err
	}

	destination.Name = destinationReq.Name
	destination.Description = destinationReq.Description
	destination.OpenTime = destinationReq.OpenTime
	destination.CloseTime = destinationReq.CloseTime
	destination.EntryPrice = destinationReq.EntryPrice
	destination.Latitude = destinationReq.Latitude
	destination.Longitude = destinationReq.Longitude

	// update category
	category, err := uc.categoryRepo.FindById(destinationReq.CategoryId)
	if err != nil || category == nil {
		return errors.New("category not found")
	}
	destination.CategoryId = destinationReq.CategoryId

	// update facilities
	if err = uc.destinationFacilityRepo.DeleteMany(destination.DestinationFacilities); err != nil {
		return fmt.Errorf("error when delete destination facility: %w", err)
	}

	destinationFacilities := dto.ToDestinationFacilities(destination.Id, destinationReq.FacilityIds)

	for _, facility := range *destinationFacilities {
		if err = uc.destinationFacilityRepo.Update(&facility); err != nil {
			return fmt.Errorf("error when update destination facility: %w", err)
		}
	}

	if err = uc.destinationRepo.Update(destination); err != nil {
		return fmt.Errorf("error when update destination: %w", err)
	}

	// update address
	destinationAddress := dto.ToDestinationAddress(destination.Id, destinationReq.DestinationAddress)
	destinationAddress.Id = destination.DestinationAddress.Id

	// check province
	province, err := uc.provinceRepo.FindById(destinationAddress.ProvinceId)

	if err != nil || province == nil {
		return errors.New("province not found")
	}

	// check city
	city, err := uc.cityRepo.FindById(destinationAddress.CityId)

	if err != nil || city == nil {
		return errors.New("city not found")
	}

	// check subdistrict
	subdistrict, err := uc.subdistrictRepo.FindById(destinationAddress.SubdistrictId)

	if err != nil || subdistrict == nil {
		return errors.New("subdistrict not found")
	}

	if err = uc.destinationAddressRepo.Update(destinationAddress); err != nil {
		return fmt.Errorf("error when update destination address: %w", err)
	}

	for _, image := range destinationReq.DestinationImages {
		urlMedia, errFile := uc.cloudinaryClient.UploadImage(image.File, "destinations")

		if errFile != nil {
			return fmt.Errorf("error when upload image to cloud: %w", errFile)
		}

		if image.Id == uuid.Nil {
			fmt.Println("Create gambar baru")
			destinationMedia := dto.ToDestinationMedia(destination.Id, "image", urlMedia, image.Title)
			if err = uc.destinationMediaRepo.Create(destinationMedia, nil); err != nil {
				return fmt.Errorf("error when create destination media: %w", err)
			}
		} else {
			fmt.Println("Update gambar lama")
			destinationMedia, errDestinationMedia := uc.destinationMediaRepo.FindById(image.Id)
			if errDestinationMedia != nil {
				return fmt.Errorf("error when find destination media: %w", errDestinationMedia)
			}

			destinationMedia.DestinationId = destination.Id
			destinationMedia.Url = urlMedia
			destinationMedia.Title = image.Title

			if err = uc.destinationMediaRepo.Update(destinationMedia); err != nil {
				return fmt.Errorf("error when update destination media: %w", err)
			}
		}
	}

	return nil
}

func (uc *DestinationUsecase) GetAllDestinations(page, limit int, searchQuery string) (*int64, *[]dto.GetAllDestination, error) {
	_, total, destinations, err := uc.destinationRepo.FindAll(page, limit, searchQuery, "", "")
	if err != nil {
		return nil, nil, err
	}

	response := dto.ToGetAllDestinationsResponse(&destinations)

	return total, response, nil
}

func (uc *DestinationUsecase) GetDestinationById(id uuid.UUID) (*dto.GetByIdDestinationResponse, error) {
	destination, err := uc.destinationRepo.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}

	response := dto.ToGetByIdDestinationResponse(destination)

	return response, nil
}

func (uc *DestinationUsecase) DeleteDestination(id uuid.UUID) error {
	destination, err := uc.destinationRepo.FindById(id)
	if err != nil {
		return &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}

	if err = uc.destinationRepo.Delete(destination); err != nil {
		return fmt.Errorf("error when delete destination: %w", err)
	}

	if err = uc.destinationAddressRepo.Delete(destination.DestinationAddress); err != nil {
		return fmt.Errorf("error when delete destination address: %w", err)
	}

	if err = uc.destinationFacilityRepo.DeleteMany(destination.DestinationFacilities); err != nil {
		return fmt.Errorf("error when delete destination facility: %w", err)
	}

	if err = uc.destinationMediaRepo.DeleteMany(&destination.DestinationMedias); err != nil {
		return fmt.Errorf("error when delete destination media: %w", err)
	}

	return nil
}

// increment visit count.
func (uc *DestinationUsecase) IncrementVisitCount(id uuid.UUID) error {
	destination, err := uc.destinationRepo.FindById(id)
	if err != nil {
		return err
	}
	destination.VisitCount++
	return uc.destinationRepo.Update(destination)
}

func (uc *DestinationUsecase) GetCitiesWithDestinations() (*[]dto.Cities, error) {
	cities, err := uc.cityRepo.GetCitiesWithDestinations()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data kota"}
	}

	response := dto.ToCitiesResponse(&cities)

	return response, nil
}

func (uc *DestinationUsecase) GetDestinationByCityId(id string) (*[]dto.DestinationsByCity, error) {
	destinations, err := uc.destinationRepo.FindDestinationByCityId(id)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data destinasi"}
	}
	if len(destinations) == 0 {
		return nil, &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"}
	}
	response := dto.ToDestinationsByCityResponse(&destinations)
	return response, nil
}
