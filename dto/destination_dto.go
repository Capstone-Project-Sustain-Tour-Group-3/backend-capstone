package dto

import (
	"capstone/entities"

	"github.com/google/uuid"
)

type SearchDestinationsResponse struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"nama"`
	Url      *string   `json:"url_media"`
	Province string    `json:"provinsi"`
	City     string    `json:"kota"`
}

func ToSearchDestinationsResponse(destinations *[]entities.Destination) *[]SearchDestinationsResponse {
	responses := make([]SearchDestinationsResponse, len(*destinations))

	for idx, destination := range *destinations {
		var url *string

		if len(destination.DestinationMedias) > 0 {
			url = &destination.DestinationMedias[0].Url
		}

		responses[idx] = SearchDestinationsResponse{
			Id:       destination.Id,
			Name:     destination.Name,
			Url:      url,
			Province: destination.DestinationAddress.Province.Name,
			City:     destination.DestinationAddress.City,
		}
	}

	return &responses
}

type UrlImage struct {
	Id   uuid.UUID `json:"id_media"`
	Url  string    `json:"url_media"`
	Type string    `json:"tipe"`
}

type UrlVideo struct {
	Id    uuid.UUID `json:"id_media"`
	Url   string    `json:"url_media"`
	Type  string    `json:"tipe"`
	Title string    `json:"judul"`
}

type DestinationAddress struct {
	Province   string `json:"provinsi"`
	City       string `json:"kota"`
	Regency    string `json:"kabupaten"`
	StreetName string `json:"nama_jalan"`
	PostalCode string `json:"kode_pos"`
}

type Category struct {
	Id   uuid.UUID `json:"id_kategori"`
	Name string    `json:"nama"`
}

type Facility struct {
	Id   uuid.UUID `json:"id_fasilitas"`
	Name string    `json:"nama"`
	Url  string    `json:"url_logo"`
}

type DestinationSort string

const (
	Terbaru DestinationSort = "terbaru"
	Terlama DestinationSort = "terlama"
	Populer DestinationSort = "populer"
)

type DetailDestinationResponse struct {
	Id                 uuid.UUID           `json:"id_destinasi"`
	Name               string              `json:"nama_destinasi"`
	OpenTime           string              `json:"jam_buka"`
	CloseTime          string              `json:"jam_tutup"`
	EntryPrice         float64             `json:"harga_masuk"`
	Description        string              `json:"deskripsi"`
	DestinationAddress *DestinationAddress `json:"alamat_destinasi"`
	UrlImages          *[]UrlImage         `json:"url_gambar"`
	UrlVideos          *[]UrlVideo         `json:"url_video"`
	Categories         *[]Category         `json:"kategori"`
	Facilities         *[]Facility         `json:"fasilitas"`
}

func ToUrlImages(destination *entities.Destination) *[]UrlImage {
	var images []UrlImage
	for _, media := range destination.DestinationMedias {
		if media.Type == "image" {
			images = append(images, UrlImage{
				Id:   media.Id,
				Url:  media.Url,
				Type: media.Type,
			})
		}
	}
	return &images
}

func ToUrlVideos(destination *entities.Destination) *[]UrlVideo {
	var videos []UrlVideo
	for _, media := range destination.DestinationMedias {
		if media.Type == "video" {
			videos = append(videos, UrlVideo{
				Id:    media.Id,
				Url:   media.Url,
				Type:  media.Type,
				Title: media.Title,
			})
		}
	}
	return &videos
}

func ToCategories(destination *entities.Destination) *[]Category {
	var categories []Category
	for _, category := range *destination.Categories {
		categories = append(categories, Category{
			Id:   category.Id,
			Name: category.Category.Name,
		})
	}
	return &categories
}

func ToFacilities(destination *entities.Destination) *[]Facility {
	var facilities []Facility
	for _, facility := range *destination.Facilities {
		facilities = append(facilities, Facility{
			Id:   facility.Id,
			Name: facility.Facility.Name,
			Url:  facility.Facility.Url,
		})
	}
	return &facilities
}

func ToDetailDestinationResponse(destination *entities.Destination) *DetailDestinationResponse {
	return &DetailDestinationResponse{
		Id:          destination.Id,
		Name:        destination.Name,
		OpenTime:    destination.OpenTime,
		CloseTime:   destination.CloseTime,
		EntryPrice:  destination.EntryPrice,
		Description: destination.Description,
		DestinationAddress: &DestinationAddress{
			Province:   destination.DestinationAddress.Province.Name,
			City:       destination.DestinationAddress.City,
			Regency:    destination.DestinationAddress.Regency,
			StreetName: destination.DestinationAddress.StreetName,
			PostalCode: destination.DestinationAddress.PostalCode,
		},
		UrlImages:  ToUrlImages(destination),
		UrlVideos:  ToUrlVideos(destination),
		Categories: ToCategories(destination),
		Facilities: ToFacilities(destination),
	}
}
