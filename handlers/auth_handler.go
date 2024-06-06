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

type authHandler struct {
	usecase usecases.AuthUsecase
}

func NewAuthHandler(usecase usecases.AuthUsecase) *authHandler {
	return &authHandler{usecase}
}

func (h *authHandler) Register(ctx echo.Context) error {
	var req dto.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}
	refId, err := h.usecase.Register(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Registrasi berhasil!",
		Data:       refId,
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *authHandler) ResendOTP(ctx echo.Context) error {
	var req dto.ResendOTPRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	refId, err := h.usecase.ResendOTP(req.Email)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "OTP berhasil dikirimkan",
		Data:       refId,
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *authHandler) VerifyEmail(ctx echo.Context) error {
	var req dto.VerifyEmailRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}
	err := h.usecase.VerifyEmail(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "OTP Valid",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *authHandler) Login(ctx echo.Context) error {
	var req dto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
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
		Message:    "Login berhasil!",
		Data:       result,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *authHandler) Pong(ctx echo.Context) error {
	id, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
		})
	}
	return ctx.JSON(http.StatusOK, id)
}

func (h *authHandler) ForgotPassword(ctx echo.Context) error {
	var req dto.ForgotPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}
	err := h.usecase.ForgotPassword(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Password berhasil diubah",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *authHandler) Logout(ctx echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	err := h.usecase.Logout(req.RefreshToken)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Logout berhasil!",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *authHandler) GetNewAccessToken(ctx echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	token, err := h.usecase.GetNewAccessToken(req.RefreshToken)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan token baru",
		Data:       token,
	})
	return ctx.JSON(http.StatusOK, response)
}
