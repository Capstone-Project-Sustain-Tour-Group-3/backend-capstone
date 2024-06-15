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

type ProfileUsecase interface {
	GetDetailUser(id uuid.UUID) (*entities.User, error)
	InsertUserDetail(request *dto.UserDetailRequest, id uuid.UUID) error
	ChangePassword(request *dto.ChangePasswordRequest, id uuid.UUID) error
}

type profileUsecase struct {
	userRepo         repositories.UserRepository
	cloudinaryClient cloudinary.ICloudinaryClient
	passwordHelper   helpers.PasswordHelper
}

func NewProfileUsecase(
	repository repositories.UserRepository,
	client cloudinary.ICloudinaryClient,
	passwordHelper helpers.PasswordHelper,
) *profileUsecase {
	return &profileUsecase{userRepo: repository, cloudinaryClient: client, passwordHelper: passwordHelper}
}

func (uc *profileUsecase) GetDetailUser(id uuid.UUID) (*entities.User, error) {
	user, err := uc.userRepo.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "User tidak ditemukan"}
	}
	return user, nil
}

func (uc *profileUsecase) InsertUserDetail(request *dto.UserDetailRequest, id uuid.UUID) error {
	user, _ := uc.userRepo.FindById(id)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "User tidak ditemukan"}
	}
	if request.Email != user.Email {
		existEmail, _ := uc.userRepo.FindByEmail(request.Email)
		if existEmail != nil {
			return &errorHandlers.ConflictError{Message: "Email sudah digunakan"}
		}
	}

	user.Username = request.Username
	user.Fullname = request.NamaLengkap
	user.PhoneNumber = request.NoTelepon
	user.Bio = request.Bio
	user.Gender = request.JenisKelamin
	user.City = request.Kota
	user.Province = request.Provinsi
	user.Email = request.Email

	if err := uc.handleProfilePictureUpdate(request, user); err != nil {
		return err
	}

	if err := uc.userRepo.Update(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"}
	}
	return nil
}

func (uc *profileUsecase) handleProfilePictureUpdate(request *dto.UserDetailRequest, user *entities.User) error {
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

	if user.ProfileImageUrl != nil && *user.ProfileImageUrl != "" {
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

func (uc *profileUsecase) ChangePassword(request *dto.ChangePasswordRequest, id uuid.UUID) error {
	user, _ := uc.userRepo.FindById(id)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "User tidak ditemukan"}
	}

	if err := uc.passwordHelper.VerifyPassword(user.Password, request.PasswordLama); err != nil {
		return &errorHandlers.ConflictError{Message: "Password lama tidak valid"}
	}

	if request.PasswordBaru != request.KonfirmasiPassword {
		return &errorHandlers.BadRequestError{Message: "Password dan konfirmasi password tidak cocok"}
	}

	hashPassword, err := uc.passwordHelper.HashPassword(request.PasswordBaru)
	if err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal hashing password"}
	}
	user.Password = hashPassword

	if err = uc.userRepo.Update(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui password"}
	}
	return nil
}
