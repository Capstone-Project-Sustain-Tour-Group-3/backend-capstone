package handlers

import (
	"errors"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type adminHandler struct {
	usecase         usecases.AdminUsecase
	imageValidation helpers.IValidationHelper
}

func NewAdminHandler(uc usecases.AdminUsecase, iVal helpers.IValidationHelper) *adminHandler {
	return &adminHandler{uc, iVal}
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

func (h *adminHandler) GetAllAdmins(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 || err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 || err != nil {
		limit = 10
	}
	searchQuery := ctx.QueryParam("search")

	res, totalPtr, err := h.usecase.GetAllAdmins(page, limit, searchQuery)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	total := *totalPtr
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		page = lastPage
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode:  http.StatusOK,
		Message:     "Berhasil mendapatkan data admin",
		Data:        res,
		IsPaginate:  true,
		Total:       total,
		PerPage:     len(*res),
		CurrentPage: page,
		LastPage:    lastPage,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *adminHandler) GetAdminDetail(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.BadRequestError{
				Message: "Id admin tidak valid",
			},
		)
	}

	result, err := h.usecase.GetAdminDetail(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan detail admin",
		Data:       result,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *adminHandler) CreateAdmin(ctx echo.Context) error {
	var file multipart.File = nil
	req := new(dto.AdminRequest)

	if err := ctx.Bind(req); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "Permintaan tidak valid. Silakan periksa kembali data yang anda masukkan."})
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "Failed",
			Message: "Permintaan tidak valid. Silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	fileHeader, err := ctx.FormFile("foto_profil")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "File tidak dapat di baca"})
	}

	if fileHeader != nil {
		if !h.imageValidation.IsValidImageType(fileHeader) || !h.imageValidation.IsValidImageSize(fileHeader) {
			return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "File harus berupa gambar dan kurang dari 2 MB"})
		}

		file, err = fileHeader.Open()
		if err != nil {
			return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "File tidak dapat di baca"})
		}
		defer file.Close()
	}

	req.ProfileImage = file
	err = h.usecase.CreateAdmin(req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil menambah data admin",
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *adminHandler) UpdateAdmin(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.BadRequestError{
				Message: "Id admin tidak valid",
			},
		)
	}

	var file multipart.File = nil
	req := new(dto.UpdateAdminRequest)

	if err = ctx.Bind(req); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "Permintaan tidak valid. Silakan periksa kembali data yang anda masukkan."})
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "Permintaan tidak valid. Silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	fileHeader, err := ctx.FormFile("foto_profil")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "File tidak dapat di baca"})
	}

	if fileHeader != nil {
		if !h.imageValidation.IsValidImageType(fileHeader) || !h.imageValidation.IsValidImageSize(fileHeader) {
			return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "File harus berupa gambar dan kurang dari 2 MB"})
		}

		file, err = fileHeader.Open()
		if err != nil {
			return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "File tidak dapat di baca"})
		}
		defer file.Close()
	}

	req.ProfileImage = file
	err = h.usecase.UpdateAdmin(req, id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mengubah data admin",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *adminHandler) DeleteAdmin(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.BadRequestError{
				Message: "Id admin tidak valid",
			},
		)
	}

	err = h.usecase.DeleteAdmin(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil menghapus data admin",
	})
	return ctx.JSON(http.StatusOK, response)
}
