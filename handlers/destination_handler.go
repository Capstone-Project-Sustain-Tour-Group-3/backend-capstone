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

	if h.usecase.IncrementVisitCount(destinationId) != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	destination.VisitCount++

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

	// Bind basic fields
	req.Name = ctx.FormValue("nama_destinasi")
	req.Description = ctx.FormValue("deskripsi")
	req.OpenTime = ctx.FormValue("jam_buka")
	req.CloseTime = ctx.FormValue("jam_tutup")
	req.EntryPrice, _ = strconv.ParseFloat(ctx.FormValue("harga_masuk"), 64)
	req.CategoryId, _ = uuid.Parse(ctx.FormValue("id_kategori"))
	req.Latitude, _ = strconv.ParseFloat(ctx.FormValue("latitude"), 64)
	req.Longitude, _ = strconv.ParseFloat(ctx.FormValue("longitude"), 64)

	// Unmarshal nested JSON fields
	if err := json.Unmarshal([]byte(ctx.FormValue("fasilitas")), &req.FacilityIds); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "error",
			Message: "Invalid fasilitas",
			Errors:  err.Error(),
		})
	}

	if err := json.Unmarshal([]byte(ctx.FormValue("alamat_destinasi")), &req.DestinationAddress); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "error",
			Message: "Invalid alamat destinasi",
			Errors:  err.Error(),
		})
	}

	files, _ := ctx.MultipartForm()

	fileHeaders := files.File["files"]

	for i, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return errorHandlers.HandleError(ctx, err)
		}
		defer file.Close()

		var judul []string
		if err = json.Unmarshal([]byte(ctx.FormValue("judul")), &judul); err != nil {
			return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
				Status:  "error",
				Message: "Invalid judul",
				Errors:  err.Error(),
			})
		}

		req.DestinationImages = append(req.DestinationImages, dto.CreateDestinationImageRequest{
			File:  file, // atau simpan dalam format yang diperlukan
			Title: judul[i],
		})
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	reqJSON, _ := json.MarshalIndent(req, "", "  ")

	fmt.Println(string(reqJSON))

	err := h.usecase.CreateDestination(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "data destinasi berhasil ditambah",
		Data:       req,
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *DestinationHandler) UpdateDestination(ctx echo.Context) error {
	var req dto.UpdateDestinationRequest

	id := ctx.Param("id")
	destinationId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	// Bind basic fields
	req.Name = ctx.FormValue("nama_destinasi")
	req.Description = ctx.FormValue("deskripsi")
	req.OpenTime = ctx.FormValue("jam_buka")
	req.CloseTime = ctx.FormValue("jam_tutup")
	req.EntryPrice, _ = strconv.ParseFloat(ctx.FormValue("harga_masuk"), 64)
	req.CategoryId, _ = uuid.Parse(ctx.FormValue("id_kategori"))
	req.Latitude, _ = strconv.ParseFloat(ctx.FormValue("latitude"), 64)
	req.Longitude, _ = strconv.ParseFloat(ctx.FormValue("longitude"), 64)

	// Unmarshal nested JSON fields
	if err = json.Unmarshal([]byte(ctx.FormValue("fasilitas")), &req.FacilityIds); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "error",
			Message: "Invalid fasilitas",
			Errors:  err.Error(),
		})
	}

	if err = json.Unmarshal([]byte(ctx.FormValue("alamat_destinasi")), &req.DestinationAddress); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "error",
			Message: "Invalid alamat destinasi",
			Errors:  err.Error(),
		})
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

	err = h.usecase.UpdateDestination(destinationId, &req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data destinasi berhasil diubah",
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) DeleteDestination(ctx echo.Context) error {
	id := ctx.Param("id")
	destinationId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	err = h.usecase.DeleteDestination(destinationId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data destinasi berhasil dihapus",
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) GetCitiesWithDestinations(ctx echo.Context) error {
	cities, err := h.usecase.GetCitiesWithDestinations()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "berhasil menampilkan kota yang memiliki destinasi",
		Data:       cities,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) GetDestinationsByCityId(ctx echo.Context) error {
	id := ctx.Param("id")

	destinations, err := h.usecase.GetDestinationByCityId(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "berhasil menampilkan destinasi",
		Data:       destinations,
	})

	return ctx.JSON(http.StatusOK, response)
}
