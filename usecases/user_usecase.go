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
	Register(request *dto.RegisterRequest) (*entities.User, error)
}

type userUsecase struct {
	repository repositories.UserRepository
}

func NewUserUsecase(repository repositories.UserRepository) *userUsecase {
	return &userUsecase{repository}
}

func (uc *userUsecase) Register(request *dto.RegisterRequest) (*entities.User, error) {
	isExist, _ := uc.repository.FindByEmail(request.Email)
	if isExist != nil {
		return nil, &errorHandlers.ConflictError{Message: "username/email sudah terdaftar"}
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	user := entities.User{
		Id:          uuid.New(),
		Email:       request.Email,
		Password:    password,
		Username:    request.Username,
		Fullname:    request.NamaLengkap,
		PhoneNumber: request.NoTelepon,
	}
}
