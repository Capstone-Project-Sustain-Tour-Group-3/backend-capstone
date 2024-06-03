package cloudinary

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type ICloudinaryClient interface {
	UploadImage(file multipart.File) (string, error)
	DeleteImage(mediaUrl string) error
}

type CloudinaryClient struct {
	client *cloudinary.Cloudinary
}

func NewCloudinaryClient(apiUrl string) *CloudinaryClient {
	client, err := cloudinary.NewFromURL(apiUrl)
	if err != nil {
		log.Fatal("cannot create cloudinary client :: ", err)
	}

	return &CloudinaryClient{client: client}
}

func (c *CloudinaryClient) UploadImage(file multipart.File) (string, error) {
	ctx := context.Background()

	result, err := c.client.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "tourease"})
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}

func (c *CloudinaryClient) DeleteImage(mediaUrl string) error {
	ctx := context.Background()

	parts := strings.Split(mediaUrl, "/")
	publicIDWithExtension := parts[len(parts)-2] + "/" + parts[len(parts)-1]
	publicID := strings.TrimSuffix(publicIDWithExtension, filepath.Ext(publicIDWithExtension))

	_, err := c.client.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})
	if err != nil {
		return fmt.Errorf("gagal menghapus media dengan ID '%s': %w", publicID, err)
	}

	return nil
}
