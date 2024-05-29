package errorHandlers

import (
	"net/http"

	"capstone/dto"
	"capstone/helpers"

	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	var statusCode int
	switch err.(type) {
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *UnAuthorizedError:
		statusCode = http.StatusUnauthorized
	case *ConflictError:
		statusCode = http.StatusConflict
	case *ForbiddenError:
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
