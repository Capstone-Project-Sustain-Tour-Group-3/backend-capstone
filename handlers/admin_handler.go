package handlers

import (
	"net/http"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"

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

	ctx.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    result.RefreshToken,
		Path:     "/",
		Domain:   "",
		MaxAge:   24 * 60 * 60,
		Secure:   true,
		HttpOnly: true,
	})
	adminResponse := dto.ToLoginAdminResponse(result)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login berhasil!",
		Data:       adminResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *adminHandler) Logout(ctx echo.Context) error {
	cookie, err := ctx.Cookie("refreshToken")
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "Token tidak ditemukan"})
	}
	if err = h.usecase.Logout(cookie.Value); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	ctx.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		Domain:   "",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	})
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Logout berhasil!",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *adminHandler) GetNewAccessToken(ctx echo.Context) error {
	cookie, err := ctx.Cookie("refreshToken")
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.ConflictError{Message: "Token tidak ditemukan"})
	}
	result, err := h.usecase.GetNewAccessToken(cookie.Value)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Token berhasil diperbarui!",
		Data:       result,
	})
	return ctx.JSON(http.StatusOK, response)
}
