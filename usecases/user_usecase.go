package usecases

import (
	"fmt"
	"time"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ResendOTP(email string) (*dto.RegisterResponse, error)
	VerifyEmail(request *dto.VerifyEmailRequest) error
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
}

type userUsecase struct {
	userRepo  repositories.UserRepository
	cacheRepo *repositories.CacheRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, cacheRepo *repositories.CacheRepository) *userUsecase {
	return &userUsecase{userRepo, cacheRepo}
}

func (uc *userUsecase) Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	ifExistUsername, _ := uc.userRepo.FindByUsername(request.Username)
	if ifExistUsername != nil {
		return nil, &errorHandlers.ConflictError{Message: "username sudah ada, silahkan username lain"}
	}
	isExistEmail, _ := uc.userRepo.FindByEmail(request.Email)
	if isExistEmail != nil {
		return nil, &errorHandlers.ConflictError{Message: "email sudah terdaftar"}
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

	fmt.Println("otp", otp)
	uc.cacheRepo.Set(ref, otp)
	uc.cacheRepo.Set("email", user.Email)
	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal untuk mendaftar"}
	}

	if err := helpers.SendOTP(user.Email, user.Fullname, otp); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mengirimkan email"}
	}
	response := dto.RegisterResponse{ReferenceId: ref}
	return &response, nil
}

func (uc *userUsecase) ResendOTP(email string) (*dto.RegisterResponse, error) {
	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	otp := helpers.GenerateOTP()
	referenceId := helpers.GenerateReferenceId()
	uc.cacheRepo.Set(referenceId, otp)
	uc.cacheRepo.Set("email", email)

	if err := helpers.SendOTP(user.Email, user.Fullname, otp); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mengirimkan email"}
	}
	response := dto.RegisterResponse{ReferenceId: referenceId}
	return &response, nil
}

func (uc *userUsecase) VerifyEmail(request *dto.VerifyEmailRequest) error {
	cachedOTP, exists := uc.cacheRepo.Get(request.RefId)
	fmt.Println(cachedOTP, exists)
	if !exists {
		return &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	fmt.Println("get otp", cachedOTP)

	if cachedOTP != request.OTP {
		return &errorHandlers.BadRequestError{Message: "Kode OTP tidak cocok. Mohon periksa kembali dan masukkan dengan benar."}
	}

	now := time.Now()
	email, _ := uc.cacheRepo.Get("email")
	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}

	user.EmailVerifiedAt = &now
	_, err := uc.userRepo.Update(user)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Akun tidak ditemukan"}
	}
	return nil
}

func (uc *userUsecase) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := uc.userRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	if user.EmailVerifiedAt == nil {
		return nil, &errorHandlers.UnAuthorizedError{Message: "Email belum terverifikasi"}
	}
	if err := helpers.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Email atau password salah"}
	}

	accessToken, err := helpers.GenerateAccessToken(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	response := &dto.LoginResponse{
		Token: accessToken,
	}
	return response, nil
}
