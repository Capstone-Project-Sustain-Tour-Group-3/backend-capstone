package tests

import (
	"errors"
	"mime/multipart"
	"testing"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/db"
	"capstone/mocks/externals"
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
		mockSetup     func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository, mockDestinationRepo *repositories.MockDestinationRepository, mockCloudinaryClient *externals.MockCloudinaryClient)
		expectedError error
	}{
		{
			name: "Success delete destination media",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository, mockDestinationRepo *repositories.MockDestinationRepository, mockCloudinaryClient *externals.MockCloudinaryClient) {
				tx := db.MockDB()
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&existingMedia, nil)
				mockDestinationMediaRepo.On("BeginTx").Return(tx)
				mockDestinationMediaRepo.On("Delete", &existingMedia, tx).
					Return(nil)
				mockCloudinaryClient.On("DeleteImage", mock.Anything).Return(nil)
				mockDestinationRepo.On("CommitTx", tx).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error destination media not found",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository, mockDestinationRepo *repositories.MockDestinationRepository, mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
				tx := db.MockDB()
				mockDestinationMediaRepo.On("BeginTx").Return(tx)
				mockDestinationMediaRepo.On("Delete", &existingMedia, tx).
					Return(nil)

				mockCloudinaryClient.On("DeleteImage", mock.Anything).Return(nil)
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"},
		},
		{
			name: "Error deleting destination media",
			id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			mockSetup: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository, mockDestinationRepo *repositories.MockDestinationRepository, mockCloudinaryClient *externals.MockCloudinaryClient) {
				tx := db.MockDB()
				mockDestinationMediaRepo.On("FindById", mock.Anything).
					Return(&existingMedia, nil)
				mockDestinationMediaRepo.On("BeginTx").Return(tx)
				mockDestinationMediaRepo.On("Delete", &existingMedia, tx).
					Return(errors.New("delete error"))
				mockCloudinaryClient.On("DeleteImage", mock.Anything).Return(nil)
				mockDestinationRepo.On("CommitTx", tx).Return(nil)
			},
			expectedError: errors.New("delete error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
			mockDestinationRepo := new(repositories.MockDestinationRepository)
			mockCloudinaryClient := new(externals.MockCloudinaryClient)

			uc := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

			tc.mockSetup(mockDestinationMediaRepo, mockDestinationRepo, mockCloudinaryClient)

			err := uc.Delete(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			// mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}

func TestUploadImage(t *testing.T) {
	tests := []struct {
		name                string
		mockCloudinarySetup func(*externals.MockCloudinaryClient)
		mockRepoSetup       func(*repositories.MockDestinationMediaRepository)
		expectedURL         string
		expectedError       string
	}{
		{
			name: "success",
			mockCloudinarySetup: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("http://example.com/image.jpg", nil)
			},
			mockRepoSetup: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("Create", mock.AnythingOfType("*entities.DestinationMedia"), (*gorm.DB)(nil)).Return(nil)
			},
			expectedURL:   "http://example.com/image.jpg",
			expectedError: "",
		},
		{
			name: "upload image error",
			mockCloudinarySetup: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("", errors.New("upload error"))
			},
			mockRepoSetup: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {},
			expectedURL:   "",
			expectedError: "upload error",
		},
		{
			name: "create media error",
			mockCloudinarySetup: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("http://example.com/image.jpg", nil)
			},
			mockRepoSetup: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("Create", mock.AnythingOfType("*entities.DestinationMedia"), (*gorm.DB)(nil)).Return(errors.New("create error"))
			},
			expectedURL:   "",
			expectedError: "create error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCloudinaryClient := new(externals.MockCloudinaryClient)
			mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
			usecase := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, nil, mockCloudinaryClient)

			file := new(multipart.FileHeader)
			open, _ := file.Open()
			destinationId := uuid.New()
			title := "Test Title"
			uploadRequest := dto.UploadDestinationMediaRequest{
				DestinationId: destinationId,
				File:          open,
				Title:         title,
			}

			tt.mockCloudinarySetup(mockCloudinaryClient)
			tt.mockRepoSetup(mockDestinationMediaRepo)

			url, err := usecase.UploadImage(uploadRequest)

			if tt.expectedError == "" {
				require.NoError(t, err)
				require.Equal(t, tt.expectedURL, url)
			} else {
				require.Error(t, err)
				require.Empty(t, url)
				require.Equal(t, tt.expectedError, err.Error())
			}

			mockCloudinaryClient.AssertExpectations(t)
			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateImage(t *testing.T) {
	tests := []struct {
		name                 string
		mockRepoFindById     func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository)
		mockCloudinaryUpload func(mockCloudinaryClient *externals.MockCloudinaryClient)
		mockRepoUpdate       func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository)
		mockCloudinaryDelete func(mockCloudinaryClient *externals.MockCloudinaryClient)
		expectedURL          string
		expectedError        string
	}{
		{
			name: "success",
			mockRepoFindById: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("FindById", mock.Anything).Return(&entities.DestinationMedia{Url: "http://example.com/oldimage.jpg"}, nil)
			},
			mockCloudinaryUpload: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("http://example.com/newimage.jpg", nil)
			},
			mockRepoUpdate: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("Update", mock.AnythingOfType("*entities.DestinationMedia")).Return(nil)
			},
			mockCloudinaryDelete: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("DeleteImage", "http://example.com/oldimage.jpg").Return(nil)
			},
			expectedURL:   "http://example.com/newimage.jpg",
			expectedError: "",
		},
		{
			name: "find by id error",
			mockRepoFindById: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("FindById", mock.Anything).Return(nil, errors.New("not found"))
			},
			mockCloudinaryUpload: func(mockCloudinaryClient *externals.MockCloudinaryClient) {},
			mockRepoUpdate:       func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {},
			mockCloudinaryDelete: func(mockCloudinaryClient *externals.MockCloudinaryClient) {},
			expectedURL:          "",
			expectedError:        "Destinasi tidak ditemukan",
		},
		{
			name: "upload image error",
			mockRepoFindById: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("FindById", mock.Anything).Return(&entities.DestinationMedia{Url: "http://example.com/oldimage.jpg"}, nil)
			},
			mockCloudinaryUpload: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("", errors.New("upload error"))
			},
			mockRepoUpdate:       func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {},
			mockCloudinaryDelete: func(mockCloudinaryClient *externals.MockCloudinaryClient) {},
			expectedURL:          "",
			expectedError:        "upload error",
		},
		{
			name: "update media error",
			mockRepoFindById: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("FindById", mock.Anything).Return(&entities.DestinationMedia{Url: "http://example.com/oldimage.jpg"}, nil)
			},
			mockCloudinaryUpload: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("http://example.com/newimage.jpg", nil)
			},
			mockRepoUpdate: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("Update", mock.AnythingOfType("*entities.DestinationMedia")).Return(errors.New("update error"))
			},
			mockCloudinaryDelete: func(mockCloudinaryClient *externals.MockCloudinaryClient) {},
			expectedURL:          "",
			expectedError:        "update error",
		},
		{
			name: "delete image error",
			mockRepoFindById: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("FindById", mock.Anything).Return(&entities.DestinationMedia{Url: "http://example.com/oldimage.jpg"}, nil)
			},
			mockCloudinaryUpload: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("UploadImage", mock.Anything, "destination_media").Return("http://example.com/newimage.jpg", nil)
			},
			mockRepoUpdate: func(mockDestinationMediaRepo *repositories.MockDestinationMediaRepository) {
				mockDestinationMediaRepo.On("Update", mock.AnythingOfType("*entities.DestinationMedia")).Return(nil)
			},
			mockCloudinaryDelete: func(mockCloudinaryClient *externals.MockCloudinaryClient) {
				mockCloudinaryClient.On("DeleteImage", "http://example.com/oldimage.jpg").Return(errors.New("delete error"))
			},
			expectedURL:   "",
			expectedError: "Gagal menghapus media destinasi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCloudinaryClient := new(externals.MockCloudinaryClient)
			mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
			usecase := usecases.NewDestinationMediaUsecase(mockDestinationMediaRepo, nil, mockCloudinaryClient)

			id := uuid.New()
			file := new(multipart.FileHeader)
			open, _ := file.Open()
			title := "Updated Title"
			updateRequest := dto.UpdateImageDestinationMediaRequest{
				File:  open,
				Title: title,
			}

			tt.mockRepoFindById(mockDestinationMediaRepo)
			tt.mockCloudinaryUpload(mockCloudinaryClient)
			tt.mockRepoUpdate(mockDestinationMediaRepo)
			tt.mockCloudinaryDelete(mockCloudinaryClient)

			url, err := usecase.UpdateImage(id, updateRequest)

			if tt.expectedError == "" {
				require.NoError(t, err)
				require.Equal(t, tt.expectedURL, url)
			} else {
				require.Error(t, err)
				require.Empty(t, url)
				require.Equal(t, tt.expectedError, err.Error())
			}

			mockCloudinaryClient.AssertExpectations(t)
			mockDestinationMediaRepo.AssertExpectations(t)
		})
	}
}
