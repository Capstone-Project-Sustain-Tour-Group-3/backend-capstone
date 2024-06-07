package handlers

import (
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

type RouteHandler struct {
	RouteUsecase usecases.RouteInterface
}

func NewRouteHandler(routeUsecase usecases.RouteInterface) *RouteHandler {
	return &RouteHandler{RouteUsecase: routeUsecase}
}

func (h *RouteHandler) FindAll(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	searchQuery := ctx.QueryParam("search")
	routes, totalPtr, err := h.RouteUsecase.FindAll(page, limit, searchQuery)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	total := *totalPtr
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		page = lastPage
	}
	findAllResponse := dto.ToFindAllRouteResponse(routes)
	response := helpers.Response(dto.ResponseParams{
		StatusCode:  http.StatusOK,
		Message:     "data rute berhasil ditemukan",
		Data:        findAllResponse,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *RouteHandler) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "id tidak valid"})
	}
	route, err := h.RouteUsecase.FindById(uid)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	findResponse := dto.ToFindByIdRouteResponse(route)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data rute berhasil didapatkan",
		Data:       findResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *RouteHandler) DeleteRoute(ctx echo.Context) error {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "id tidak valid"})
	}
	err = h.RouteUsecase.DeleteRoute(uid)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data rute berhasil dihapus",
	})
	return ctx.JSON(http.StatusOK, response)
}