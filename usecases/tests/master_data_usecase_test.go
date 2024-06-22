package tests

import (
	"testing"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetCategoriesMasterData(t *testing.T) {
	mockCategoryRepo := new(repositories.MockCategoryRepository)

	uc := usecases.NewMasterDataUsecase(nil, nil, nil, mockCategoryRepo, nil)

	testCases := []struct {
		name                  string
		mockCategoryRepoSetup func()
		expectedResponse      *[]dto.CategoryResponse
		expectedError         error
	}{
		{
			name: "Success get categories",
			mockCategoryRepoSetup: func() {
				mockCategories := []entities.Category{
					{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4108"), Name: "Category 1"},
					{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4109"), Name: "Category 2"},
				}
				mockCategoryRepo.On("FindAll").Return(mockCategories, nil).Once()
			},
			expectedResponse: &[]dto.CategoryResponse{
				{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4108"), Name: "Category 1"},
				{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4109"), Name: "Category 2"},
			},
			expectedError: nil,
		},
		{
			name: "Error getting categories",
			mockCategoryRepoSetup: func() {
				mockCategoryRepo.On("FindAll").Return(nil, errors.New("some error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kategori"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCategoryRepoSetup()

			res, err := uc.GetCategories()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			}

			mockCategoryRepo.AssertExpectations(t)
		})
	}
}

func TestGetFacilities(t *testing.T) {
	mockFacilityRepo := new(repositories.MockFacilityRepository)

	uc := usecases.NewMasterDataUsecase(nil, nil, nil, nil, mockFacilityRepo)

	testCases := []struct {
		name                  string
		mockFacilityRepoSetup func()
		expectedResponse      *[]dto.FacilityResponse
		expectedError         error
	}{
		{
			name: "Success get facilities",
			mockFacilityRepoSetup: func() {
				mockFacilities := []entities.Facility{
					{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4108"), Name: "Facility 1"},
					{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4109"), Name: "Facility 2"},
				}
				mockFacilityRepo.On("FindAll").Return(mockFacilities, nil).Once()
			},
			expectedResponse: &[]dto.FacilityResponse{
				{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4108"), Name: "Facility 1"},
				{Id: uuid.MustParse("4c5a191e-1a38-4f3e-929b-db312c1c4109"), Name: "Facility 2"},
			},
			expectedError: nil,
		},
		{
			name: "Error getting facilities",
			mockFacilityRepoSetup: func() {
				mockFacilityRepo.On("FindAll").Return(nil, errors.New("some error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan fasilitas"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFacilityRepoSetup()

			res, err := uc.GetFacilities()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			}

			mockFacilityRepo.AssertExpectations(t)
		})
	}
}

func TestGetProvincesMasterData(t *testing.T) {
	mockProvinceRepo := new(repositories.MockProvinceRepository)

	uc := usecases.NewMasterDataUsecase(mockProvinceRepo, nil, nil, nil, nil)

	testCases := []struct {
		name                  string
		mockProvinceRepoSetup func()
		expectedResponse      *[]dto.ProvinceResponse
		expectedError         error
	}{
		{
			name: "Success get provinces",
			mockProvinceRepoSetup: func() {
				mockProvinces := []entities.Province{
					{Id: "ABC", Name: "Province 1"},
					{Id: "BCA", Name: "Province 2"},
				}
				mockProvinceRepo.On("FindAll").Return(mockProvinces, nil).Once()
			},
			expectedResponse: &[]dto.ProvinceResponse{
				{Id: "ABC", Name: "Province 1"},
				{Id: "BCA", Name: "Province 2"},
			},
			expectedError: nil,
		},
		{
			name: "Error getting provinces",
			mockProvinceRepoSetup: func() {
				mockProvinceRepo.On("FindAll").Return(nil, errors.New("some error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan provinsi"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockProvinceRepoSetup()

			res, err := uc.GetProvinces()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			}

			mockProvinceRepo.AssertExpectations(t)
		})
	}
}

func TestGetCities(t *testing.T) {
	mockCityRepo := new(repositories.MockCityRepository)

	uc := usecases.NewMasterDataUsecase(nil, mockCityRepo, nil, nil, nil)

	testCases := []struct {
		name              string
		provinceId        string
		mockCityRepoSetup func()
		expectedResponse  *[]dto.CityResponse
		expectedError     error
	}{
		{
			name:       "Success get cities",
			provinceId: "4c5a191e-1a38-4f3e-929b-db312c1c4108",
			mockCityRepoSetup: func() {
				mockCities := []entities.City{
					{Id: "ABC", Name: "City 1", ProvinceId: "EFG"},
					{Id: "BCA", Name: "City 2", ProvinceId: "HIJ"},
				}
				mockCityRepo.On("FindAll", "4c5a191e-1a38-4f3e-929b-db312c1c4108").Return(mockCities, nil).Once()
			},
			expectedResponse: &[]dto.CityResponse{
				{Id: "ABC", Name: "City 1"},
				{Id: "BCA", Name: "City 2"},
			},
			expectedError: nil,
		},
		{
			name:       "Error getting cities",
			provinceId: "4c5a191e-1a38-4f3e-929b-db312c1c4108",
			mockCityRepoSetup: func() {
				mockCityRepo.On("FindAll", "4c5a191e-1a38-4f3e-929b-db312c1c4108").Return(nil, errors.New("some error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kota"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCityRepoSetup()

			res, err := uc.GetCities(tc.provinceId)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			}

			mockCityRepo.AssertExpectations(t)
		})
	}
}

func TestGetSubdistricts(t *testing.T) {
	mockSubdistrictRepo := new(repositories.MockSubdistrictRepository)

	uc := usecases.NewMasterDataUsecase(nil, nil, mockSubdistrictRepo, nil, nil)

	testCases := []struct {
		name                     string
		cityId                   string
		mockSubdistrictRepoSetup func()
		expectedResponse         *[]dto.SubdistrictResponse
		expectedError            error
	}{
		{
			name:   "Success get subdistricts",
			cityId: "4c5a191e-1a38-4f3e-929b-db312c1c4108",
			mockSubdistrictRepoSetup: func() {
				mockSubdistricts := []entities.Subdistrict{
					{Id: "ABC", Name: "Subdistrict 1", CityId: "FGH"},
					{Id: "BCA", Name: "Subdistrict 2", CityId: "JKL"},
				}
				mockSubdistrictRepo.On("FindAll", "4c5a191e-1a38-4f3e-929b-db312c1c4108").Return(mockSubdistricts, nil).Once()
			},
			expectedResponse: &[]dto.SubdistrictResponse{
				{Id: "ABC", Name: "Subdistrict 1"},
				{Id: "BCA", Name: "Subdistrict 2"},
			},
			expectedError: nil,
		},
		{
			name:   "Error getting subdistricts",
			cityId: "4c5a191e-1a38-4f3e-929b-db312c1c4108",
			mockSubdistrictRepoSetup: func() {
				mockSubdistrictRepo.On("FindAll", "4c5a191e-1a38-4f3e-929b-db312c1c4108").Return(nil, errors.New("some error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kecamatan"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSubdistrictRepoSetup()

			res, err := uc.GetSubdistricts(tc.cityId)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			}

			mockSubdistrictRepo.AssertExpectations(t)
		})
	}
}
