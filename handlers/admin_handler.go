package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type adminHandler struct {
	usecase usecases.AdminUsecase
}

func NewAdminHandler(uc usecases.AdminUsecase) *adminHandler {
	return &adminHandler{uc}
}

func (h *adminHandler) Login(ctx echo.Context) error {
	var req dto.LoginAdminRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: err.Error()})
	}
	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	result, err := h.usecase.Login(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "You have successfully logged in",
		Data:       result,
	})
	return ctx.JSON(http.StatusOK, response)
}