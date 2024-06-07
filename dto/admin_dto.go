package dto

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
