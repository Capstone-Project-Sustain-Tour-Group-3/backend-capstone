package usecases

import (
	"strings"
	"time"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type AuthUsecase interface {
	Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ResendOTP(email string) (*dto.RegisterResponse, error)
	VerifyEmail(request *dto.VerifyEmailRequest) error
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
	Logout(token string) error
	GetNewAccessToken(refreshToken string) (*dto.NewToken, error)
	ForgotPassword(request *dto.ChangePasswordRequest) error
}

type authUsecase struct {
	userRepo  repositories.AuthRepository
	cacheRepo *repositories.CacheRepository
}

func NewAuthUsecase(userRepo repositories.AuthRepository, cacheRepo *repositories.CacheRepository) *authUsecase {
	return &authUsecase{userRepo, cacheRepo}
}

func (uc *authUsecase) Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	ifExistUsername, _ := uc.userRepo.FindByUsername(request.Username)
	if ifExistUsername != nil {
		return nil, &errorHandlers.ConflictError{Message: "username sudah ada, silahkan username lain"}
	}
	isExistEmail, _ := uc.userRepo.FindByEmail(request.Email)
	if isExistEmail != nil {
		return nil, &errorHandlers.ConflictError{Message: "email sudah terdaftar"}
	}

	if strings.Contains(request.Username, " ") {
		return nil, &errorHandlers.BadRequestError{Message: "username tidak boleh mengandung spasi"}
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal hashing password"}
	}
	user := &entities.User{
		Id:          uuid.New(),
		Email:       request.Email,
		Password:    password,
		Username:    request.Username,
		Fullname:    request.NamaLengkap,
		PhoneNumber: request.NoTelepon,
	}
	ref := helpers.GenerateReferenceId()
	otp := helpers.GenerateOTP()

	uc.cacheRepo.Set(ref, otp)
	uc.cacheRepo.Set(ref+"_email", request.Email)
	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendaftar"}
	}

	if err = helpers.SendOTP(user.Email, user.Fullname, otp); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mengirimkan email"}
	}
	response := dto.RegisterResponse{ReferenceId: ref}
	return &response, nil
}

func (uc *authUsecase) ResendOTP(email string) (*dto.RegisterResponse, error) {
	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	otp := helpers.GenerateOTP()
	referenceId := helpers.GenerateReferenceId()
	uc.cacheRepo.Set(referenceId, otp)
	uc.cacheRepo.Set(referenceId+"_email", user.Email)

	if err := helpers.SendOTP(user.Email, user.Fullname, otp); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mengirimkan email"}
	}
	response := dto.RegisterResponse{ReferenceId: referenceId}
	return &response, nil
}

func (uc *authUsecase) VerifyEmail(request *dto.VerifyEmailRequest) error {
	cachedOTP, exists := uc.cacheRepo.Get(request.RefId)
	if !exists {
		return &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}

	if cachedOTP != request.OTP {
		return &errorHandlers.BadRequestError{Message: "Kode OTP tidak cocok. Mohon periksa kembali dan masukkan dengan benar."}
	}

	now := time.Now()
	email, _ := uc.cacheRepo.Get(request.RefId + "_email")
	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}

	user.EmailVerifiedAt = &now
	err := uc.userRepo.Update(user)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Akun tidak ditemukan"}
	}
	return nil
}

func (uc *authUsecase) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, _ := uc.userRepo.FindByEmail(request.Email)
	if user == nil {
		return nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	if err := helpers.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Email atau password salah"}
	}

	accessToken, err := helpers.GenerateAccessToken(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	refreshToken, err := helpers.GenerateRefreshToken(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	user.RefreshToken = refreshToken
	err = uc.userRepo.Update(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	response := &dto.LoginResponse{
		Username:     user.Username,
		ProfileImage: user.ProfileImageUrl,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, nil
}

func (uc *authUsecase) ForgotPassword(request *dto.ChangePasswordRequest) error {
	email, _ := uc.cacheRepo.Get(request.RefId + "_email")

	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	if request.Password != request.KonfirmasiPassword {
		return &errorHandlers.BadRequestError{Message: "Password tidak cocok"}
	}
	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal hashing password"}
	}
	user.Password = password
	err = uc.userRepo.Update(user)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal mengubah password"}
	}
	return nil
}

func (uc *authUsecase) Logout(token string) error {
	user, err := uc.userRepo.GetUserByRefreshToken(token)
	if err != nil {
		return &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}
	}
	user.RefreshToken = ""
	err = uc.userRepo.Update(user)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (uc *authUsecase) GetNewAccessToken(refreshToken string) (*dto.NewToken, error) {
	user, err := uc.userRepo.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}
	}
	if user.RefreshToken != refreshToken {
		return nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}
	}
	accessToken, err := helpers.GenerateAccessToken(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	token := &dto.NewToken{
		AccessToken: accessToken,
	}
	return token, nil
}
