package errorHandlers

import (
	"errors"
	"net/http"

	"capstone/dto"
	"capstone/helpers"

	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	var statusCode int
	var badRequestError *BadRequestError
	var internalServerError *InternalServerError
	var notFoundError *NotFoundError
	var unAuthorizedError *UnAuthorizedError
	var conflictError *ConflictError
	var forbiddenError *ForbiddenError
	switch {
	case errors.As(err, &badRequestError):
		statusCode = http.StatusBadRequest
	case errors.As(err, &internalServerError):
		statusCode = http.StatusInternalServerError
	case errors.As(err, &notFoundError):
		statusCode = http.StatusNotFound
	case errors.As(err, &unAuthorizedError):
		statusCode = http.StatusUnauthorized
	case errors.As(err, &conflictError):
		statusCode = http.StatusConflict
	case errors.As(err, &forbiddenError):
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	return c.JSON(statusCode, response)
}
