package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func errorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("Kolom %s harus diisi", fieldError.Field())
	case "number":
		return fmt.Sprintf("Kolom %s harus sebuah angka", fieldError.Field())
	case "startswith":
		return fmt.Sprintf("Kolom %s harus dimulai dengan %s", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("Kolom %s harus lebih dari %s karakter", fieldError.Field(), fieldError.Param())
	case "max":
		return fmt.Sprintf("Kolom %s harus kurang dari %s karakter", fieldError.Field(), fieldError.Param())
	case "len":
		return fmt.Sprintf("Kolom %s harus memiliki panjang %s karakter", fieldError.Field(), fieldError.Param())
	case "email":
		return fmt.Sprintf("Kolom %s harus sebuah email yang valid", fieldError.Field())
	case "gte":
		return fmt.Sprintf("Kolom %s harus terisi dengan jumlah lebih dari sama dengan %s", fieldError.Field(), fieldError.Param())
	case "lte":
		return fmt.Sprintf("Kolom %s harus terisi dengan jumlah kurang dari sama dengan %s", fieldError.Field(), fieldError.Param())
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
	return fileSize <= int64(maxSize)
}

func ValidateRequest(str interface{}) interface{} {
	validate := validator.New()

	if err := validate.Struct(str); err != nil {
		var ve validator.ValidationErrors
		errors.As(err, &ve)
		apiErrors := make([]ApiError, len(ve))
		for i, fieldError := range ve {
			apiErrors[i] = ApiError{fieldError.Field(), errorMessage(fieldError)}
		}
		return apiErrors
	}
	return nil
}
