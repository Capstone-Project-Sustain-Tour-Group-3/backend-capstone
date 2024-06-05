package usecases

import (
	"errors"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalizationUsecase interface {
	GetProvinces() (*[]dto.ProvinceResponse, error)
	GetCategories() (*[]dto.CategoryResponse, error)
	CreatePersonalization(request *dto.PersonalizationRequest, userID uuid.UUID) error
}

type personalizationUsecase struct {
	provinceRepo            repositories.IProvinceRepository
	categoryRepo            repositories.ICategoryRepository
	userPersonalizationRepo repositories.UserPersonalizationRepository
}

func NewPersonalizationUsecase(
	provinceRepo repositories.IProvinceRepository,
	categoryRepo repositories.ICategoryRepository,
	userPersonalizationRepo repositories.UserPersonalizationRepository,
) *personalizationUsecase {
	return &personalizationUsecase{
		provinceRepo:            provinceRepo,
		categoryRepo:            categoryRepo,
		userPersonalizationRepo: userPersonalizationRepo,
	}
}

func (r *personalizationUsecase) GetProvinces() (*[]dto.ProvinceResponse, error) {
	provinces, err := r.provinceRepo.FindAll()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan provinsi"}
	}

	res := new([]dto.ProvinceResponse)

	err = mapper.MapperSlice(provinces, res)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan provinsi"}
	}

	return res, nil
}

func (r *personalizationUsecase) GetCategories() (*[]dto.CategoryResponse, error) {
	categories, err := r.categoryRepo.FindAll()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kategori destinasi"}
	}

	res := new([]dto.CategoryResponse)

	err = mapper.MapperSlice(categories, res)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kategori destinasi"}
	}

	return res, nil
}

func (r *personalizationUsecase) CreatePersonalization(request *dto.PersonalizationRequest, userID uuid.UUID) error {
	id := uuid.New()

	var persCategories []entities.PersonalizationCategory
	var persProvinces []entities.PersonalizationProvince

	for _, category := range request.CategoryIds {
		persCategories = append(
			persCategories,
			entities.PersonalizationCategory{
				Id:         uuid.New(),
				CategoryId: category,
			},
		)
	}

	for _, province := range request.ProvinceIds {
		persProvinces = append(
			persProvinces,
			entities.PersonalizationProvince{
				Id:         uuid.New(),
				ProvinceId: province,
			},
		)
	}

	userPersonalization := &entities.UserPersonalization{
		Id:                        id,
		UserId:                    userID,
		PersonalizationCategories: persCategories,
		PersonalizationProvinces:  persProvinces,
	}

	err := r.userPersonalizationRepo.Create(userPersonalization)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return &errorHandlers.ConflictError{Message: "Personalisasi sudah dibuat"}
		default:
			return &errorHandlers.InternalServerError{Message: "Gagal untuk membuat personalisasi"}
		}
	}

	return nil
}
