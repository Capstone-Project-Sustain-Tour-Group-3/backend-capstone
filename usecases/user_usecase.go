package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type userUsecase struct {
	userRepo  repositories.UserRepository
	cacheRepo repositories.CacheRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, cacheRepo repositories.CacheRepository) *userUsecase {
	return &userUsecase{userRepo, cacheRepo}
}

func (uc *userUsecase) Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	isExist, _ := uc.userRepo.FindByEmail(request.Email)
	if isExist != nil {
		return nil, &errorHandlers.ConflictError{Message: "username/email sudah terdaftar"}
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
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
	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}

	response := dto.RegisterResponse{ReferenceId: ref}
	return &response, nil
}
