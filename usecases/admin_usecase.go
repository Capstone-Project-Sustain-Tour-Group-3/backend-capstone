package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/devfeel/mapper"
)

type AdminUsecase interface {
	Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error)
	Logout(token string) error
	GetNewAccessToken(refreshToken string) (*dto.NewToken, error)
	GetAllAdmins(page, limit int, search string) (*[]dto.GetAllAdminResponse, *int64, error)
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
		return nil, &errorHandlers.ConflictError{Message: "Email atau password salah"}
	}

	accessToken, err := helpers.GenerateAccessToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	refreshToken, err := helpers.GenerateRefreshToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	admin.RefreshToken = refreshToken
	if err = uc.repository.Update(admin); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	response := &dto.LoginAdminResponse{
		Username:     admin.Username,
		ProfileImage: admin.ProfileImageURL,
		Role:         admin.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (uc *adminUsecase) Logout(token string) error {
	admin, err := uc.repository.GetUserByRefreshToken(token)
	if err != nil {
		return &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}
	}
	admin.RefreshToken = ""
	if err = uc.repository.Update(admin); err != nil {
		return &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (uc *adminUsecase) GetNewAccessToken(refreshToken string) (*dto.NewToken, error) {
	admin, err := uc.repository.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}
	}
	if admin.RefreshToken != refreshToken {
		return nil, &errorHandlers.UnAuthorizedError{Message: "Token tidak valid"}
	}
	accessToken, err := helpers.GenerateAccessToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	token := &dto.NewToken{
		AccessToken: accessToken,
	}
	return token, nil
}

func (uc *adminUsecase) GetAllAdmins(page, limit int, search string) (*[]dto.GetAllAdminResponse, *int64, error) {
	admins, total, err := uc.repository.FindAll((page-1)*limit, limit, search)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data admin"}
	}

	res := new([]dto.GetAllAdminResponse)
	err = mapper.MapperSlice(admins, res)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data admin"}
	}

	return res, total, nil
}
