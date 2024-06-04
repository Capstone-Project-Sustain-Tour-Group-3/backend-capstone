package handlers

import (
	"net/http"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type personalizationHandler struct {
	usecase usecases.PersonalizationUsecase
}

func NewPersonalizationHandler(usecase usecases.PersonalizationUsecase) *personalizationHandler {
	return &personalizationHandler{usecase}
}

func (h *personalizationHandler) GetProvinces(ctx echo.Context) error {
	provinces, err := h.usecase.GetProvinces()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan provinsi",
		Data:       provinces,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *personalizationHandler) GetCategories(ctx echo.Context) error {
	categories, err := h.usecase.GetCategories()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan kategori destinasi",
		Data:       categories,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *personalizationHandler) CreatePersonalization(ctx echo.Context) error {
	uid, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.UnAuthorizedError{Message: "User tidak dikenali"},
		)
	}

	var req dto.PersonalizationRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.BadRequestError{Message: err.Error()},
		)
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "Failed",
			Message: "Permintaan tidak valid. Silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	err := h.usecase.CreatePersonalization(&req, *uid)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Personalisasi berhasil di buat",
	})
	return ctx.JSON(http.StatusCreated, response)
}
