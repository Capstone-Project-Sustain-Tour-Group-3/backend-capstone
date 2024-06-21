package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/devfeel/mapper"
)

type IMasterDataUsecase interface {
	GetCategories() (*[]dto.CategoryResponse, error)
	GetFacilities() (*[]dto.FacilityResponse, error)
	GetProvinces() (*[]dto.ProvinceResponse, error)
	GetCities(provinceId string) (*[]dto.CityResponse, error)
	GetSubdistricts(cityId string) (*[]dto.SubdistrictResponse, error)
}

type MasterDataUsecase struct {
	provinceRepo    repositories.IProvinceRepository
	cityRepo        repositories.ICityRepository
	subdistrictRepo repositories.ISubdistrictRepository
	categoryRepo    repositories.ICategoryRepository
	facilityRepo    repositories.IFacilityRepository
}

func NewMasterDataUsecase(
	provinceRepo repositories.IProvinceRepository,
	cityRepo repositories.ICityRepository,
	subdistrictRepo repositories.ISubdistrictRepository,
	categoryRepo repositories.ICategoryRepository,
	facilityRepo repositories.IFacilityRepository,
) *MasterDataUsecase {
	return &MasterDataUsecase{
		provinceRepo:    provinceRepo,
		cityRepo:        cityRepo,
		subdistrictRepo: subdistrictRepo,
		categoryRepo:    categoryRepo,
		facilityRepo:    facilityRepo,
	}
}

func (r *MasterDataUsecase) GetCategories() (*[]dto.CategoryResponse, error) {
	categories, err := r.categoryRepo.FindAll()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kategori"}
	}

	res := new([]dto.CategoryResponse)

	err = mapper.MapperSlice(categories, res)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kategori"}
	}

	return res, nil
}

func (r *MasterDataUsecase) GetFacilities() (*[]dto.FacilityResponse, error) {
	facilities, err := r.facilityRepo.FindAll()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan fasilitas"}
	}

	res := new([]dto.FacilityResponse)

	err = mapper.MapperSlice(facilities, res)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan fasilitas"}
	}

	return res, nil
}

func (r *MasterDataUsecase) GetProvinces() (*[]dto.ProvinceResponse, error) {
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

func (r *MasterDataUsecase) GetCities(provinceId string) (*[]dto.CityResponse, error) {
	cities, err := r.cityRepo.FindAll(provinceId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kota"}
	}

	res := new([]dto.CityResponse)

	err = mapper.MapperSlice(cities, res)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kota"}
	}

	return res, nil
}

func (r *MasterDataUsecase) GetSubdistricts(cityId string) (*[]dto.SubdistrictResponse, error) {
	subdistricts, err := r.subdistrictRepo.FindAll(cityId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kecamatan"}
	}

	res := new([]dto.SubdistrictResponse)

	err = mapper.MapperSlice(subdistricts, res)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kecamatan"}
	}

	return res, nil
}
