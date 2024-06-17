package tests

import (
	"bytes"
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

type findByIdTestCase struct {
	name           string
	id             uuid.UUID
	mockSetup      func(userRepo *repositories.MockUserRepository)
	expectedError  error
	expectedResult *entities.User
}

type findAllTestCase struct {
	name          string
	page          int
	limit         int
	searchQuery   string
	mockSetup     func(userRepo *repositories.MockUserRepository)
	expectedError error
	expectedUsers *[]entities.User
	expectedTotal *int64
}

type createTestCase struct {
	name          string
	request       *dto.UserRequest
	mockSetup     func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper)
	expectedError error
}

type updateTestCase struct {
	name          string
	id            uuid.UUID
	request       *dto.UserRequest
	mockSetup     func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper)
	expectedError error
}

type deleteTestCase struct {
	name          string
	id            uuid.UUID
	mockSetup     func(userRepo *repositories.MockUserRepository, cloudinaryClient *externals.MockCloudinaryClient)
	expectedError error
}

type profilePictureTestCase struct {
	name          string
	request       *dto.UserRequest
	user          *entities.User
	mockSetup     func(cloudinaryClient *externals.MockCloudinaryClient, imageValidator *externals.MockImageValidator)
	expectedError error
}

func TestFindById(t *testing.T) {
	id := uuid.New()
	user := &entities.User{
		Id:    id,
		Email: "test@example.com",
	}

	testCases := []findByIdTestCase{
		{
			name: "Success",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository) {
				userRepo.On("FindById", id).Return(user, nil)
			},
			expectedError:  nil,
			expectedResult: user,
		},
		{
			name: "User not found",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository) {
				userRepo.On("FindById", id).Return(nil, &errorHandlers.NotFoundError{Message: "User tidak ditemukan"})
			},
			expectedError:  &errorHandlers.NotFoundError{Message: "User tidak ditemukan"},
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			passwordHelper := new(externals.MockPasswordHelper)
			cloudinaryClient := new(externals.MockCloudinaryClient)
			imageValidation := new(externals.MockImageValidator)
			tc.mockSetup(userRepo)

			uc := usecases.NewUserUsecase(userRepo, cloudinaryClient, passwordHelper, imageValidation)

			result, err := uc.FindById(tc.id)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
			}

			require.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestFindAll(t *testing.T) {
	page := 1
	limit := 10
	searchQuery := "test"
	users := &[]entities.User{
		{Id: uuid.New(), Email: "user1@example.com"},
		{Id: uuid.New(), Email: "user2@example.com"},
	}
	total := int64(2)

	testCases := []findAllTestCase{
		{
			name:        "Success",
			page:        page,
			limit:       limit,
			searchQuery: searchQuery,
			mockSetup: func(userRepo *repositories.MockUserRepository) {
				userRepo.On("FindAll", page, limit, searchQuery).Return(users, &total, nil)
			},
			expectedError: nil,
			expectedUsers: users,
			expectedTotal: &total,
		},
		{
			name:        "Internal server error",
			page:        page,
			limit:       limit,
			searchQuery: searchQuery,
			mockSetup: func(userRepo *repositories.MockUserRepository) {
				userRepo.On("FindAll", page, limit, searchQuery).Return(nil, nil, &errorHandlers.InternalServerError{Message: "Gagal untuk menampilkan data user"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk menampilkan data user"},
			expectedUsers: nil,
			expectedTotal: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			passwordHelper := new(externals.MockPasswordHelper)
			cloudinaryClient := new(externals.MockCloudinaryClient)
			imageValidation := new(externals.MockImageValidator)
			tc.mockSetup(userRepo)

			uc := usecases.NewUserUsecase(userRepo, cloudinaryClient, passwordHelper, imageValidation)

			res, count, err := uc.FindAll(tc.page, tc.limit, tc.searchQuery)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotNil(t, count)
			}

			require.Equal(t, tc.expectedUsers, res)
			require.Equal(t, tc.expectedTotal, count)
			userRepo.AssertExpectations(t)
		})
	}
}

func TestCreate(t *testing.T) {
	request := &dto.UserRequest{
		Username:     "testuser",
		Email:        "test@example.com",
		Password:     "password",
		NamaLengkap:  "Test User",
		Bio:          "Test Bio",
		NoTelepon:    "1234567890",
		JenisKelamin: "Male",
		Kota:         "Test City",
		Provinsi:     "Test Province",
	}

	passwordHash := "hashed_password"

	testCases := []createTestCase{
		{
			name:    "Success create user",
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return(passwordHash, nil)
				userRepo.On("Create", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:    "Username already exists",
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(&entities.User{}, nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Username sudah digunakan"},
		},
		{
			name:    "Email already exists",
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(&entities.User{}, nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Email sudah digunakan"},
		},
		{
			name:    "Hashing password failed",
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return("", &errorHandlers.InternalServerError{Message: "Gagal hashing password"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal hashing password"},
		},
		{
			name:    "Create user failed",
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return(passwordHash, nil)
				userRepo.On("Create", mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal untuk menambah data user"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk menambah data user"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			passwordHelper := new(externals.MockPasswordHelper)
			cloudinaryClient := new(externals.MockCloudinaryClient)
			imageValidation := new(externals.MockImageValidator)

			uc := usecases.NewUserUsecase(userRepo, cloudinaryClient, passwordHelper, imageValidation)

			tc.mockSetup(userRepo, passwordHelper)

			err := uc.Create(tc.request)

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

func TestUpdate(t *testing.T) {
	id := uuid.New()
	request := &dto.UserRequest{
		Username:     "updateduser",
		Email:        "updated@example.com",
		Password:     "newpassword",
		NamaLengkap:  "Updated User",
		Bio:          "Updated Bio",
		NoTelepon:    "9876543210",
		JenisKelamin: "Female",
		Kota:         "Updated City",
		Provinsi:     "Updated Province",
	}

	existingUser := &entities.User{
		Id:       id,
		Username: "existinguser",
		Email:    "existing@example.com",
	}

	passwordHash := "hashed_new_password"

	testCases := []updateTestCase{
		{
			name:    "Success update user",
			id:      id,
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", id).Return(existingUser, nil)
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return(passwordHash, nil)
				userRepo.On("Update", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:    "User not found",
			id:      id,
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", id).Return(nil, nil)
			},
			expectedError: &errorHandlers.NotFoundError{Message: "User tidak ditemukan"},
		},
		{
			name:    "Hashing password failed",
			id:      id,
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", id).Return(existingUser, nil)
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return("", &errorHandlers.InternalServerError{Message: "Gagal hashing password"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal hashing password"},
		},
		{
			name:    "Update user failed",
			id:      id,
			request: request,
			mockSetup: func(userRepo *repositories.MockUserRepository, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindById", id).Return(existingUser, nil)
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return(passwordHash, nil)
				userRepo.On("Update", mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			passwordHelper := new(externals.MockPasswordHelper)
			cloudinaryClient := new(externals.MockCloudinaryClient)
			imageValidation := new(externals.MockImageValidator)

			uc := usecases.NewUserUsecase(userRepo, cloudinaryClient, passwordHelper, imageValidation)

			tc.mockSetup(userRepo, passwordHelper)

			err := uc.Update(tc.id, tc.request)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	id := uuid.New()
	profileImageUrl := "http://example.com/image.jpg"

	existingUserWithImage := &entities.User{
		Id:              id,
		ProfileImageUrl: &profileImageUrl,
	}

	existingUserWithoutImage := &entities.User{
		Id: id,
	}

	testCases := []deleteTestCase{
		{
			name: "Success delete user with profile image",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinaryClient *externals.MockCloudinaryClient) {
				userRepo.On("FindById", id).Return(existingUserWithImage, nil)
				cloudinaryClient.On("DeleteImage", profileImageUrl).Return(nil)
				userRepo.On("Delete", existingUserWithImage).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Success delete user without profile image",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinaryClient *externals.MockCloudinaryClient) {
				userRepo.On("FindById", id).Return(existingUserWithoutImage, nil)
				userRepo.On("Delete", existingUserWithoutImage).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "User not found",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinaryClient *externals.MockCloudinaryClient) {
				userRepo.On("FindById", id).Return(nil, nil)
			},
			expectedError: &errorHandlers.NotFoundError{Message: "User tidak ditemukan"},
		},
		{
			name: "Failed to delete profile image",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinaryClient *externals.MockCloudinaryClient) {
				userRepo.On("FindById", id).Return(existingUserWithImage, nil)
				cloudinaryClient.On("DeleteImage", profileImageUrl).Return(&errorHandlers.InternalServerError{Message: "Gagal menghapus foto profil"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menghapus foto profil"},
		},
		{
			name: "Failed to delete user data",
			id:   id,
			mockSetup: func(userRepo *repositories.MockUserRepository, cloudinaryClient *externals.MockCloudinaryClient) {
				userRepo.On("FindById", id).Return(existingUserWithoutImage, nil)
				userRepo.On("Delete", existingUserWithoutImage).Return(&errorHandlers.InternalServerError{Message: "Gagal untuk menghapus data user"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk menghapus data user"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockUserRepository)
			cloudinaryClient := new(externals.MockCloudinaryClient)
			passwordHelper := new(externals.MockPasswordHelper)
			imageValidation := new(externals.MockImageValidator)

			uc := usecases.NewUserUsecase(userRepo, cloudinaryClient, passwordHelper, imageValidation)

			tc.mockSetup(userRepo, cloudinaryClient)

			err := uc.Delete(tc.id)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			userRepo.AssertExpectations(t)
			cloudinaryClient.AssertExpectations(t)
		})
	}
}

func createMultipartFileHeader(filename string, content []byte) *multipart.FileHeader {
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

func TestHandleProfilePictureUpdate(t *testing.T) {
	fileContent := []byte("file content")
	fileHeader := createMultipartFileHeader("test.jpg", fileContent)
	request := &dto.UserRequest{
		FotoProfil: fileHeader,
	}

	existingProfileImageUrl := "http://example.com/image.jpg"
	userWithProfileImage := &entities.User{
		ProfileImageUrl: &existingProfileImageUrl,
	}

	userWithoutProfileImage := &entities.User{}
	testCases := []profilePictureTestCase{
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
			uc := usecases.NewUserUsecase(nil, cloudinaryClient, nil, imageValidator)

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
