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
	if args.Get(0) == "" {
		return "", args.Error(1)
	}
	return args.String(0), nil
}

func (m *MockCloudinaryClient) DeleteImage(mediaUrl string) error {
	args := m.Called(mediaUrl)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
