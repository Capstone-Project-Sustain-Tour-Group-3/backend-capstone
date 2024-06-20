package tests

import (
	"errors"
	"testing"

	"capstone/mocks/externals"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreateDestinationMedia(t *testing.T) {
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

	testCases := []struct {
		name          string
		request       dto.CreateDestinationMediaRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success create destination media",
			request: dto.CreateDestinationMediaRequest{
				DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
				Url:           "https://example.com/media.jpg",
				Type:          "image",
				Title:         "Beautiful Landscape",
			},
			mockSetup: func() {
				mockDestinationRepo.On("FindById", mock.Anything).Return(&entities.Destination{Id: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a")}, nil).Once()
				mockDestinationMediaRepo.On("Create", mock.AnythingOfType("*entities.DestinationMedia"), mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Error destination not found",
			request: dto.CreateDestinationMediaRequest{
				DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
				Url:           "https://example.com/media.jpg",
				Type:          "image",
				Title:         "Beautiful Landscape",
			},
			mockSetup: func() {
				mockDestinationRepo.On("FindById", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"},
		},
		{
			name: "Error creating destination media",
			request: dto.CreateDestinationMediaRequest{
				DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
				Url:           "https://example.com/media.jpg",
				Type:          "image",
				Title:         "Beautiful Landscape",
			},
			mockSetup: func() {
				mockDestinationRepo.On("FindById", mock.Anything).Return(&entities.Destination{Id: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a")}, nil).Once()
				mockDestinationMediaRepo.On("Create", mock.AnythingOfType("*entities.DestinationMedia"), mock.Anything).Return(errors.New("create error")).Once()
			},
			expectedError: errors.New("create error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.Create(tc.request)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockDestinationRepo.AssertExpectations(t)
			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}

func TestFindAllDestinationMedia(t *testing.T) {
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

	destinationMedias := []entities.DestinationMedia{
		{
			Id:            uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
			Url:           "https://example.com/media1.jpg",
			Type:          "image",
			Title:         "Beautiful Landscape 1",
		},
		{
			Id:            uuid.MustParse("b2c3d4e5-f678-90ab-cdef-0123456789ab"),
			DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
			Url:           "https://example.com/media2.jpg",
			Type:          "image",
			Title:         "Beautiful Landscape 2",
		},
	}

	testCases := []struct {
		name          string
		page          int
		limit         int
		searchQuery   string
		mockSetup     func()
		expectedTotal *int64
		expectedResp  []dto.GetDetailDestinationMediaResponse
		expectedError error
	}{
		{
			name:        "Success find all destination media",
			page:        1,
			limit:       10,
			searchQuery: "Beautiful",
			mockSetup: func() {
				total := int64(2)
				mockDestinationMediaRepo.On("FindAll", 1, 10, "Beautiful").
					Return(&total, destinationMedias, nil).Once()
			},
			expectedTotal: func() *int64 { i := int64(2); return &i }(),
			expectedResp:  dto.ToGetAllDestinationMediaResponse(destinationMedias),
			expectedError: nil,
		},
		{
			name:        "Error in repository",
			page:        1,
			limit:       10,
			searchQuery: "Beautiful",
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindAll", 1, 10, "Beautiful").
					Return(nil, nil, errors.New("repository error")).Once()
			},
			expectedTotal: nil,
			expectedResp:  nil,
			expectedError: errors.New("repository error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			total, resp, err := uc.FindAll(tc.page, tc.limit, tc.searchQuery)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, total)
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTotal, total)
				require.Equal(t, tc.expectedResp, resp)
			}

			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}

func TestFindByIdDestinationMedia(t *testing.T) {
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

	destinationMedia := entities.DestinationMedia{
		Id:            uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
		DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
		Url:           "https://example.com/media1.jpg",
		Type:          "image",
		Title:         "Beautiful Landscape 1",
	}
	toRes := dto.ToGetDetailDestinationMediaResponse(destinationMedia)
	testCases := []struct {
		name          string
		id            uuid.UUID
		mockSetup     func()
		expectedResp  *dto.GetDetailDestinationMediaResponse
		expectedError error
	}{
		{
			name: "Success find destination media by id",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&destinationMedia, nil).Once()
			},
			expectedResp:  &toRes,
			expectedError: nil,
		},
		{
			name: "Error destination media not found",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.NotFoundError{Message: "Konten media destinasi tidak ditemukan"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			resp, err := uc.FindById(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tc.expectedResp, resp)
			}

			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateDestinationMedia(t *testing.T) {
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

	existingMedia := entities.DestinationMedia{
		Id:            uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
		DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
		Url:           "https://example.com/media1.jpg",
		Type:          "image",
		Title:         "Beautiful Landscape 1",
	}

	updateRequest := dto.UpdateDestinationMediaRequest{
		DestinationId: uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
		Url:           "https://example.com/media2.jpg",
		Type:          "video",
		Title:         "Amazing View",
	}

	updatedMedia := existingMedia
	updatedMedia.DestinationId = updateRequest.DestinationId
	updatedMedia.Url = updateRequest.Url
	updatedMedia.Type = updateRequest.Type
	updatedMedia.Title = updateRequest.Title

	testCases := []struct {
		name          string
		id            uuid.UUID
		request       dto.UpdateDestinationMediaRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name:    "Success update destination media",
			id:      uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			request: updateRequest,
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&existingMedia, nil).Once()
				mockDestinationMediaRepo.On("Update", &updatedMedia).
					Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name:    "Error destination media not found",
			id:      uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			request: updateRequest,
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"},
		},
		{
			name:    "Error updating destination media",
			id:      uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			request: updateRequest,
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&existingMedia, nil).Once()
				mockDestinationMediaRepo.On("Update", &updatedMedia).
					Return(errors.New("update error")).Once()
			},
			expectedError: errors.New("update error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.Update(tc.id, tc.request)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteDestinationMedia(t *testing.T) {
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

	existingMedia := entities.DestinationMedia{
		Id:            uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
		DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
		Url:           "https://example.com/media1.jpg",
		Type:          "image",
		Title:         "Beautiful Landscape 1",
	}

	testCases := []struct {
		name          string
		id            uuid.UUID
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success delete destination media",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&existingMedia, nil).Once()
				mockDestinationMediaRepo.On("Delete", &existingMedia).
					Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Error destination media not found",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"},
		},
		{
			name: "Error deleting destination media",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func() {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&existingMedia, nil).Once()
				mockDestinationMediaRepo.On("Delete", &existingMedia).
					Return(errors.New("delete error")).Once()
			},
			expectedError: errors.New("delete error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.Delete(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}
