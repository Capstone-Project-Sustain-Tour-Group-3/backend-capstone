package dto

type LoginAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginAdminResponse struct {
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
	Role         string `json:"role"`
	Token        string `json:"token"`
}
