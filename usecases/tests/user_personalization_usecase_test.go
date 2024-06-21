package tests

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/repositories"
	"capstone/usecases"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestGetProvinces(t *testing.T) {
	provinceList := []dto.ProvinceResponse{
		{Id: "ABC", Name: "Province 1"},
		{Id: "BCA", Name: "Province 2"},
	}

	testCases := []struct {
		name           string
		mockSetup      func(provinceRepo *repositories.MockProvinceRepository)
		expectedResult *[]dto.ProvinceResponse
		expectedError  error
	}{
		{
			name: "Success get all provinces",
			mockSetup: func(provinceRepo *repositories.MockProvinceRepository) {
				provinces := []entities.Province{
					{Id: "ABC", Name: "Province 1"},
					{Id: "BCA", Name: "Province 2"},
				}
				provinceRepo.On("FindAll").Return(provinces, nil)
			},
			expectedResult: &provinceList,
			expectedError:  nil,
		},
		{
			name: "Failed to get provinces - repository error",
			mockSetup: func(provinceRepo *repositories.MockProvinceRepository) {
				provinceRepo.On("FindAll").Return(nil, errors.New("repository error"))
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan provinsi"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			provinceRepo := new(repositories.MockProvinceRepository)
			uc := usecases.NewPersonalizationUsecase(provinceRepo, nil, nil)

			tc.mockSetup(provinceRepo)

			result, err := uc.GetProvinces()

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, result)
			}

			provinceRepo.AssertExpectations(t)
		})
	}
}

func TestGetCategories(t *testing.T) {
	categoryList := []dto.CategoryResponse{
		{Id: uuid.New(), Name: "Category 1"},
		{Id: uuid.New(), Name: "Category 2"},
	}

	testCases := []struct {
		name           string
		mockSetup      func(categoryRepo *repositories.MockCategoryRepository)
		expectedResult *[]dto.CategoryResponse
		expectedError  error
	}{
		{
			name: "Success get all categories",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository) {
				categories := []entities.Category{
					{Id: uuid.New(), Name: "Category 1"},
					{Id: uuid.New(), Name: "Category 2"},
				}
				categoryRepo.On("FindAll").Return(categories, nil)
			},
			expectedResult: &categoryList,
			expectedError:  nil,
		},
		{
			name: "Failed to get categories - repository error",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository) {
				categoryRepo.On("FindAll").Return(nil, errors.New("repository error"))
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Gagal untuk mendapatkan kategori destinasi"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			categoryRepo := new(repositories.MockCategoryRepository)
			uc := usecases.NewPersonalizationUsecase(nil, categoryRepo, nil)

			tc.mockSetup(categoryRepo)

			result, err := uc.GetCategories()

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, result)
			}

			categoryRepo.AssertExpectations(t)
		})
	}
}

func TestCreatePersonalization(t *testing.T) {
	testUserID := uuid.New()
	testRequest := &dto.PersonalizationRequest{
		CategoryIds: []uuid.UUID{uuid.New(), uuid.New()},
		ProvinceIds: []string{"ABC", "BCA"},
	}

	testCases := []struct {
		name          string
		mockSetup     func(categoryRepo *repositories.MockCategoryRepository, provinceRepo *repositories.MockProvinceRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository)
		expectedError error
	}{
		{
			name: "Success create personalization",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository, provinceRepo *repositories.MockProvinceRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository) {
				for _, id := range testRequest.CategoryIds {
					categoryRepo.On("FindById", id).Return(&entities.Category{Id: id}, nil)
				}
				for _, id := range testRequest.ProvinceIds {
					provinceRepo.On("FindById", id).Return(&entities.Province{Id: id}, nil)
				}
				userPersonalizationRepo.On("Create", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Category not found",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository, provinceRepo *repositories.MockProvinceRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository) {
				categoryRepo.On("FindById", mock.Anything).Return(nil, &errorHandlers.NotFoundError{Message: "Kategori destinasi pilihan ke-1 tidak ditemukan"})
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Kategori destinasi pilihan ke-1 tidak ditemukan"},
		},
		{
			name: "Province not found",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository, provinceRepo *repositories.MockProvinceRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository) {
				categoryRepo.On("FindById", mock.Anything).Return(&entities.Category{}, nil)
				provinceRepo.On("FindById", mock.Anything).Return(nil, &errorHandlers.NotFoundError{Message: "Provinsi pilihan ke-1 tidak ditemukan"})
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Provinsi pilihan ke-1 tidak ditemukan"},
		},
		{
			name: "Personalization already exists",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository, provinceRepo *repositories.MockProvinceRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository) {
				for _, id := range testRequest.CategoryIds {
					categoryRepo.On("FindById", id).Return(&entities.Category{Id: id}, nil)
				}
				for _, id := range testRequest.ProvinceIds {
					provinceRepo.On("FindById", id).Return(&entities.Province{Id: id}, nil)
				}
				userPersonalizationRepo.On("Create", mock.Anything).Return(gorm.ErrDuplicatedKey)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Personalisasi sudah dibuat"},
		},
		{
			name: "Internal server error",
			mockSetup: func(categoryRepo *repositories.MockCategoryRepository, provinceRepo *repositories.MockProvinceRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository) {
				for _, id := range testRequest.CategoryIds {
					categoryRepo.On("FindById", id).Return(&entities.Category{Id: id}, nil)
				}
				for _, id := range testRequest.ProvinceIds {
					provinceRepo.On("FindById", id).Return(&entities.Province{Id: id}, nil)
				}
				userPersonalizationRepo.On("Create", mock.Anything).Return(errors.New("internal error"))
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk membuat personalisasi"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			categoryRepo := new(repositories.MockCategoryRepository)
			provinceRepo := new(repositories.MockProvinceRepository)
			userPersonalizationRepo := new(repositories.MockUserPersonalizationRepository)

			uc := usecases.NewPersonalizationUsecase(provinceRepo, categoryRepo, userPersonalizationRepo)

			tc.mockSetup(categoryRepo, provinceRepo, userPersonalizationRepo)

			err := uc.CreatePersonalization(testRequest, testUserID)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			categoryRepo.AssertExpectations(t)
			provinceRepo.AssertExpectations(t)
			userPersonalizationRepo.AssertExpectations(t)
		})
	}
}
