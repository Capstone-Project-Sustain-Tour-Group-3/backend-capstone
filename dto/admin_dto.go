package dto

import (
	"time"

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
	CreatedAt time.Time `json:"tanggal_pembuatan"`
}

func ToLoginAdminResponse(response *LoginAdminResponse) *LoginAdminResponse {
	return &LoginAdminResponse{
		Username:     response.Username,
		ProfileImage: response.ProfileImage,
		Role:         response.Role,
		AccessToken:  response.AccessToken,
	}
}

type NewToken struct {
	AccessToken string `json:"access_token"`
}
