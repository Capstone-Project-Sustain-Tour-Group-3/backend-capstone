package dto

import (
	"github.com/google/uuid"
)

type ProvinceResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"nama"`
	Url  string    `json:"url"`
}

type CategoryResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"nama"`
	Url  string    `json:"url"`
}

type PersonalizationRequest struct {
	CategoryId  uuid.UUID `json:"id_kategori" validate:"required"`
	ProvinceIds []string  `json:"id_provinsi" validate:"gte=1,lte=3,dive,required,len=2"`
}
