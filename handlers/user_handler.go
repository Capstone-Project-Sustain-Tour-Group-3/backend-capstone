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

type userHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) *userHandler {
	return &userHandler{usecase}
}

func (h *userHandler) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	user, err := h.usecase.FindById(userId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	findResponse := dto.ToFindByIdResponse(user)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data user berhasil ditemukan",
		Data:       findResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) FindAll(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	searchQuery := ctx.QueryParam("search")
	users, totalPtr, err := h.usecase.FindAll(page, limit, searchQuery)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	total := *totalPtr
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		page = lastPage
	}
	findAllResponse := dto.ToFindAllUserResponse(users)
	response := helpers.Response(dto.ResponseParams{
		StatusCode:  http.StatusOK,
		Message:     "data user berhasil ditemukan",
		Data:        findAllResponse,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Create(ctx echo.Context) error {
	var req dto.UserRequest
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
	err := h.usecase.Create(&req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "data user berhasil ditambah",
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *userHandler) Update(ctx echo.Context) error {
	var req dto.UserRequest
	if err := ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	err = h.usecase.Update(userId, &req)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data user berhasil diperbarui",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	err = h.usecase.Delete(userId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data user berhasil dihapus",
	})
	return ctx.JSON(http.StatusOK, response)
}
