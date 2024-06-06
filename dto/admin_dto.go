package dto

import (
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
