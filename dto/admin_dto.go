package dto

import (
	"mime/multipart"

	"capstone/entities"

	"github.com/google/uuid"
)

type LoginAdminRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginAdminResponse struct {
	Username     string  `json:"username"`
	ProfileImage *string `json:"profile_image"`
	Role         string  `json:"role"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token,omitempty"`
}

type GetAllAdminResponse struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt string    `json:"tanggal_pembuatan"`
}

type GetDetailAdminResponse struct {
	Username        string  `json:"username"`
	ProfileImageURL *string `json:"foto_profil"`
	CreatedAt       string  `json:"tanggal_pembuatan"`
}

type AdminRequest struct {
	Username     string `form:"username" validate:"required,max=16"`
	Password     string `form:"password" validate:"required,min=8"`
	ProfileImage multipart.File
}

type UpdateAdminRequest struct {
	Username     string `form:"username" validate:"required,max=16"`
	Password     string `form:"password" validate:"omitempty,min=8"`
	ProfileImage multipart.File
}

func ToCreateAdminRequest(request *AdminRequest, imageURL *string) *entities.Admin {
	return &entities.Admin{
		Id:              uuid.New(),
		Username:        request.Username,
		Password:        request.Password,
		ProfileImageURL: imageURL,
		Role:            "admin",
		DeleteMilli:     0,
	}
}

func ToUpdateAdminRequest(request *UpdateAdminRequest, admin *entities.Admin, imageURL *string) *entities.Admin {
	return &entities.Admin{
		Id:              admin.Id,
		Username:        request.Username,
		Password:        request.Password,
		ProfileImageURL: imageURL,
		Role:            admin.Role,
		RefreshToken:    admin.RefreshToken,
		CreatedAt:       admin.CreatedAt,
		DeletedAt:       admin.DeletedAt,
		DeleteMilli:     admin.DeleteMilli,
	}
}

func ToLoginAdminResponse(response *LoginAdminResponse) *LoginAdminResponse {
	return &LoginAdminResponse{
		Username:     response.Username,
		ProfileImage: response.ProfileImage,
		Role:         response.Role,
		AccessToken:  response.AccessToken,
	}
}

func ToGetDetailAdminResponse(response *entities.Admin) *GetDetailAdminResponse {
	return &GetDetailAdminResponse{
		Username:        response.Username,
		ProfileImageURL: response.ProfileImageURL,
		CreatedAt:       response.CreatedAt.Format("02-01-2006"),
	}
}

func ToGetAllAdminResponse(response *[]entities.Admin) *[]GetAllAdminResponse {
	res := make([]GetAllAdminResponse, len(*response))

	for i := range *response {
		res[i] = GetAllAdminResponse{
			Id:        (*response)[i].Id,
			Username:  (*response)[i].Username,
			CreatedAt: (*response)[i].CreatedAt.Format("02-01-2006"),
		}
	}

	return &res
}

type NewToken struct {
	AccessToken string `json:"access_token"`
}
