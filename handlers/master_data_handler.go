package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MasterDataHandler struct {
	MasterDataUsecase usecases.IMasterDataUsecase
}

func NewMasterDataHandler(masterDataUsecase usecases.IMasterDataUsecase) *MasterDataHandler {
	return &MasterDataHandler{MasterDataUsecase: masterDataUsecase}
}

func (h *MasterDataHandler) GetCategories(ctx echo.Context) error {
	categories, err := h.MasterDataUsecase.GetCategories()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data kategori destinasi berhasil ditemukan",
		Data:       categories,
		IsPaginate: false,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *MasterDataHandler) GetFacilities(ctx echo.Context) error {
	facilities, err := h.MasterDataUsecase.GetFacilities()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data fasilitas destinasi berhasil ditemukan",
		Data:       facilities,
		IsPaginate: false,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *MasterDataHandler) GetProvinces(ctx echo.Context) error {
	provinces, err := h.MasterDataUsecase.GetProvinces()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data provinsi berhasil ditemukan",
		Data:       provinces,
		IsPaginate: false,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *MasterDataHandler) GetCities(ctx echo.Context) error {
	provinceId := ctx.QueryParam("province_id")

	cities, err := h.MasterDataUsecase.GetCities(provinceId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data kota berhasil ditemukan",
		Data:       cities,
		IsPaginate: false,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *MasterDataHandler) GetSubdistricts(ctx echo.Context) error {
	cityId := ctx.QueryParam("city_id")

	subdistricts, err := h.MasterDataUsecase.GetSubdistricts(cityId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data kecamatan berhasil ditemukan",
		Data:       subdistricts,
		IsPaginate: false,
	})

	return ctx.JSON(http.StatusOK, response)
}
