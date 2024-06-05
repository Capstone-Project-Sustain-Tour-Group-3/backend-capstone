package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DestinationHandler struct {
	usecase usecases.IDestinationUsecase
}

func NewDestinationHandler(usecase usecases.IDestinationUsecase) *DestinationHandler {
	return &DestinationHandler{usecase}
}

func (h *DestinationHandler) SearchDestinations(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	searchQuery := ctx.QueryParam("search")
	sortQuery := ctx.QueryParam("sort")
	filterQuery := ctx.QueryParam("filter")

	categoryName, totalPtr, destinations, err := h.usecase.SearchDestinations(page, limit, searchQuery, sortQuery, filterQuery)
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
		Message:     "berhasil menampilkan destinasi",
		Data:        destinations,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
		IsSort:      true,
		Sort:        sortQuery,
		IsFilter:    true,
		Filter:      categoryName,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) DetailDestination(ctx echo.Context) error {
	id := ctx.Param("id")
	destinationId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	destination, err := h.usecase.DetailDestination(destinationId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "berhasil menampilkan destinasi",
		Data:       destination,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) GetAllDestinations(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	searchQuery := ctx.QueryParam("search")

	totalPtr, destinations, err := h.usecase.GetAllDestinations(page, limit, searchQuery)
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
		Message:     "berhasil menampilkan destinasi",
		Data:        destinations,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) GetDestinationById(ctx echo.Context) error {
	id := ctx.Param("id")
	destinationId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	destination, err := h.usecase.GetDestinationById(destinationId)

	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "berhasil menampilkan destinasi",
		Data:       destination,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) CreateDestination(ctx echo.Context) error {
	var req dto.CreateDestinationRequest

	if err := ctx.Bind(&req); err != nil {
		fmt.Println("masuk siniii")
		return errorHandlers.HandleError(ctx, err)
	}

	reqJSON, _ := json.MarshalIndent(req, "", "  ")

	fmt.Println(string(reqJSON))

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	err := h.usecase.CreateDestination(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "data destinasi berhasil ditambah",
	})

	return ctx.JSON(http.StatusCreated, response)
}
