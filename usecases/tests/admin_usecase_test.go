package tests

import (
	"errors"
	"testing"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/externals"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestLoginAdmin(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	mockPasswordHelper := new(externals.MockPasswordHelper)
	mockTokenHelper := new(externals.MockTokenHelper)

	uc := usecases.NewAdminUsecase(mockRepo, nil, mockPasswordHelper, mockTokenHelper)

	testCases := []struct {
		name           string
		request        *dto.LoginAdminRequest
		mockSetup      func()
		expectedResult *dto.LoginAdminResponse
		expectedError  error
	}{
		{
			name: "Success login",
			request: &dto.LoginAdminRequest{
				Username: "testuser",
				Password: "password123",
			},
			mockSetup: func() {
				admin := &entities.Admin{
					Username:        "testuser",
					Password:        "hashedPassword",
					ProfileImageURL: nil,
					Role:            "admin",
				}
				mockRepo.On("FindByUsername", "testuser").Return(admin, nil).Once()
				mockPasswordHelper.On("VerifyPassword", "hashedPassword", "password123").Return(nil).Once()
				mockTokenHelper.On("GenerateAccessToken", admin).Return("accessToken", nil).Once()
				mockTokenHelper.On("GenerateRefreshToken", admin).Return("refreshToken", nil).Once()
				admin.RefreshToken = "refreshToken"
				mockRepo.On("Update", admin).Return(nil).Once()
			},
			expectedResult: &dto.LoginAdminResponse{
				Username:     "testuser",
				ProfileImage: nil,
				Role:         "admin",
				AccessToken:  "accessToken",
				RefreshToken: "refreshToken",
			},
			expectedError: nil,
		},
		{
			name: "User not found",
			request: &dto.LoginAdminRequest{
				Username: "unknownuser",
				Password: "password123",
			},
			mockSetup: func() {
				mockRepo.On("FindByUsername", "unknownuser").Return(nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"},
		},
		{
			name: "Incorrect password",
			request: &dto.LoginAdminRequest{
				Username: "testuser",
				Password: "wrongpassword",
			},
			mockSetup: func() {
				admin := &entities.Admin{
					Username: "testuser",
					Password: "hashedPassword",
				}
				mockRepo.On("FindByUsername", "testuser").Return(admin, nil).Once()
				mockPasswordHelper.On("VerifyPassword", "hashedPassword", "wrongpassword").Return(&errorHandlers.ConflictError{Message: "Email atau password salah"}).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.ConflictError{Message: "Email atau password salah"},
		},
		{
			name: "Access token generation failure",
			request: &dto.LoginAdminRequest{
				Username: "testuser",
				Password: "password123",
			},
			mockSetup: func() {
				admin := &entities.Admin{
					Username: "testuser",
					Password: "hashedPassword",
				}
				mockRepo.On("FindByUsername", "testuser").Return(admin, nil).Once()
				mockPasswordHelper.On("VerifyPassword", "hashedPassword", "password123").Return(nil).Once()
				mockTokenHelper.On("GenerateAccessToken", admin).Return("", &errorHandlers.InternalServerError{Message: "Token generation error"}).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Token generation error"},
		},
		{
			name: "Refresh token generation failure",
			request: &dto.LoginAdminRequest{
				Username: "testuser",
				Password: "password123",
			},
			mockSetup: func() {
				admin := &entities.Admin{
					Username: "testuser",
					Password: "hashedPassword",
				}
				mockRepo.On("FindByUsername", "testuser").Return(admin, nil).Once()
				mockPasswordHelper.On("VerifyPassword", "hashedPassword", "password123").Return(nil).Once()
				mockTokenHelper.On("GenerateAccessToken", admin).Return("accessToken", nil).Once()
				mockTokenHelper.On("GenerateRefreshToken", admin).Return("", &errorHandlers.InternalServerError{Message: "Token generation error"}).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Token generation error"},
		},
		{
			name: "Update failure",
			request: &dto.LoginAdminRequest{
				Username: "testuser",
				Password: "password123",
			},
			mockSetup: func() {
				admin := &entities.Admin{
					Username: "testuser",
					Password: "hashedPassword",
				}
				mockRepo.On("FindByUsername", "testuser").Return(admin, nil).Once()
				mockPasswordHelper.On("VerifyPassword", "hashedPassword", "password123").Return(nil).Once()
				mockTokenHelper.On("GenerateAccessToken", admin).Return("accessToken", nil).Once()
				mockTokenHelper.On("GenerateRefreshToken", admin).Return("refreshToken", nil).Once()
				admin.RefreshToken = "refreshToken"
				mockRepo.On("Update", admin).Return(&errorHandlers.InternalServerError{Message: "Update error"}).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Update error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			result, err := uc.Login(tc.request)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResult, result)
			}

			mockRepo.AssertExpectations(t)
			mockPasswordHelper.AssertExpectations(t)
			mockTokenHelper.AssertExpectations(t)
		})
	}
}

func TestLogoutAdmin(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	uc := usecases.NewAdminUsecase(mockRepo, nil, nil, nil)

	testCases := []struct {
		name          string
		token         string
		mockSetup     func()
		expectedError error
	}{
		{
			name:  "Success logout",
			token: "validRefreshToken",
			mockSetup: func() {
				admin := &entities.Admin{
					RefreshToken: "validRefreshToken",
				}
				mockRepo.On("GetUserByRefreshToken", "validRefreshToken").Return(admin, nil).Once()
				admin.RefreshToken = ""
				mockRepo.On("Update", admin).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name:  "Invalid token",
			token: "invalidRefreshToken",
			mockSetup: func() {
				mockRepo.On("GetUserByRefreshToken", "invalidRefreshToken").Return(nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}).Once()
			},
			expectedError: &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"},
		},
		{
			name:  "Update failure",
			token: "validRefreshToken",
			mockSetup: func() {
				admin := &entities.Admin{
					RefreshToken: "validRefreshToken",
				}
				mockRepo.On("GetUserByRefreshToken", "validRefreshToken").Return(admin, nil).Once()
				admin.RefreshToken = ""
				mockRepo.On("Update", admin).Return(&errorHandlers.InternalServerError{Message: "Update error"}).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Update error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.Logout(tc.token)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetNewAccessTokenAdmin(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	mockTokenHelper := new(externals.MockTokenHelper)
	uc := usecases.NewAdminUsecase(mockRepo, nil, nil, mockTokenHelper)

	testCases := []struct {
		name          string
		refreshToken  string
		mockSetup     func()
		expectedToken *dto.NewToken
		expectedError error
	}{
		{
			name:         "Success get new access token",
			refreshToken: "validRefreshToken",
			mockSetup: func() {
				admin := &entities.Admin{
					RefreshToken: "validRefreshToken",
				}
				mockRepo.On("GetUserByRefreshToken", "validRefreshToken").Return(admin, nil).Once()
				mockTokenHelper.On("GenerateAccessToken", admin).Return("newAccessToken", nil).Once()
			},
			expectedToken: &dto.NewToken{
				AccessToken: "newAccessToken",
			},
			expectedError: nil,
		},
		{
			name:         "Invalid refresh token",
			refreshToken: "invalidRefreshToken",
			mockSetup: func() {
				mockRepo.On("GetUserByRefreshToken", "invalidRefreshToken").Return(nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}).Once()
			},
			expectedToken: nil,
			expectedError: &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"},
		},
		{
			name:         "Refresh token mismatch",
			refreshToken: "mismatchedRefreshToken",
			mockSetup: func() {
				admin := &entities.Admin{
					RefreshToken: "validRefreshToken",
				}
				mockRepo.On("GetUserByRefreshToken", "mismatchedRefreshToken").Return(admin, nil).Once()
			},
			expectedToken: nil,
			expectedError: &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"},
		},
		{
			name:         "Generate access token failure",
			refreshToken: "validRefreshToken",
			mockSetup: func() {
				admin := &entities.Admin{
					RefreshToken: "validRefreshToken",
				}
				mockRepo.On("GetUserByRefreshToken", "validRefreshToken").Return(admin, nil).Once()
				mockTokenHelper.On("GenerateAccessToken", admin).Return("", &errorHandlers.InternalServerError{Message: "Token generation error"}).Once()
			},
			expectedToken: nil,
			expectedError: &errorHandlers.InternalServerError{Message: "Token generation error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			token, err := uc.GetNewAccessToken(tc.refreshToken)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToken, token)
			}

			mockRepo.AssertExpectations(t)
			mockTokenHelper.AssertExpectations(t)
		})
	}
}

func TestGetAllAdmins(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	uc := usecases.NewAdminUsecase(mockRepo, nil, nil, nil)

	testCases := []struct {
		name           string
		page           int
		limit          int
		search         string
		mockSetup      func()
		expectedAdmins *[]dto.GetAllAdminResponse
		expectedTotal  *int64
		expectedError  error
	}{
		{
			name:   "Success get all admins",
			page:   1,
			limit:  10,
			search: "admin",
			mockSetup: func() {
				admins := []entities.Admin{
					{Username: "admin1", Role: "superadmin"},
					{Username: "admin2", Role: "admin"},
				}
				total := int64(2)
				mockRepo.On("FindAll", 0, 10, "admin").Return(&admins, &total, nil).Once()
			},
			expectedAdmins: &[]dto.GetAllAdminResponse{
				{Username: "admin1", CreatedAt: "01-01-0001"},
				{Username: "admin2", CreatedAt: "01-01-0001"},
			},
			expectedTotal: func() *int64 { t := int64(2); return &t }(),
			expectedError: nil,
		},
		{
			name:   "Error getting admins",
			page:   1,
			limit:  10,
			search: "admin",
			mockSetup: func() {
				mockRepo.On("FindAll", 0, 10, "admin").Return(nil, nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data admin"}).Once()
			},
			expectedAdmins: nil,
			expectedTotal:  nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data admin"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			admins, total, err := uc.GetAllAdmins(tc.page, tc.limit, tc.search)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedAdmins, admins)
				require.Equal(t, tc.expectedTotal, total)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAdminDetail(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	uc := usecases.NewAdminUsecase(mockRepo, nil, nil, nil)

	testCases := []struct {
		name           string
		id             uuid.UUID
		mockSetup      func()
		expectedResult *dto.GetDetailAdminResponse
		expectedError  error
	}{
		{
			name: "Success get admin detail",
			id:   uuid.New(),
			mockSetup: func() {
				admin := entities.Admin{
					Id:       uuid.New(),
					Username: "admin1",
					Role:     "superadmin",
				}
				mockRepo.On("FindById", mock.Anything).Return(&admin, nil).Once()
			},
			expectedResult: &dto.GetDetailAdminResponse{
				Username:        "admin1",
				ProfileImageURL: nil,
				CreatedAt:       "01-01-0001",
			},
			expectedError: nil,
		},
		{
			name: "Admin not found",
			id:   uuid.New(),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.NotFoundError{Message: "Data admin tidak ditemukan"},
		},
		{
			name: "Internal server error",
			id:   uuid.New(),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(nil, errors.New("some internal error")).Once()
			},
			expectedResult: nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Gagal mendapatkan detail admin"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			result, err := uc.GetAdminDetail(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResult, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestCreateAdmin(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	mockPasswordHelper := new(externals.MockPasswordHelper)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)
	uc := usecases.NewAdminUsecase(mockRepo, mockCloudinaryClient, mockPasswordHelper, nil)

	testCases := []struct {
		name          string
		request       *dto.AdminRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success create admin",
			request: &dto.AdminRequest{
				Username:     "admin1",
				Password:     "password123",
				ProfileImage: nil,
			},
			mockSetup: func() {
				// mockCloudinaryClient.On("UploadImage", mock.Anything, "admin").Return("http://image.url/profile.png", nil).Once()
				mockPasswordHelper.On("HashPassword", "password123").Return("hashedPassword123", nil).Once()
				mockRepo.On("Create", mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Error on password hashing",
			request: &dto.AdminRequest{
				Username:     "admin1",
				Password:     "password123",
				ProfileImage: nil,
			},
			mockSetup: func() {
				mockPasswordHelper.On("HashPassword", "password123").Return("", errors.New("hashing error")).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menambah data admin"},
		},
		{
			name: "Error on duplicate username",
			request: &dto.AdminRequest{
				Username:     "admin1",
				Password:     "password123",
				ProfileImage: nil,
			},
			mockSetup: func() {
				mockPasswordHelper.On("HashPassword", "password123").Return("hashedPassword123", nil).Once()
				mockRepo.On("Create", mock.Anything).Return(gorm.ErrDuplicatedKey).Once()
			},
			expectedError: &errorHandlers.ConflictError{Message: "Username sudah digunakan"},
		},
		{
			name: "Internal server error on create",
			request: &dto.AdminRequest{
				Username:     "admin1",
				Password:     "password123",
				ProfileImage: nil,
			},
			mockSetup: func() {
				mockPasswordHelper.On("HashPassword", "password123").Return("hashedPassword123", nil).Once()
				mockRepo.On("Create", mock.Anything).Return(errors.New("some internal error")).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menambah data admin"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.CreateAdmin(tc.request)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockPasswordHelper.AssertExpectations(t)
			mockCloudinaryClient.AssertExpectations(t)
		})
	}
}

func TestUpdateAdmin(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	mockPasswordHelper := new(externals.MockPasswordHelper)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)
	uc := usecases.NewAdminUsecase(mockRepo, mockCloudinaryClient, mockPasswordHelper, nil)

	admin := &entities.Admin{
		Id:              uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
		Username:        "existing_admin",
		Password:        "hashed_password",
		ProfileImageURL: nil,
	}

	testCases := []struct {
		name          string
		request       *dto.UpdateAdminRequest
		adminID       uuid.UUID
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success update admin",
			request: &dto.UpdateAdminRequest{
				Username:     "updated_admin",
				Password:     "new_password",
				ProfileImage: nil,
			},
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(admin, nil).Once()
				mockPasswordHelper.On("HashPassword", "new_password").Return("hashed_new_password", nil).Once()
				mockRepo.On("Update", mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Success update admin without profile image",
			request: &dto.UpdateAdminRequest{
				Username: "updated_admin",
				Password: "new_password",
			},
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(admin, nil).Once()
				mockPasswordHelper.On("HashPassword", "new_password").Return("hashed_new_password", nil).Once()
				mockRepo.On("Update", mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Error updating admin with invalid admin ID",
			request: &dto.UpdateAdminRequest{
				Username: "updated_admin",
			},
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Data admin tidak ditemukan"},
		},
		{
			name: "Error updating admin with password hashing failure",
			request: &dto.UpdateAdminRequest{
				Password: "new_password",
			},
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(admin, nil).Once()
				mockPasswordHelper.On("HashPassword", "new_password").Return("", errors.New("hashing error")).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mengubah data admin"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.UpdateAdmin(tc.request, tc.adminID)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockPasswordHelper.AssertExpectations(t)
			mockCloudinaryClient.AssertExpectations(t)
		})
	}
}

func TestDeleteAdmin(t *testing.T) {
	mockRepo := new(repositories.MockAdminRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)
	uc := usecases.NewAdminUsecase(mockRepo, mockCloudinaryClient, nil, nil)

	testCases := []struct {
		name          string
		adminID       uuid.UUID
		mockSetup     func()
		expectedError error
	}{
		{
			name:    "Success delete admin",
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(&entities.Admin{
					Id:              uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
					Username:        "existing_admin",
					Password:        "hashed_password",
					ProfileImageURL: nil,
				}, nil).Once()
				mockRepo.On("Delete", mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name:    "Error deleting admin with invalid admin ID",
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Data admin tidak ditemukan"},
		},
		{
			name:    "Error deleting admin with repository deletion failure",
			adminID: uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
			mockSetup: func() {
				mockRepo.On("FindById", mock.Anything).Return(&entities.Admin{
					Id:              uuid.MustParse("bce24300-4108-4e7f-86f7-d3c7fb28c68b"),
					Username:        "existing_admin",
					Password:        "hashed_password",
					ProfileImageURL: nil,
				}, nil).Once()
				mockRepo.On("Delete", mock.Anything).Return(errors.New("delete error")).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menghapus data admin"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := uc.DeleteAdmin(tc.adminID)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockCloudinaryClient.AssertExpectations(t)
		})
	}
}
