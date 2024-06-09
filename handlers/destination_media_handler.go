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

type DestinationMediaHandler struct {
	usecase usecases.DestinationMediaUsecase
}

func NewDestinationMediaHandler(usecase usecases.DestinationMediaUsecase) *DestinationMediaHandler {
	return &DestinationMediaHandler{usecase}
}

func (h *DestinationMediaHandler) AllDestinationMedia(ctx echo.Context) error {
	destinationMedias, err := h.usecase.FindAll()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, destinationMedias)
}

func (h *DestinationMediaHandler) DetailMedia(ctx echo.Context) error {
	id := ctx.Param("id")
	destinationMediaId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	destinationMedia, err := h.usecase.FindById(destinationMediaId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data destination media berhasil ditemukan",
		Data:       destinationMedia,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationMediaHandler) Create(ctx echo.Context) error {
	var req dto.CreateDestinationMediaRequest
	err := ctx.Bind(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	if err = h.usecase.Create(req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "data konten berhasil ditambah",
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *DestinationMediaHandler) Update(ctx echo.Context) error {
	var req dto.UpdateDestinationMediaRequest

	id := ctx.Param("id")
	destinationMediaId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	if err = ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	if err = h.usecase.Update(destinationMediaId, req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data konten berhasil diubah",
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationMediaHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	destinationMediaId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	if err = h.usecase.Delete(destinationMediaId); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data konten berhasil dihapus",
	})

	return ctx.JSON(http.StatusOK, response)
}
