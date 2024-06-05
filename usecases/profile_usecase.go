package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type ProfileUsecase interface {
	GetDetailUser(id uuid.UUID) (*entities.User, error)
	InsertUserDetail(request *dto.UserDetailRequest, id uuid.UUID) (string, error)
}

type profileUsecase struct {
	userRepo  repositories.UserRepository
	cacheRepo *repositories.CacheRepository
}

func NewProfileUsecase(repository repositories.UserRepository, cacheRepository *repositories.CacheRepository) *profileUsecase {
	return &profileUsecase{userRepo: repository, cacheRepo: cacheRepository}
}

func (uc *profileUsecase) GetDetailUser(id uuid.UUID) (*entities.User, error) {
	user, err := uc.userRepo.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "User tidak ditemukan"}
	}
	return user, nil
}

func (uc *profileUsecase) InsertUserDetail(request *dto.UserDetailRequest, id uuid.UUID) (string, error) {
	user, _ := uc.userRepo.FindById(id)
	if user == nil {
		return "", &errorHandlers.ConflictError{Message: "User tidak ditemukan"}
	}

	user.Username = request.Username
	user.Fullname = request.NamaLengkap
	user.PhoneNumber = request.NoTelepon
	user.Bio = request.Bio
	user.Gender = request.JenisKelamin
	user.City = request.Kota
	user.Province = request.Provinsi

	if request.Email != "" {
		isExist, _ := uc.userRepo.FindByEmail(request.Email)
		if isExist != nil {
			return "", &errorHandlers.ConflictError{Message: "Email sudah digunakan"}
		}
		otp := helpers.GenerateOTP()
		ref := helpers.GenerateReferenceId()
		uc.cacheRepo.Set(ref, otp)
		uc.cacheRepo.Set(ref+"_email", request.Email)

		if err := helpers.SendOTP(request.Email, request.NamaLengkap, otp); err != nil {
			return "", &errorHandlers.InternalServerError{Message: "Gagal mengirim OTP"}
		}
	}

	if request.FotoProfil != nil {
		if imageType := helpers.IsValidImageType(request.FotoProfil); !imageType {
			return "", &errorHandlers.BadRequestError{Message: "Tipe foto profil tidak valid"}
		}
		if size := helpers.IsValidImageSize(request.FotoProfil); !size {
			return "", &errorHandlers.BadRequestError{Message: "Ukuran foto profil tidak valid"}
		}
		user.ProfileImageUrl = "testing.jpg"
	}

	if err := uc.userRepo.Update(user); err != nil {
		return "", &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"}
	}
	return "", nil
}
