package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"
)

type AdminUsecase interface {
	Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error)
}

type adminUsecase struct {
	repository repositories.AdminRepository
}

func NewAdminUsecase(repository repositories.AdminRepository) *adminUsecase {
	return &adminUsecase{repository}
}

func (uc *adminUsecase) Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error) {
	admin, err := uc.repository.FindByUsername(request.Username)
	if err != nil {
		return nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	if err = helpers.VerifyPassword(admin.Password, request.Password); err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Email atau password salah"}
	}

	accessToken, err := helpers.GenerateAccessToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	response := &dto.LoginAdminResponse{
		Username:     admin.Username,
		ProfileImage: admin.ProfileImageURL,
		Role:         admin.Role,
		Token:        accessToken,
	}

	return response, nil
}
