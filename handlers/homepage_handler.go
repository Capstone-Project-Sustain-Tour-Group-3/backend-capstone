package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type homePageHandler struct {
	usecase usecases.HomepageUsecase
}

func NewHomePageHandler(usecase usecases.HomepageUsecase) *homePageHandler {
	return &homePageHandler{usecase}
}

func (h *homePageHandler) GetHomepageData(ctx echo.Context) error {
	uid, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.UnAuthorizedError{Message: "User tidak dikenali"},
		)
	}

	cityName := strings.TrimSpace(ctx.QueryParam("provinsi"))
	if cityName == "" {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.BadRequestError{
				Message: "Nama kota tidak boleh kosong",
			},
		)
	}

	req := dto.ToHomepageRequest(cityName)

	data, err := h.usecase.GetHomepageData(req, *uid)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	res := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan data pada homepage",
		Data:       data,
	})

	return ctx.JSON(http.StatusOK, res)
}
