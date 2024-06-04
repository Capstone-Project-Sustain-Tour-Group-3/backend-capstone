package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
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
	var userPersonalizations []*entities.UserPersonalization

	for _, v := range request.ProvinceIds {
		userPersonalizations = append(userPersonalizations, &entities.UserPersonalization{
			Id:         uuid.New(),
			UserId:     userID,
			ProvinceId: v,
			CategoryId: request.CategoryId,
		})
	}

	err := r.userPersonalizationRepo.Create(userPersonalizations)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk membuat personalisasi"}
	}

	return nil
}
