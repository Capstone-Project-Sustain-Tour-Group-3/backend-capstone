package dto

import (
	"github.com/google/uuid"
)

type ProvinceResponse struct {
	Id   string `json:"id"`
	Name string `json:"nama"`
	Url  string `json:"url"`
}

type CategoryResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"nama"`
	Url  string    `json:"url"`
}

type PersonalizationRequest struct {
	CategoryIds []uuid.UUID `json:"id_kategori" validate:"gte=1,lte=3,dive,required"`
	ProvinceIds []string    `json:"id_provinsi" validate:"gte=1,lte=3,dive,required,len=2"`
}
