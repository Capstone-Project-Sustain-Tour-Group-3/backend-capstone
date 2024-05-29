package dto

type ChatbotRequest struct {
	Question string `json:"question" validate:"required"`
}

type ChatbotResponse struct {
	Answer string `json:"answer"`
}
