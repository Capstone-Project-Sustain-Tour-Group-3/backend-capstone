package usecases

import (
	"fmt"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/externals/cloudinary"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type UserUsecase interface {
	FindById(id uuid.UUID) (*entities.User, error)
	FindAll(page, limit int, searchQuery string) (*[]entities.User, *int64, error)
	Create(request *dto.UserRequest) error
	Update(id uuid.UUID, request *dto.UserRequest) error
	Delete(id uuid.UUID) error
}

type userUsecase struct {
	repository       repositories.UserRepository
	cloudinaryClient cloudinary.ICloudinaryClient
}

func NewUserUsecase(repository repositories.UserRepository, client cloudinary.ICloudinaryClient) *userUsecase {
	return &userUsecase{repository: repository, cloudinaryClient: client}
}

func (uc *userUsecase) FindById(id uuid.UUID) (*entities.User, error) {
	user, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "User tidak ditemukan"}
	}
	return user, nil
}

func (uc *userUsecase) FindAll(page, limit int, searchQuery string) (*[]entities.User, *int64, error) {
	users, total, err := uc.repository.FindAll(page, limit, searchQuery)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: "Gagal untuk menampilkan data user"}
	}
	return users, total, nil
}

func (uc *userUsecase) Create(request *dto.UserRequest) error {
	existUsername, _ := uc.repository.FindByUsername(request.Username)
	if existUsername != nil {
		return &errorHandlers.ConflictError{Message: "Username sudah digunakan"}
	}

	existEmail, _ := uc.repository.FindByEmail(request.Email)
	if existEmail != nil {
		return &errorHandlers.ConflictError{Message: "Email sudah digunakan"}
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal hashing password"}
	}
	user := &entities.User{
		Id:              uuid.New(),
		Email:           request.Email,
		Password:        password,
		Username:        request.Username,
		Fullname:        request.NamaLengkap,
		Bio:             request.Bio,
		PhoneNumber:     request.NoTelepon,
		Gender:          request.JenisKelamin,
		City:            request.Kota,
		Province:        request.Provinsi,
		EmailVerifiedAt: nil,
	}

	if err = uc.handleProfilePictureUpdate(request, user); err != nil {
		return err
	}

	if err = uc.repository.Create(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk menambah data user"}
	}
	return nil
}

func (uc *userUsecase) Update(id uuid.UUID, request *dto.UserRequest) error {
	user, _ := uc.repository.FindById(id)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "User tidak ditemukan"}
	}

	if request.Username != user.Username {
		existUsername, _ := uc.repository.FindByUsername(request.Username)
		if existUsername != nil {
			return &errorHandlers.ConflictError{Message: "Username sudah digunakan"}
		}
	}

	if request.Email != user.Email {
		existEmail, _ := uc.repository.FindByEmail(request.Email)
		if existEmail != nil {
			return &errorHandlers.ConflictError{Message: "Email sudah digunakan"}
		}
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal hashing password"}
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = password
	user.Fullname = request.NamaLengkap
	user.Bio = request.Bio
	user.PhoneNumber = request.NoTelepon
	user.Gender = request.JenisKelamin
	user.City = request.Kota
	user.Province = request.Provinsi

	if err = uc.handleProfilePictureUpdate(request, user); err != nil {
		return err
	}

	if err = uc.repository.Update(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"}
	}
	return nil
}

func (uc *userUsecase) Delete(id uuid.UUID) error {
	user, _ := uc.repository.FindById(id)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "User tidak ditemukan"}
	}

	if user.ProfileImageUrl != nil {
		if err := uc.cloudinaryClient.DeleteImage(*user.ProfileImageUrl); err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal menghapus foto profil"}
		}
	}

	if err := uc.repository.Delete(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk menghapus data user"}
	}
	return nil
}

func (uc *userUsecase) handleProfilePictureUpdate(request *dto.UserRequest, user *entities.User) error {
	if request.FotoProfil == nil {
		return nil
	}

	if !helpers.IsValidImageType(request.FotoProfil) {
		return &errorHandlers.BadRequestError{Message: "Tipe foto profil tidak valid"}
	}
	if !helpers.IsValidImageSize(request.FotoProfil) {
		return &errorHandlers.BadRequestError{Message: "Ukuran foto profil tidak valid"}
	}

	file, err := request.FotoProfil.Open()
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal membuka file foto profil"}
	}
	defer file.Close()

	fmt.Println(user.ProfileImageUrl)
	if user.ProfileImageUrl != nil {
		if err = uc.cloudinaryClient.DeleteImage(*user.ProfileImageUrl); err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal menghapus foto profil"}
		}
	}

	urlMedia, err := uc.cloudinaryClient.UploadImage(file, "users")
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal upload foto profil"}
	}
	user.ProfileImageUrl = &urlMedia

	return nil
}
