package usecases

import (
	"errors"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/externals/cloudinary"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminUsecase interface {
	Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error)
	Logout(token string) error
	GetNewAccessToken(refreshToken string) (*dto.NewToken, error)
	GetAllAdmins(page, limit int, search string) (*[]dto.GetAllAdminResponse, *int64, error)
	GetAdminDetail(id uuid.UUID) (*dto.GetDetailAdminResponse, error)
	CreateAdmin(request *dto.AdminRequest) error
	UpdateAdmin(request *dto.UpdateAdminRequest, id uuid.UUID) error
	DeleteAdmin(id uuid.UUID) error
}

type adminUsecase struct {
	repository       repositories.AdminRepository
	cloudinaryClient cloudinary.ICloudinaryClient
	passwordHelper   helpers.PasswordHelper
	tokenHelper      helpers.TokenHelper
}

func NewAdminUsecase(
	repository repositories.AdminRepository,
	cloudinaryClient cloudinary.ICloudinaryClient,
	passwordHelper helpers.PasswordHelper,
	tokenHelper helpers.TokenHelper,
) *adminUsecase {
	return &adminUsecase{
		repository:       repository,
		cloudinaryClient: cloudinaryClient,
		passwordHelper:   passwordHelper,
		tokenHelper:      tokenHelper,
	}
}

func (uc *adminUsecase) Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error) {
	admin, err := uc.repository.FindByUsername(request.Username)
	if err != nil {
		return nil, &errorHandlers.ConflictError{Message: "Akun tidak ditemukan"}
	}
	if err = uc.passwordHelper.VerifyPassword(admin.Password, request.Password); err != nil {
		return nil, &errorHandlers.ConflictError{Message: "Email atau password salah"}
	}

	accessToken, err := uc.tokenHelper.GenerateAccessToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	refreshToken, err := uc.tokenHelper.GenerateRefreshToken(admin)
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
	accessToken, err := uc.tokenHelper.GenerateAccessToken(admin)
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

	res := dto.ToGetAllAdminResponse(admins)

	return res, total, nil
}

func (uc *adminUsecase) GetAdminDetail(id uuid.UUID) (*dto.GetDetailAdminResponse, error) {
	admin, err := uc.repository.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, &errorHandlers.NotFoundError{Message: "Data admin tidak ditemukan"}
		default:
			return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan detail admin"}
		}
	}

	res := dto.ToGetDetailAdminResponse(admin)

	return res, nil
}

func (uc *adminUsecase) CreateAdmin(request *dto.AdminRequest) error {
	var profileImageURL *string

	if request.ProfileImage != nil {
		url, err := uc.cloudinaryClient.UploadImage(request.ProfileImage, "admin")
		if err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal menambah data admin"}
		}
		profileImageURL = &url
	}

	hashedPassword, err := uc.passwordHelper.HashPassword(request.Password)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menambah data admin"}
	}

	request.Password = hashedPassword
	admin := dto.ToCreateAdminRequest(request, profileImageURL)

	if err = uc.repository.Create(admin); err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return &errorHandlers.ConflictError{Message: "Username sudah digunakan"}
		default:
			return &errorHandlers.InternalServerError{Message: "Gagal menambah data admin"}
		}
	}

	return nil
}

func (uc *adminUsecase) UpdateAdmin(request *dto.UpdateAdminRequest, id uuid.UUID) error {
	var profileImageURL *string

	admin, err := uc.repository.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return &errorHandlers.NotFoundError{Message: "Data admin tidak ditemukan"}
		default:
			return &errorHandlers.InternalServerError{Message: "Gagal mengupdate data admin"}
		}
	}

	if request.Password == "" {
		request.Password = admin.Password
	} else {
		hashedPassword, hashErr := uc.passwordHelper.HashPassword(request.Password)
		if hashErr != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal mengubah data admin"}
		}

		request.Password = hashedPassword
	}

	if admin.ProfileImageURL != nil {
		if err = uc.cloudinaryClient.DeleteImage(*admin.ProfileImageURL); err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal mengupdate data admin"}
		}
	}

	if request.ProfileImage != nil {
		var url string
		url, err = uc.cloudinaryClient.UploadImage(request.ProfileImage, "admin")
		if err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal mengubah data admin"}
		}
		profileImageURL = &url
	}

	adminReq := dto.ToUpdateAdminRequest(request, admin, profileImageURL)

	if err = uc.repository.Update(adminReq); err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return &errorHandlers.ConflictError{Message: "Username sudah digunakan"}
		default:
			return &errorHandlers.InternalServerError{Message: "Gagal mengubah data admin"}
		}
	}

	return nil
}

func (uc *adminUsecase) DeleteAdmin(id uuid.UUID) error {
	admin, err := uc.repository.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return &errorHandlers.NotFoundError{Message: "Data admin tidak ditemukan"}
		default:
			return &errorHandlers.InternalServerError{Message: "Gagal menghapus data admin"}
		}
	}

	if admin.ProfileImageURL != nil {
		if err = uc.cloudinaryClient.DeleteImage(*admin.ProfileImageURL); err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal menghapus data admin"}
		}
	}

	if err = uc.repository.Delete(admin); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menghapus data admin"}
	}

	return nil
}
