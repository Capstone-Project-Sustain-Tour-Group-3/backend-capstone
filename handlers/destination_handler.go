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

func (h *DestinationHandler) CreateDestination(ctx echo.Context) error {
	var req dto.CreateDestinationRequest

	form, _ := ctx.MultipartForm()

	err := json.Unmarshal([]byte(form.Value["info"][0]), &req.DestinationInfo)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	err = json.Unmarshal([]byte(form.Value["alamat_destinasi"][0]), &req.DestinationAddress)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	for i, fileHeader := range form.File["files"] {
		file, err := fileHeader.Open()
		if err != nil {
			return errorHandlers.HandleError(ctx, err)
		}
		defer file.Close()

		req.DestinationImages = append(req.DestinationImages, dto.CreateDestinationImageRequest{
			File:  file,
			Title: form.Value["judul"][i],
		})
	}

	// fmt.Println(ctx.FormValue("alamat_destinasi"))

	// if err := ctx.Bind(&req); err != nil {
	// 	fmt.Println("masuk siniii")
	// 	return errorHandlers.HandleError(ctx, err)
	// }

	reqJSON, _ := json.MarshalIndent(req, "", "  ")

	fmt.Println(string(reqJSON))

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	err = h.usecase.CreateDestination(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "data destinasi berhasil ditambah",
	})

	return ctx.JSON(http.StatusCreated, response)
}
