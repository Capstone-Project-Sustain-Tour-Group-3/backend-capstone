package dto

import (
	"capstone/entities"

	"github.com/google/uuid"
)

type UserRequest struct {
	Username     string `json:"username" validate:"required,max=16"`
	Password     string `json:"password" validate:"required,min=8"`
	NamaLengkap  string `json:"nama_lengkap" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Bio          string `json:"bio" validate:"required"`
	NoTelepon    string `json:"no_telepon" validate:"required,number,startswith=08,min=11,max=13"`
	FotoProfil   string `json:"foto_profil" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Kota         string `json:"kota" validate:"required"`
	Provinsi     string `json:"provinsi" validate:"required"`
}

type findByIdResponse struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	NamaLengkap  string    `json:"nama_lengkap"`
	Bio          string    `json:"bio"`
	Email        string    `json:"email"`
	NoTelepon    string    `json:"no_telepon"`
	FotoProfil   string    `json:"foto_profil"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Kota         string    `json:"kota"`
	Provinsi     string    `json:"provinsi"`
}

type findAllUserResponse struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	NoTelepon    string    `json:"no_telepon"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Kota         string    `json:"kota"`
	Provinsi     string    `json:"provinsi"`
}

func ToFindByIdResponse(user *entities.User) *findByIdResponse {
	return &findByIdResponse{
		Id:           user.Id,
		Username:     user.Username,
		NamaLengkap:  user.Fullname,
		Bio:          user.Bio,
		Email:        user.Email,
		NoTelepon:    user.PhoneNumber,
		FotoProfil:   user.ProfileImageUrl,
		JenisKelamin: user.Gender,
		Kota:         user.City,
		Provinsi:     user.Province,
	}
}

type UserDetailRequest struct{}

func ToFindAllUserResponse(user *[]entities.User) *[]findAllUserResponse {
	var users []findAllUserResponse
	for _, u := range *user {
		users = append(users, findAllUserResponse{
			Id:           u.Id,
			Username:     u.Username,
			Email:        u.Email,
			NoTelepon:    u.PhoneNumber,
			JenisKelamin: u.Gender,
			Kota:         u.City,
			Provinsi:     u.Province,
		})
	}
	return &users
}
