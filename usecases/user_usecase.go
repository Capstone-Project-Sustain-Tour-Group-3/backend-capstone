package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type UserUsecase interface {
	Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ResendOTP() (*dto.RegisterResponse, error)
	VerifyEmail(request *dto.VerifyEmailRequest) error
}

type userUsecase struct {
	userRepo  repositories.UserRepository
	cacheRepo *repositories.CacheRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, cacheRepo *repositories.CacheRepository) *userUsecase {
	return &userUsecase{userRepo, cacheRepo}
}

func (uc *userUsecase) Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	isExist, _ := uc.userRepo.FindByEmail(request.Email)
	if isExist != nil {
		return nil, &errorHandlers.ConflictError{Message: "username/email sudah terdaftar"}
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{"Gagal hashing password"}
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
		return nil, &errorHandlers.InternalServerError{"Gagal untuk mendaftar"}
	}

	if err := helpers.SendOTP(user.Email, otp); err != nil {
		return nil, &errorHandlers.InternalServerError{"Gagal mengirimkan email"}
	}
	response := dto.RegisterResponse{ReferenceId: ref}
	return &response, nil
}

func (uc *userUsecase) ResendOTP() (*dto.RegisterResponse, error) {
	email, _ := uc.cacheRepo.Get("email")
	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return nil, &errorHandlers.ConflictError{"Akun tidak ditemukan"}
	}
	otp := helpers.GenerateOTP()
	referenceId := helpers.GenerateReferenceId()
	uc.cacheRepo.Set(referenceId, otp)

	if err := helpers.SendOTP(user.Email, otp); err != nil {
		return nil, &errorHandlers.InternalServerError{"Gagal mengirimkan email"}
	}
	response := dto.RegisterResponse{ReferenceId: referenceId}
	return &response, nil
}

func (uc *userUsecase) VerifyEmail(request *dto.VerifyEmailRequest) error {
	cachedOTP, exists := uc.cacheRepo.Get(request.RefId)
	if !exists {
		return &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	fmt.Println("get otp", cachedOTP)

	if cachedOTP != request.OTP {
		return &errorHandlers.BadRequestError{"Kode OTP tidak cocok. Mohon periksa kembali dan masukkan dengan benar."}
	}

	now := time.Now()
	email, _ := uc.cacheRepo.Get("email")
	user, _ := uc.userRepo.FindByEmail(email)
	if user == nil {
		return &errorHandlers.ConflictError{"Akun tidak ditemukan"}
	}

	user.EmailVerifiedAt = &now
	_, err := uc.userRepo.Update(user)
	if err != nil {
		return &errorHandlers.InternalServerError{"Akun tidak ditemukan"}
	}
	return nil
}
