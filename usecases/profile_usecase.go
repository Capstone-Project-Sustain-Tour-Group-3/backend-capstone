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
	HandleProfilePictureUpdate(request *dto.UserDetailRequest, user *entities.User) error
}

type profileUsecase struct {
	userRepo         repositories.UserRepository
	cloudinaryClient cloudinary.ICloudinaryClient
	passwordHelper   helpers.PasswordHelper
	iValidation      helpers.IValidationHelper
}

func NewProfileUsecase(
	repository repositories.UserRepository,
	client cloudinary.ICloudinaryClient,
	passwordHelper helpers.PasswordHelper,
	iValidation helpers.IValidationHelper,
) *profileUsecase {
	return &profileUsecase{userRepo: repository, cloudinaryClient: client, passwordHelper: passwordHelper, iValidation: iValidation}
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

	if err := uc.HandleProfilePictureUpdate(request, user); err != nil {
		return err
	}

	if err := uc.userRepo.Update(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"}
	}
	return nil
}

func (uc *profileUsecase) HandleProfilePictureUpdate(request *dto.UserDetailRequest, user *entities.User) error {
	if request.FotoProfil == nil {
		return nil
	}

	if !uc.iValidation.IsValidImageType(request.FotoProfil) {
		return &errorHandlers.BadRequestError{Message: "Tipe foto profil tidak valid"}
	}
	if !uc.iValidation.IsValidImageSize(request.FotoProfil) {
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
