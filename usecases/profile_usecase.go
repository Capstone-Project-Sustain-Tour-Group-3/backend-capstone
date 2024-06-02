package usecases

import (
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"
	"github.com/google/uuid"
)

type ProfileUsecase interface {
	GetDetailUser(id uuid.UUID) (*entities.User, error)
}

type profileUsecase struct {
	repository repositories.UserRepository
}

func NewProfileUsecase(repository repositories.UserRepository) *profileUsecase {
	return &profileUsecase{repository: repository}
}

func (uc *profileUsecase) GetDetailUser(id uuid.UUID) (*entities.User, error) {
	user, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "User tidak ditemukan"}
	}
	return user, nil
}
