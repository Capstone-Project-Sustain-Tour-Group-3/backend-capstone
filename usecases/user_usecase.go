package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"
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

}
