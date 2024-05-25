package handlers

import (
	"fmt"
	"net/http"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) *userHandler {
	return &userHandler{usecase}
}

func (h *userHandler) Register(ctx echo.Context) error {
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
	fmt.Println(refId)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Registrasi berhasil!",
		Data:       refId,
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *userHandler) ResendOTP(ctx echo.Context) error {
	refId, err := h.usecase.ResendOTP()
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

func (h *userHandler) VerifyEmail(ctx echo.Context) error {
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

func (h *userHandler) Login(ctx echo.Context) error {
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
	loginResponse, err := h.usecase.Login(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login berhasil!",
		Data:       loginResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Pong(ctx echo.Context) error {
	id := ctx.Get("userId").(*uuid.UUID)
	return ctx.JSON(http.StatusOK, id)
}
