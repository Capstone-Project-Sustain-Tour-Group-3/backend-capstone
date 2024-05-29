package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DestinationHandler struct {
	usecase usecases.IDestinationUsecase
}

func NewDestinationHandler(usecase usecases.IDestinationUsecase) *DestinationHandler {
	return &DestinationHandler{usecase}
}

func (h *DestinationHandler) SearchDestinations(ctx echo.Context) error {
	destinations, err := h.usecase.SearchDestinations()
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
