package dto

import (
	"mime/multipart"

	"capstone/entities"

	"github.com/google/uuid"
)

type UserRequest struct {
	Username     string                `form:"username" validate:"required,max=16"`
	Password     string                `form:"password" validate:"required,min=8"`
	NamaLengkap  string                `form:"nama_lengkap" validate:"required"`
	Email        string                `form:"email" validate:"required,email"`
	Bio          string                `form:"bio"`
	NoTelepon    string                `form:"no_telepon" validate:"required,number,startswith=08,min=11,max=13"`
	FotoProfil   *multipart.FileHeader `form:"foto_profil"`
	JenisKelamin string                `form:"jenis_kelamin"`
	Kota         string                `form:"kota"`
	Provinsi     string                `form:"provinsi"`
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
	NamaLengkap  string    `json:"nama_lengkap"`
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
		FotoProfil:   *user.ProfileImageUrl,
		JenisKelamin: user.Gender,
		Kota:         user.City,
		Provinsi:     user.Province,
	}
}

type UserDetailRequest struct {
	Username     string                `form:"username" validate:"required"`
	NamaLengkap  string                `form:"nama_lengkap" validate:"required"`
	FotoProfil   *multipart.FileHeader `form:"foto_profil"`
	Email        string                `form:"email" validate:"required,email"`
	NoTelepon    string                `form:"no_telepon" validate:"required,number,startswith=08,min=11,max=13"`
	Bio          string                `form:"bio" validate:"required"`
	JenisKelamin string                `form:"jenis_kelamin" validate:"required"`
	Kota         string                `form:"kota" validate:"required"`
	Provinsi     string                `form:"provinsi" validate:"required"`
}

type ChangePasswordRequest struct {
	PasswordLama       string `json:"password_lama" validate:"required"`
	PasswordBaru       string `json:"password_baru" validate:"required,min=8"`
	KonfirmasiPassword string `json:"konfirmasi_password" validate:"required,min=8"`
}

func ToFindAllUserResponse(user *[]entities.User) *[]findAllUserResponse {
	var users []findAllUserResponse
	for _, u := range *user {
		users = append(users, findAllUserResponse{
			Id:           u.Id,
			Username:     u.Username,
			NamaLengkap:  u.Fullname,
			Email:        u.Email,
			NoTelepon:    u.PhoneNumber,
			JenisKelamin: u.Gender,
			Kota:         u.City,
			Provinsi:     u.Province,
		})
	}
	return &users
}
