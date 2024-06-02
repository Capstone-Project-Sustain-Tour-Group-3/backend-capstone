package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type profileHandler struct {
	usecase usecases.ProfileUsecase
}

func NewProfileHandler(uc usecases.ProfileUsecase) *profileHandler {
	return &profileHandler{usecase: uc}
}

func (h *profileHandler) GetDetailUser(ctx echo.Context) error {
	uid, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return errorHandlers.HandleError(ctx, &errorHandlers.UnAuthorizedError{Message: "user tidak dikenali"})
	}

	user, err := h.usecase.GetDetailUser(*uid)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	findResponse := dto.ToFindByIdResponse(user)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data profil berhasil didapatkan",
		Data:       findResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}
