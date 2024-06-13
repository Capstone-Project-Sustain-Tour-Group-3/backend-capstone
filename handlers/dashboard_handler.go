package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type dashboardHandler struct {
	dashboardUsecase usecases.DashboardUsecase
}

func NewDashboardHandler(dashboardUsecase usecases.DashboardUsecase) *dashboardHandler {
	return &dashboardHandler{dashboardUsecase: dashboardUsecase}
}

func (h *dashboardHandler) GetAllData(c echo.Context) error {
	data, err := h.dashboardUsecase.GetAllData()
	if err != nil {
		return errorHandlers.HandleError(c, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil menampilkan data dashboard",
		Data:       data,
	})

	return c.JSON(http.StatusOK, response)
}
