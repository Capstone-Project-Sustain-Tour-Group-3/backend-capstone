package dto

type ChatbotRequest struct {
	Message string `json:"message" validate:"required"`
}

type ChatbotResponse struct {
	Answer string `json:"answer"`
}
