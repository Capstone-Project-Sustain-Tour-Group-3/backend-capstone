package dto

import (
	"mime/multipart"

	"capstone/entities"

	"github.com/google/uuid"
)

type SearchDestination struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"nama"`
	Url        *string   `json:"url_media"`
	Province   string    `json:"provinsi"`
	City       string    `json:"kota"`
	VisitCount int       `json:"visit_count"`
	Category   Category  `json:"kategori"`
}

type GetAllDestination struct {
	Id                 uuid.UUID           `json:"id"`
	Name               string              `json:"nama"`
	OpenTime           string              `json:"jam_buka"`
	CloseTime          string              `json:"jam_tutup"`
	EntryPrice         float64             `json:"harga_masuk"`
	VisitCount         int                 `json:"visit_count"`
	Category           Category            `json:"kategori"`
	DestinationAddress *DestinationAddress `json:"alamat"`
}

func ToSearchDestinationsResponse(destinations *[]entities.Destination) *[]SearchDestination {
	responses := make([]SearchDestination, len(*destinations))

	for idx, destination := range *destinations {
		var url *string

		if len(destination.DestinationMedias) > 0 {
			url = &destination.DestinationMedias[0].Url
		}

		responses[idx] = SearchDestination{
			Id:         destination.Id,
			Name:       destination.Name,
			Url:        url,
			Province:   destination.DestinationAddress.Province.Name,
			City:       destination.DestinationAddress.City.Name,
			VisitCount: destination.VisitCount,
			Category:   *ToCategory(&destination),
		}
	}

	return &responses
}

func ToGetAllDestinationsResponse(destinations *[]entities.Destination) *[]GetAllDestination {
	responses := make([]GetAllDestination, len(*destinations))

	for idx, destination := range *destinations {
		responses[idx] = GetAllDestination{
			Id:         destination.Id,
			Name:       destination.Name,
			OpenTime:   destination.OpenTime,
			CloseTime:  destination.CloseTime,
			VisitCount: destination.VisitCount,
			EntryPrice: destination.EntryPrice,
			Category:   *ToCategory(&destination),
			DestinationAddress: &DestinationAddress{
				Province:    destination.DestinationAddress.Province.Name,
				City:        destination.DestinationAddress.City.Name,
				Subdistrict: destination.DestinationAddress.Subdistrict.Name,
				StreetName:  destination.DestinationAddress.StreetName,
				PostalCode:  destination.DestinationAddress.PostalCode,
			},
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
	Province    string `json:"provinsi"`
	City        string `json:"kota"`
	Subdistrict string `json:"kecamatan"`
	StreetName  string `json:"nama_jalan"`
	PostalCode  string `json:"kode_pos"`
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

type DestinationFilter string

const (
	Alam          DestinationFilter = "Alam"
	SeniDanBudaya DestinationFilter = "Seni dan Budaya"
	Sejarah       DestinationFilter = "Sejarah"
)

type DetailDestinationResponse struct {
	Id                 uuid.UUID            `json:"id_destinasi"`
	Name               string               `json:"nama_destinasi"`
	OpenTime           string               `json:"jam_buka"`
	CloseTime          string               `json:"jam_tutup"`
	EntryPrice         float64              `json:"harga_masuk"`
	VisitCount         int                  `json:"visit_count"`
	Description        string               `json:"deskripsi"`
	DestinationAddress *DestinationAddress  `json:"alamat_destinasi"`
	UrlImages          *[]UrlImage          `json:"url_gambar"`
	UrlVideos          *[]UrlVideo          `json:"url_video"`
	Categories         *Category            `json:"kategori"`
	Facilities         *[]Facility          `json:"fasilitas"`
	SimilarDestination *[]SearchDestination `json:"destinasi_serupa"`
}

type GetByIdDestinationResponse struct {
	Id                 uuid.UUID           `json:"id_destinasi"`
	Name               string              `json:"nama_destinasi"`
	OpenTime           string              `json:"jam_buka"`
	CloseTime          string              `json:"jam_tutup"`
	EntryPrice         float64             `json:"harga_masuk"`
	VisitCount         int                 `json:"visit_count"`
	Description        string              `json:"deskripsi"`
	DestinationAddress *DestinationAddress `json:"alamat_destinasi"`
	UrlImages          *[]UrlImage         `json:"url_gambar"`
	Categories         *Category           `json:"kategori"`
	Facilities         *[]Facility         `json:"fasilitas"`
}

type Cities struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type DestinationsByCity struct {
	Id            string `json:"id"`
	NamaDestinasi string `json:"nama_destinasi"`
	NamaJalan     string `json:"nama_jalan"`
	Kecamatan     string `json:"kecamatan"`
	Kota          string `json:"kota"`
	Provinsi      string `json:"provinsi"`
}

func ToDestinationsByCityResponse(destinations *[]entities.Destination) *[]DestinationsByCity {
	responses := make([]DestinationsByCity, len(*destinations))

	for idx, destination := range *destinations {
		responses[idx] = DestinationsByCity{
			Id:            destination.Id.String(),
			NamaDestinasi: destination.Name,
			NamaJalan:     destination.DestinationAddress.StreetName,
			Kecamatan:     destination.DestinationAddress.Subdistrict.Name,
			Kota:          destination.DestinationAddress.City.Name,
			Provinsi:      destination.DestinationAddress.Province.Name,
		}
	}

	return &responses
}

func ToCitiesResponse(cities *[]entities.City) *[]Cities {
	responses := make([]Cities, len(*cities))

	for idx, city := range *cities {
		responses[idx] = Cities{
			Id:   city.Id,
			Nama: city.Name,
		}
	}

	return &responses
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

func ToCategory(destination *entities.Destination) *Category {
	return &Category{
		Id:   destination.CategoryId,
		Name: destination.Category.Name,
	}
}

func ToFacilities(destination *entities.Destination) *[]Facility {
	var facilities []Facility
	for _, destinationFacility := range *destination.DestinationFacilities {
		facilities = append(facilities, Facility{
			Id:   destinationFacility.FacilityId,
			Name: destinationFacility.Facility.Name,
			Url:  destinationFacility.Facility.Url,
		})
	}
	return &facilities
}

func ToDetailDestinationResponse(destination *entities.Destination, similarDestinations *[]entities.Destination) *DetailDestinationResponse {
	return &DetailDestinationResponse{
		Id:          destination.Id,
		Name:        destination.Name,
		OpenTime:    destination.OpenTime,
		CloseTime:   destination.CloseTime,
		EntryPrice:  destination.EntryPrice,
		VisitCount:  destination.VisitCount,
		Description: destination.Description,
		DestinationAddress: &DestinationAddress{
			Province:    destination.DestinationAddress.Province.Name,
			City:        destination.DestinationAddress.City.Name,
			Subdistrict: destination.DestinationAddress.Subdistrict.Name,
			StreetName:  destination.DestinationAddress.StreetName,
			PostalCode:  destination.DestinationAddress.PostalCode,
		},
		UrlImages:          ToUrlImages(destination),
		UrlVideos:          ToUrlVideos(destination),
		Categories:         ToCategory(destination),
		Facilities:         ToFacilities(destination),
		SimilarDestination: ToSearchDestinationsResponse(similarDestinations),
	}
}

func ToGetByIdDestinationResponse(destination *entities.Destination) *GetByIdDestinationResponse {
	return &GetByIdDestinationResponse{
		Id:          destination.Id,
		Name:        destination.Name,
		OpenTime:    destination.OpenTime,
		CloseTime:   destination.CloseTime,
		EntryPrice:  destination.EntryPrice,
		VisitCount:  destination.VisitCount,
		Description: destination.Description,
		DestinationAddress: &DestinationAddress{
			Province:    destination.DestinationAddress.Province.Name,
			City:        destination.DestinationAddress.City.Name,
			Subdistrict: destination.DestinationAddress.Subdistrict.Name,
			StreetName:  destination.DestinationAddress.StreetName,
			PostalCode:  destination.DestinationAddress.PostalCode,
		},
		UrlImages:  ToUrlImages(destination),
		Categories: ToCategory(destination),
		Facilities: ToFacilities(destination),
	}
}

func ToDestinationFacilities(destinationId uuid.UUID, facilityIds []uuid.UUID) *[]entities.DestinationFacility {
	var destinationFacilities []entities.DestinationFacility
	for _, facilityId := range facilityIds {
		destinationFacilities = append(destinationFacilities, entities.DestinationFacility{
			Id:            uuid.New(),
			DestinationId: destinationId,
			FacilityId:    facilityId,
		})
	}
	return &destinationFacilities
}

func ToDestinationMedia(destinationId uuid.UUID, mediaType string, mediaUrl string, title string) *entities.DestinationMedia {
	return &entities.DestinationMedia{
		Id:            uuid.New(),
		DestinationId: destinationId,
		Type:          mediaType,
		Url:           mediaUrl,
		Title:         title,
	}
}

func ToDestinationAddress(destinationId uuid.UUID, request CreateDestinationAddressRequest) *entities.DestinationAddress {
	return &entities.DestinationAddress{
		Id:            uuid.New(),
		ProvinceId:    request.ProvinceId,
		CityId:        request.CityId,
		SubdistrictId: request.SubdistrictId,
		StreetName:    request.StreetName,
		PostalCode:    request.PostalCode,
		DestinationId: destinationId,
	}
}

type CreateDestinationImageRequest struct {
	File  multipart.File `json:"file" form:"file" validate:"required"`
	Title string         `json:"judul" form:"judul" validate:"required"`
}

type CreateDestinationAddressRequest struct {
	ProvinceId    string `json:"id_provinsi" form:"id_provinsi" validate:"required"`
	CityId        string `json:"id_kota" form:"id_kota" validate:"required"`
	SubdistrictId string `json:"id_kecamatan" form:"id_kecamatan" validate:"required"`
	StreetName    string `json:"jalan" form:"jalan" validate:"required"`
	PostalCode    string `json:"kode_pos" form:"kode_pos" validate:"required"`
}

type CreateDestinationRequest struct {
	Name               string                          `json:"nama_destinasi" form:"nama_destinasi" validate:"required"`
	Description        string                          `json:"deskripsi" form:"deskripsi" validate:"required"`
	OpenTime           string                          `json:"jam_buka" form:"jam_buka" validate:"required"`
	CloseTime          string                          `json:"jam_tutup" form:"jam_tutup" validate:"required"`
	EntryPrice         float64                         `json:"harga_masuk" form:"harga_masuk" validate:"required"`
	CategoryId         uuid.UUID                       `json:"id_kategori" form:"id_kategori" validate:"required"`
	Latitude           float64                         `json:"latitude" form:"latitude" validate:"required"`
	Longitude          float64                         `json:"longitude" form:"longitude" validate:"required"`
	FacilityIds        []uuid.UUID                     `json:"fasilitas" form:"fasilitas" validate:"required"`
	DestinationAddress CreateDestinationAddressRequest `json:"alamat_destinasi" form:"alamat_destinasi" validate:"required"`
}

type UpdateDestinationImageRequest struct {
	Id    uuid.UUID      `json:"id" form:"id" validate:"required"`
	File  multipart.File `json:"file" form:"file" validate:"required"`
	Title string         `json:"judul" form:"judul" validate:"required"`
}

type UpdateDestinationRequest struct {
	Name               string                          `json:"nama_destinasi" form:"nama_destinasi" validate:"required"`
	Description        string                          `json:"deskripsi" form:"deskripsi" validate:"required"`
	OpenTime           string                          `json:"jam_buka" form:"jam_buka" validate:"required"`
	CloseTime          string                          `json:"jam_tutup" form:"jam_tutup" validate:"required"`
	EntryPrice         float64                         `json:"harga_masuk" form:"harga_masuk" validate:"required"`
	CategoryId         uuid.UUID                       `json:"id_kategori" form:"id_kategori" validate:"required"`
	Latitude           float64                         `json:"latitude" form:"latitude" validate:"required"`
	Longitude          float64                         `json:"longitude" form:"longitude" validate:"required"`
	FacilityIds        []uuid.UUID                     `json:"fasilitas" form:"fasilitas" validate:"required"`
	DestinationAddress CreateDestinationAddressRequest `json:"alamat_destinasi" form:"alamat_destinasi" validate:"required"`
}

type CreateDestinationResponse struct {
	Id uuid.UUID `json:"id_destinasi"`
}
