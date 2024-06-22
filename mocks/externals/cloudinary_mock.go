package externals

import (
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type MockCloudinaryClient struct {
	mock.Mock
}

func (m *MockCloudinaryClient) UploadImage(file multipart.File, folder string) (string, error) {
	args := m.Called(file, folder)
	return args.String(0), args.Error(1)
}

func (m *MockCloudinaryClient) DeleteImage(mediaUrl string) error {
	args := m.Called(mediaUrl)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
