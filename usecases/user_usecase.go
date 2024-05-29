package usecases

import (
	"fmt"
	"time"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
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
	repository repositories.UserRepository
}

func NewUserUsecase(repository repositories.UserRepository) *userUsecase {
	return &userUsecase{repository}
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
		ProfileImageUrl: request.FotoProfil,
		Gender:          request.JenisKelamin,
		City:            request.Kota,
		Province:        request.Provinsi,
		EmailVerifiedAt: nil,
	}
	if err := uc.repository.Create(user); err != nil {
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
	user.ProfileImageUrl = request.FotoProfil
	user.Gender = request.JenisKelamin
	user.City = request.Kota
	user.Province = request.Provinsi

	if err := uc.repository.Update(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk memperbarui data user"}
	}
	return nil
}

func (uc *userUsecase) Delete(id uuid.UUID) error {
	user, _ := uc.repository.FindById(id)
	if user == nil {
		return &errorHandlers.ConflictError{Message: "User tidak ditemukan"}
	}

	if err := uc.repository.Delete(user); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal untuk menghapus data user"}
	}
	return nil
}
