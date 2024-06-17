package externals

import (
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type MockImageValidator struct {
	mock.Mock
}

func (m *MockImageValidator) IsValidImageType(fileHeader *multipart.FileHeader) bool {
	args := m.Called(fileHeader)
	return args.Bool(0)
}

func (m *MockImageValidator) IsValidImageSize(fileHeader *multipart.FileHeader) bool {
	args := m.Called(fileHeader)
	return args.Bool(0)
}
