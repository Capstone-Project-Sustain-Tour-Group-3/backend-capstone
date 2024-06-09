package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

type DestinationMediaHandler struct {
	usecase usecases.IDestinationMediaUsecase
}

func NewDestinationMediaHandler(usecase usecases.IDestinationMediaUsecase) *DestinationMediaHandler {
	return &DestinationMediaHandler{usecase}
}

func (h *DestinationMediaHandler) AllDestinationMedia(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	searchQuery := ctx.QueryParam("search")

	totalPtr, destinationMedias, err := h.usecase.FindAll(page, limit, searchQuery)

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
		Message:     "berhasil menampilkan media destinasi",
		Data:        destinationMedias,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
		IsSort:      false,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationMediaHandler) DetailDestinationMedia(ctx echo.Context) error {
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

func (h *DestinationMediaHandler) CreateDestinationMedia(ctx echo.Context) error {
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

func (h *DestinationMediaHandler) UpdateDestinationMedia(ctx echo.Context) error {
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

func (h *DestinationMediaHandler) DeleteDestinationMedia(ctx echo.Context) error {
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
