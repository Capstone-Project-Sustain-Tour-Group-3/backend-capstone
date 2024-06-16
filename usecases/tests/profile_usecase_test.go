package tests

import (
	"bytes"
	"errors"
	"mime"
	"mime/multipart"
	"net/http"
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
)

type detailUser struct {
	name          string
	userID        uuid.UUID
	mockSetup     func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, iValidation *externals.MockImageValidator)
	expectedUser  *entities.User
	expectedError error
}

func TestGetDetailUser(t *testing.T) {
	user := &entities.User{
		Id:       uuid.New(),
		Username: "TestUser",
	}
	testCases := []detailUser{
		{
			name:   "Success when user is found",
			userID: uuid.New(),
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, iValidation *externals.MockImageValidator) {
				userRepo.On("FindById", mock.Anything).Return(user, nil)
			},
			expectedUser:  user,
			expectedError: nil,
		},
		{
			name:   "Error when user is not found",
			userID: uuid.New(),
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, iValidation *externals.MockImageValidator) {
				userRepo.On("FindById", mock.Anything).Return(nil, errors.New("user not found"))
			},
			expectedUser:  nil,
			expectedError: &errorHandlers.NotFoundError{Message: "User tidak ditemukan"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			cloudinary := new(externals.MockCloudinaryClient)
			passwordHelper := new(externals.MockPasswordHelper)
			iValidation := new(externals.MockImageValidator)
			uc := usecases.NewProfileUsecase(userRepo, cloudinary, passwordHelper, iValidation)

			tc.mockSetup(userRepo, cloudinary, passwordHelper, iValidation)

			res, err := uc.GetDetailUser(tc.userID)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedUser, res)

			userRepo.AssertExpectations(t)
			cloudinary.AssertExpectations(t)
			passwordHelper.AssertExpectations(t)
			iValidation.AssertExpectations(t)
		})
	}
}

func TestInsertUserDetail(t *testing.T) {
	userID := uuid.New()
	request := &dto.UserDetailRequest{
		Username:     "newusername",
		NamaLengkap:  "New Fullname",
		NoTelepon:    "08123456789",
		Bio:          "New bio",
		JenisKelamin: "Male",
		Kota:         "New City",
		Provinsi:     "New Province",
		Email:        "newemail@example.com",
	}

	existingUser := &entities.User{
		Id:          userID,
		Username:    "oldusername",
		Fullname:    "Old Fullname",
		PhoneNumber: "08123456789",
		Bio:         "Old bio",
		Gender:      "Male",
		City:        "Old City",
		Province:    "Old Province",
		Email:       "oldemail@example.com",
	}

	testCases := []struct {
		name          string
		request       *dto.UserDetailRequest
		userID        uuid.UUID
		mockSetup     func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, imageValidator *externals.MockImageValidator)
		expectedError error
	}{
		{
			name:    "Success update user detail",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, imageValidator *externals.MockImageValidator) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				userRepo.On("Update", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:    "User not found",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, imageValidator *externals.MockImageValidator) {
				userRepo.On("FindById", userID).Return(nil, errors.New("user not found"))
			},
			expectedError: &errorHandlers.ConflictError{Message: "User tidak ditemukan"},
		},
		{
			name:    "Failed to update user",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinary *externals.MockCloudinaryClient, passwordHelper *externals.MockPasswordHelper, imageValidator *externals.MockImageValidator) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				userRepo.On("Update", mock.Anything).Return(errors.New("update failed"))
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			cloudinaryClient := new(externals.MockCloudinaryClient)
			passwordHelper := new(externals.MockPasswordHelper)
			imageValidator := new(externals.MockImageValidator)
			uc := usecases.NewProfileUsecase(userRepo, cloudinaryClient, passwordHelper, imageValidator)

			tc.mockSetup(userRepo, cloudinaryClient, passwordHelper, imageValidator)

			err := uc.InsertUserDetail(tc.request, tc.userID)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	userID := uuid.New()
	request := &dto.ChangePasswordRequest{
		PasswordLama:       "oldpassword",
		PasswordBaru:       "newpassword",
		KonfirmasiPassword: "newpassword",
	}

	existingUser := &entities.User{
		Id:       userID,
		Password: "hashedOldPassword",
	}

	testCases := []struct {
		name          string
		request       *dto.ChangePasswordRequest
		userID        uuid.UUID
		mockSetup     func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper)
		expectedError error
	}{
		{
			name:    "Success change password",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				passwordHelper.On("VerifyPassword", existingUser.Password, request.PasswordLama).Return(nil)
				passwordHelper.On("HashPassword", request.PasswordBaru).Return("hashedNewPassword", nil)
				userRepo.On("Update", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:    "User not found",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", userID).Return(nil, errors.New("user not found"))
			},
			expectedError: &errorHandlers.ConflictError{Message: "User tidak ditemukan"},
		},
		{
			name:    "Old password not valid",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				passwordHelper.On("VerifyPassword", existingUser.Password, request.PasswordLama).Return(errors.New("invalid password"))
			},
			expectedError: &errorHandlers.ConflictError{Message: "Password lama tidak valid"},
		},
		{
			name: "New password and confirmation password do not match",
			request: &dto.ChangePasswordRequest{
				PasswordLama:       "oldpassword",
				PasswordBaru:       "newpassword",
				KonfirmasiPassword: "differentpassword",
			},
			userID: userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				passwordHelper.On("VerifyPassword", existingUser.Password, request.PasswordLama).Return(nil)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Password dan konfirmasi password tidak cocok"},
		},
		{
			name:    "Failed to hash password",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				passwordHelper.On("VerifyPassword", existingUser.Password, request.PasswordLama).Return(nil)
				passwordHelper.On("HashPassword", request.PasswordBaru).Return("", errors.New("hashing error"))
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal hashing password"},
		},
		{
			name:    "Failed to update password",
			request: request,
			userID:  userID,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", userID).Return(existingUser, nil)
				passwordHelper.On("VerifyPassword", existingUser.Password, request.PasswordLama).Return(nil)
				passwordHelper.On("HashPassword", request.PasswordBaru).Return("hashedNewPassword", nil)
				userRepo.On("Update", mock.Anything).Return(errors.New("update failed"))
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui password"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			passwordHelper := new(externals.MockPasswordHelper)
			uc := usecases.NewProfileUsecase(userRepo, nil, passwordHelper, nil)

			tc.mockSetup(userRepo, passwordHelper)

			err := uc.ChangePassword(tc.request, tc.userID)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			userRepo.AssertExpectations(t)
			passwordHelper.AssertExpectations(t)
		})
	}
}

func createMultipartFileHeader2(filename string, content []byte) *multipart.FileHeader {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		panic(err)
	}
	_, err = part.Write(content)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req := &http.Request{
		Header: make(http.Header),
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	_, params, err := mime.ParseMediaType(writer.FormDataContentType())
	if err != nil {
		panic(err)
	}
	mr := multipart.NewReader(body, params["boundary"])
	form, err := mr.ReadForm(10 << 20) // 10 MB max memory
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = form.RemoveAll(); err != nil {
			panic(err)
		}
	}()

	return form.File["file"][0]
}

func TestHandleProfilePictureUpdate2(t *testing.T) {
	fileContent := []byte("file content")
	fileHeader := createMultipartFileHeader2("test.jpg", fileContent)
	request := &dto.UserDetailRequest{
		FotoProfil: fileHeader,
	}

	existingProfileImageUrl := "http://example.com/image.jpg"
	userWithProfileImage := &entities.User{
		ProfileImageUrl: &existingProfileImageUrl,
	}

	userWithoutProfileImage := &entities.User{}
	testCases := []struct {
		name          string
		request       *dto.UserDetailRequest
		user          *entities.User
		mockSetup     func(cloudinaryClient *externals.MockCloudinaryClient, imageValidator *externals.MockImageValidator)
		expectedError error
	}{
		{
			name:    "Success handle profile picture update with existing profile image",
			request: request,
			user:    userWithProfileImage,
			mockSetup: func(cloudinaryClient *externals.MockCloudinaryClient, imageValidator *externals.MockImageValidator) {
				imageValidator.On("IsValidImageType", request.FotoProfil).Return(true)
				imageValidator.On("IsValidImageSize", request.FotoProfil).Return(true)
				cloudinaryClient.On("DeleteImage", existingProfileImageUrl).Return(nil)
				cloudinaryClient.On("UploadImage", mock.Anything, "users").Return("http://newimage.com", nil)
			},
			expectedError: nil,
		},
		{
			name:    "Success handle profile picture update without existing profile image",
			request: request,
			user:    userWithoutProfileImage,
			mockSetup: func(cloudinaryClient *externals.MockCloudinaryClient, imageValidator *externals.MockImageValidator) {
				imageValidator.On("IsValidImageType", request.FotoProfil).Return(true)
				imageValidator.On("IsValidImageSize", request.FotoProfil).Return(true)
				cloudinaryClient.On("UploadImage", mock.Anything, "users").Return("http://newimage.com", nil)
			},
			expectedError: nil,
		},
		{
			name:    "Invalid image type",
			request: request,
			user:    userWithoutProfileImage,
			mockSetup: func(cloudinaryClient *externals.MockCloudinaryClient, imageValidator *externals.MockImageValidator) {
				imageValidator.On("IsValidImageType", request.FotoProfil).Return(false)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Tipe foto profil tidak valid"},
		},
		{
			name:    "Invalid image size",
			request: request,
			user:    userWithoutProfileImage,
			mockSetup: func(cloudinaryClient *externals.MockCloudinaryClient, imageValidator *externals.MockImageValidator) {
				imageValidator.On("IsValidImageType", request.FotoProfil).Return(true)
				imageValidator.On("IsValidImageSize", request.FotoProfil).Return(false)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Ukuran foto profil tidak valid"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cloudinaryClient := new(externals.MockCloudinaryClient)
			imageValidator := new(externals.MockImageValidator)
			uc := usecases.NewProfileUsecase(nil, cloudinaryClient, nil, imageValidator)

			tc.mockSetup(cloudinaryClient, imageValidator)

			err := uc.HandleProfilePictureUpdate(tc.request, tc.user)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			cloudinaryClient.AssertExpectations(t)
			imageValidator.AssertExpectations(t)
		})
	}
}
