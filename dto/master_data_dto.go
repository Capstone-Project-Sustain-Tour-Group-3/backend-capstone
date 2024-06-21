package dto

import "github.com/google/uuid"

type CategoryResponse struct {
	Id   uuid.UUID `json:"kategori_id"`
	Name string    `json:"nama"`
	Url  string    `json:"url"`
}

type ProvinceResponse struct {
	Id   string `json:"provinsi_id"`
	Name string `json:"nama"`
	Url  string `json:"url"`
}

type CityResponse struct {
	Id         string `json:"kota_id"`
	ProvinceId string `json:"provinsi_id"`
	Name       string `json:"nama"`
}

type SubdistrictResponse struct {
	Id     string `json:"kecamatan_id"`
	CityId string `json:"kota_id"`
	Name   string `json:"nama"`
}

type FacilityResponse struct {
	Id   uuid.UUID `json:"fasilitas_id"`
	Name string    `json:"nama"`
	Url  string    `json:"url"`
}
