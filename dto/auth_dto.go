package dto

type RegisterRequest struct {
	Username    string `json:"username" validate:"required,max=16"`
	NamaLengkap string `json:"nama_lengkap" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	NoTelepon   string `json:"no_telepon" validate:"required,number,startswith=08,min=11,max=13"`
	Password    string `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	ReferenceId string `json:"reference_id"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Username     string  `json:"username"`
	ProfileImage *string `json:"profile_image"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token,omitempty"`
}

type VerifyEmailRequest struct {
	RefId string `json:"ref_id"`
	OTP   string `json:"otp"`
}

type ResendOTPRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordRequest struct {
	RefId              string `json:"ref_id"`
	Password           string `json:"password" validate:"required,min=8"`
	KonfirmasiPassword string `json:"konfirmasi_password" validate:"required,min=8"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
