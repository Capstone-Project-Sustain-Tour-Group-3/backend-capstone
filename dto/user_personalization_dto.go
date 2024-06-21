package dto

import (
	"capstone/entities"

	"github.com/google/uuid"
)

type PersonalizationRequest struct {
	CategoryIds []uuid.UUID `json:"id_kategori" validate:"gte=1,lte=3,dive,required"`
	ProvinceIds []string    `json:"id_provinsi" validate:"gte=1,lte=3,dive,required,len=2"`
}

func ToPersonalizationProvinceIds(personalizations *[]entities.PersonalizationProvince) []string {
	var provinceIds []string
	for _, personalization := range *personalizations {
		provinceIds = append(provinceIds, personalization.ProvinceId)
	}
	return provinceIds
}

func ToCategoryIds(personalizations *[]entities.PersonalizationCategory) []uuid.UUID {
	var categoryIds []uuid.UUID
	for _, personalization := range *personalizations {
		categoryIds = append(categoryIds, personalization.CategoryId)
	}
	return categoryIds
}
