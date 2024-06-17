package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/externals"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"gorm.io/gorm/utils/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type registerTestCase struct {
	name          string
	request       *dto.RegisterRequest
	mockSetup     func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient, passwordHelper *externals.MockPasswordHelper)
	expectedError error
	expectedResp  any
}

type resendOTPTestCase struct {
	name          string
	email         string
	mockSetup     func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient)
	expectedError error
	expectedResp  any
}

type loginTestCase struct {
	name          string
	request       *dto.LoginRequest
	mockSetup     func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper)
	expectedError error
}

type verifyEmailTestCase struct {
	name          string
	request       *dto.VerifyEmailRequest
	mockSetup     func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient)
	expectedError error
}

type logoutTestCase struct {
	name          string
	token         string
	mockSetup     func(userRepo *repositories.MockAuthRepository)
	expectedError error
}

type getNewAccessTokenTestCase struct {
	name          string
	refreshToken  string
	mockSetup     func(userRepo *repositories.MockAuthRepository, tokenHelper *externals.MockTokenHelper)
	expectedError error
	expectedResp  *dto.NewToken
}
type forgotPasswordTestCase struct {
	name          string
	request       *dto.ForgotPasswordRequest
	mockSetup     func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, passwordHelper *externals.MockPasswordHelper)
	expectedError error
}

func TestRegister(t *testing.T) {
	request := &dto.RegisterRequest{
		Username:    "testing",
		NamaLengkap: "user test",
		Email:       "test@example.com",
		NoTelepon:   "08123456789",
		Password:    "12345678",
	}

	requestWithSpace := &dto.RegisterRequest{
		Username:    "user with space",
		NamaLengkap: "user test",
		Email:       "test2@example.com",
	}

	expectedUser := &entities.User{
		Id:          uuid.New(),
		Email:       "test@example.com",
		Username:    "testing",
		Fullname:    "user test",
		PhoneNumber: "08123456789",
	}

	testCases := []registerTestCase{
		{
			name:    "Success register",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return("hashed-password", nil)
				cacheRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				emailSender.On("SendOTP", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				userRepo.On("Create", mock.Anything).Return(nil)
			},
			expectedError: nil,
			expectedResp:  &dto.RegisterResponse{ReferenceId: "123"},
		},
		{
			name:    "Username already exist",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(expectedUser, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				cacheRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "username sudah ada, silahkan username lain"},
			expectedResp:  nil,
		},
		{
			name:    "Email already exist",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(expectedUser, nil)
				cacheRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "email sudah terdaftar"},
			expectedResp:  nil,
		},
		{
			name:    "Username contains space",
			request: requestWithSpace,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", requestWithSpace.Username).Return(nil, nil)
				userRepo.On("FindByEmail", requestWithSpace.Email).Return(nil, nil)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "username tidak boleh mengandung spasi"},
			expectedResp:  nil,
		},
		{
			name:    "Failed create user",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient, passwordHelper *externals.MockPasswordHelper) {
				userRepo.On("FindByUsername", request.Username).Return(nil, nil)
				userRepo.On("FindByEmail", request.Email).Return(nil, nil)
				passwordHelper.On("HashPassword", request.Password).Return("hashed-password", nil)
				userRepo.On("Create", mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal untuk mendaftar"})
				cacheRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal untuk mendaftar"},
			expectedResp:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)
			tc.mockSetup(userRepo, cacheRepo, emailSender, passwordHelper)

			resp, err := uc.Register(tc.request)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
		})
	}
}

func TestResendOTP(t *testing.T) {
	email := "test@example.com"
	emailNotFound := "notfound@example.com"

	expectedUser := &entities.User{
		Id:       uuid.New(),
		Email:    email,
		Fullname: "user test",
	}

	testCases := []resendOTPTestCase{
		{
			name:  "Success resend OTP",
			email: email,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				userRepo.On("FindByEmail", email).Return(expectedUser, nil)
				cacheRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				emailSender.On("SendOTP", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				userRepo.On("Create", mock.Anything).Return(nil)
			},
			expectedError: nil,
			expectedResp:  &dto.RegisterResponse{ReferenceId: "reference-id"},
		},
		{
			name:  "Email not found",
			email: emailNotFound,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				userRepo.On("FindByEmail", emailNotFound).Return(nil, nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"},
			expectedResp:  nil,
		},
		{
			name:  "Failed to send email",
			email: email,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				userRepo.On("FindByEmail", email).Return(expectedUser, nil)
				cacheRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				emailSender.On("SendOTP", mock.Anything, mock.Anything, mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal mengirimkan email"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mengirimkan email"},
			expectedResp:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)
			tc.mockSetup(userRepo, cacheRepo, emailSender)

			resp, err := uc.ResendOTP(tc.email)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
		})
	}
}

func TestVerifyEmail(t *testing.T) {
	request := &dto.VerifyEmailRequest{
		RefId: "reference-id",
		OTP:   "123456",
	}

	expectedUser := &entities.User{
		Id:    uuid.New(),
		Email: "test@example.com",
	}

	testCases := []verifyEmailTestCase{
		{
			name:    "Success verify email",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				cacheRepo.On("Get", request.RefId).Return("123456", true)
				cacheRepo.On("Get", mock.Anything, mock.Anything).Return(expectedUser.Email, true)
				cacheRepo.On("Set", mock.Anything, mock.Anything).Return(nil)
				userRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)
				userRepo.On("Update", expectedUser).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:    "Invalid OTP",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				cacheRepo.On("Get", request.RefId).Return("654321", true)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Kode OTP tidak cocok. Mohon periksa kembali dan masukkan dengan benar."},
		},
		{
			name:    "Account not found",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				cacheRepo.On("Get", request.RefId).Return("123456", false)
				cacheRepo.On("Get", mock.Anything, mock.Anything).Return("", true)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"},
		},
		{
			name:    "Failed to update user",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, emailSender *externals.MockEmailClient) {
				cacheRepo.On("Get", request.RefId).Return("123456", true)
				cacheRepo.On("Get", mock.Anything, mock.Anything).Return(expectedUser.Email, true)
				cacheRepo.On("Set", mock.Anything, mock.Anything).Return(nil)
				userRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)
				userRepo.On("Update", expectedUser).Return(&errorHandlers.InternalServerError{Message: "Gagal memperbarui akun"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal memperbarui data user"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)
			tc.mockSetup(userRepo, cacheRepo, emailSender)

			err := uc.VerifyEmail(tc.request)

			require.Equal(t, tc.expectedError, err)
		})
	}
}

func TestLogin(t *testing.T) {
	request := &dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	expectedUser := &entities.User{
		Id:              uuid.New(),
		Email:           "test@example.com",
		Password:        "password",
		EmailVerifiedAt: tests.Now(),
	}

	testCases := []loginTestCase{
		{
			name:    "Success login",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("FindByEmail", request.Email).Return(expectedUser, nil)
				passwordHelper.On("VerifyPassword", request.Password, expectedUser.Password).Return(nil)
				tokenHelper.On("GenerateAccessToken", expectedUser).Return("access-token", nil)
				tokenHelper.On("GenerateRefreshToken", expectedUser).Return("refresh-token", nil)
				userRepo.On("Update", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:    "Account not found",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("FindByEmail", "test@example.com").Return(nil, nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"},
		},
		{
			name:    "Email not verified",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				user := *expectedUser
				user.EmailVerifiedAt = nil
				userRepo.On("FindByEmail", "test@example.com").Return(&user, nil)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Email belum terverifikasi"},
		},
		{
			name:    "Incorrect password",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)
				passwordHelper.On("VerifyPassword", request.Password, expectedUser.Password).Return(&errorHandlers.BadRequestError{Message: "Email atau password salah"})
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Email atau password salah"},
		},
		{
			name:    "Failed to generate access token",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)
				passwordHelper.On("VerifyPassword", request.Password, expectedUser.Password).Return(nil)
				tokenHelper.On("GenerateAccessToken", expectedUser).Return("", &errorHandlers.InternalServerError{Message: "Gagal generate access token"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal generate access token"},
		},
		{
			name:    "Failed to generate refresh token",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)
				passwordHelper.On("VerifyPassword", request.Password, expectedUser.Password).Return(nil)
				tokenHelper.On("GenerateAccessToken", expectedUser).Return("access-token", nil)
				tokenHelper.On("GenerateRefreshToken", expectedUser).Return("", &errorHandlers.InternalServerError{Message: "Gagal generate refresh token"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal generate refresh token"},
		},
		{
			name:    "Failed to update user",
			request: request,
			mockSetup: func(userRepo *repositories.MockAuthRepository, passwordHelper *externals.MockPasswordHelper, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)
				passwordHelper.On("VerifyPassword", request.Password, expectedUser.Password).Return(nil)
				tokenHelper.On("GenerateAccessToken", expectedUser).Return("access-token", nil)
				tokenHelper.On("GenerateRefreshToken", expectedUser).Return("refresh-token", nil)
				userRepo.On("Update", expectedUser).Return(&errorHandlers.InternalServerError{Message: "Gagal memperbarui akun"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal memperbarui akun"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)
			tc.mockSetup(userRepo, passwordHelper, tokenHelper)

			resp, err := uc.Login(tc.request)

			require.Equal(t, tc.expectedError, err)
			if tc.expectedError == nil {
				require.NotNil(t, resp)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	token := "refresh_token"

	expectedUser := &entities.User{
		Id:           uuid.New(),
		Email:        "test@example.com",
		RefreshToken: "refresh_token",
	}

	testCases := []logoutTestCase{
		{
			name:  "Success logout",
			token: token,
			mockSetup: func(userRepo *repositories.MockAuthRepository) {
				userRepo.On("GetUserByRefreshToken", token).Return(expectedUser, nil)
				userRepo.On("Update", expectedUser).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:  "Invalid token",
			token: token,
			mockSetup: func(userRepo *repositories.MockAuthRepository) {
				userRepo.On("GetUserByRefreshToken", token).Return(nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"})
			},
			expectedError: &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"},
		},
		{
			name:  "Empty token",
			token: "",
			mockSetup: func(userRepo *repositories.MockAuthRepository) {
				userRepo.On("GetUserByRefreshToken", "").Return(nil, nil)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Token tidak boleh kosong"},
		},
		{
			name:  "Failed to update user",
			token: token,
			mockSetup: func(userRepo *repositories.MockAuthRepository) {
				userRepo.On("GetUserByRefreshToken", token).Return(expectedUser, nil)
				userRepo.On("Update", expectedUser).Return(&errorHandlers.InternalServerError{Message: "Gagal logout"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal logout"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)

			tc.mockSetup(userRepo)

			err := uc.Logout(tc.token)

			require.Equal(t, tc.expectedError, err)
			if tc.expectedError == nil {
				require.Equal(t, "", expectedUser.RefreshToken)
			}
		})
	}
}

func TestGetNewAccessToken(t *testing.T) {
	refreshToken := "refresh_token"

	expectedUser := &entities.User{
		Id:           uuid.New(),
		Email:        "test@example.com",
		RefreshToken: "refresh_token",
	}

	accessToken := "new_access_token"

	testCases := []getNewAccessTokenTestCase{
		{
			name:         "Success get new access token",
			refreshToken: refreshToken,
			mockSetup: func(userRepo *repositories.MockAuthRepository, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("GetUserByRefreshToken", refreshToken).Return(expectedUser, nil)
				tokenHelper.On("GenerateAccessToken", expectedUser).Return(accessToken, nil)
			},
			expectedError: nil,
			expectedResp:  &dto.NewToken{AccessToken: accessToken},
		},
		{
			name:         "Invalid token",
			refreshToken: refreshToken,
			mockSetup: func(userRepo *repositories.MockAuthRepository, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("GetUserByRefreshToken", refreshToken).Return(nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"})
			},
			expectedError: &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"},
			expectedResp:  nil,
		},
		{
			name:         "Empty token",
			refreshToken: "",
			mockSetup: func(userRepo *repositories.MockAuthRepository, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("GetUserByRefreshToken", "").Return(nil, nil)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Token tidak boleh kosong"},
			expectedResp:  nil,
		},
		{
			name:         "Failed to generate access token",
			refreshToken: refreshToken,
			mockSetup: func(userRepo *repositories.MockAuthRepository, tokenHelper *externals.MockTokenHelper) {
				userRepo.On("GetUserByRefreshToken", refreshToken).Return(expectedUser, nil)
				tokenHelper.On("GenerateAccessToken", expectedUser).Return("", &errorHandlers.InternalServerError{Message: "Gagal membuat access token"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal membuat access token"},
			expectedResp:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)

			tc.mockSetup(userRepo, tokenHelper)
			resp, err := uc.GetNewAccessToken(tc.refreshToken)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
		})
	}
}

func TestForgotPassword(t *testing.T) {
	refId := "ref_id"
	email := "test@example.com"
	password := "new_password"
	hashedPassword := "hashed_password"

	expectedUser := &entities.User{
		Id:       uuid.New(),
		Email:    email,
		Password: "old_password",
	}

	testCases := []forgotPasswordTestCase{
		{
			name: "Success forgot password",
			request: &dto.ForgotPasswordRequest{
				RefId:              refId,
				Password:           password,
				KonfirmasiPassword: password,
			},
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, passwordHelper *externals.MockPasswordHelper) {
				cacheRepo.On("Get", refId+"_email").Return(email, true)
				userRepo.On("FindByEmail", email).Return(expectedUser, nil)
				passwordHelper.On("HashPassword", password).Return(hashedPassword, nil)
				userRepo.On("Update", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "User not found",
			request: &dto.ForgotPasswordRequest{
				RefId:              refId,
				Password:           password,
				KonfirmasiPassword: password,
			},
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, passwordHelper *externals.MockPasswordHelper) {
				cacheRepo.On("Get", refId+"_email").Return(email, true)
				userRepo.On("FindByEmail", email).Return(nil, nil)
			},
			expectedError: &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"},
		},
		{
			name: "Password mismatch",
			request: &dto.ForgotPasswordRequest{
				RefId:              refId,
				Password:           password,
				KonfirmasiPassword: "different_password",
			},
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, passwordHelper *externals.MockPasswordHelper) {
				cacheRepo.On("Get", refId+"_email").Return(email, true)
				userRepo.On("FindByEmail", email).Return(expectedUser, nil)
			},
			expectedError: &errorHandlers.BadRequestError{Message: "Password tidak cocok"},
		},
		{
			name: "Failed to hash password",
			request: &dto.ForgotPasswordRequest{
				RefId:              refId,
				Password:           password,
				KonfirmasiPassword: password,
			},
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, passwordHelper *externals.MockPasswordHelper) {
				cacheRepo.On("Get", refId+"_email").Return(email, true)
				userRepo.On("FindByEmail", email).Return(expectedUser, nil)
				passwordHelper.On("HashPassword", password).Return("", &errorHandlers.InternalServerError{Message: "Gagal hashing password"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal hashing password"},
		},
		{
			name: "Failed to update user",
			request: &dto.ForgotPasswordRequest{
				RefId:              refId,
				Password:           password,
				KonfirmasiPassword: password,
			},
			mockSetup: func(userRepo *repositories.MockAuthRepository, cacheRepo *repositories.MockCacheRepository, passwordHelper *externals.MockPasswordHelper) {
				cacheRepo.On("Get", refId+"_email").Return(email, true)
				userRepo.On("FindByEmail", email).Return(expectedUser, nil)
				passwordHelper.On("HashPassword", password).Return(hashedPassword, nil)
				userRepo.On("Update", mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal mengubah password"})
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mengubah password"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(repositories.MockAuthRepository)
			cacheRepo := new(repositories.MockCacheRepository)
			emailSender := new(externals.MockEmailClient)
			passwordHelper := new(externals.MockPasswordHelper)
			tokenHelper := new(externals.MockTokenHelper)
			uc := usecases.NewAuthUsecase(userRepo, cacheRepo, emailSender, passwordHelper, tokenHelper)

			tc.mockSetup(userRepo, cacheRepo, passwordHelper)

			err := uc.ForgotPassword(tc.request)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			userRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
			passwordHelper.AssertExpectations(t)
		})
	}
}
