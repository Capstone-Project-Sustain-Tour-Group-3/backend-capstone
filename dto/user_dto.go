package dto

type RegisterRequest struct {
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	NoTelepon   string `json:"no_telepon"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	ReferenceId string `json:"reference_id"`
}

type VerifyEmailRequest struct {
	RefId string `json:"ref_id"`
	OTP   string `json:"otp"`
}
