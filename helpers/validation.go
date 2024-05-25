package helpers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string
	Message string
}

func errorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", fieldError.Field())
	case "number":
		return fmt.Sprintf("Field %s must be a number", fieldError.Field())
	case "startswith":
		return fmt.Sprintf("Field %s must start with %s", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("Field %s must greater than  %s", fieldError.Field(), fieldError.Param())
	case "max":
		return fmt.Sprintf("Field %s must less than  %s", fieldError.Field(), fieldError.Param())
	case "len":
		return fmt.Sprintf("Field %s must be  %s characters", fieldError.Field(), fieldError.Param())
	case "email":
		return fmt.Sprintf("Field %s must be  a valid email", fieldError.Field())
	case "maxFileSize":
		return fmt.Sprintf("Field %s must be  an image", fieldError.Field())
	}

	return fieldError.Error()
}
func IsValidImageType(fileHeader *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}

	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	fileType := http.DetectContentType(buffer)
	if _, ok := allowedTypes[fileType]; !ok {
		return false
	}

	return true
}

func IsValidImageSize(fileHeader *multipart.FileHeader) bool {
	maxSize := 2 * 1024 * 1014
	fileSize := fileHeader.Size
	fmt.Println(fileSize, maxSize)
	return fileSize <= int64(maxSize)
}

func ValidateRequest(str interface{}) interface{} {
	validate := validator.New()

	if err := validate.Struct(str); err != nil {
		ve := err.(validator.ValidationErrors)
		errors := make([]ApiError, len(ve))
		for i, fieldError := range ve {
			errors[i] = ApiError{fieldError.Field(), errorMessage(fieldError)}
		}
		return errors
	}
	return nil
}