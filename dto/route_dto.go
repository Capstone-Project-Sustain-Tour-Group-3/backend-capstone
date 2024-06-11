package dto

import (
	"capstone/entities"

	"github.com/google/uuid"
)

type findAllResponse struct {
	Id            uuid.UUID   `json:"id"`
	Username      string      `json:"username"`
	Kota          string      `json:"kota"`
	NamaRute      string      `json:"nama_rute"`
	Destinasi     []Destinasi `json:"destinasi"`
	EstimasiBiaya float64     `json:"estimasi_biaya"`
}

type Destinasi struct {
	Id            uuid.UUID `json:"id"`
	NamaDestinasi string    `json:"nama_destinasi"`
}

func ToFindAllRouteResponse(routes *[]entities.Route) *[]findAllResponse {
	var responses []findAllResponse
	for _, route := range *routes {
		var destinasi []Destinasi
		for _, routeDetail := range route.RouteDetail {
			destinasi = append(destinasi, Destinasi{
				Id:            routeDetail.DestinationId,
				NamaDestinasi: routeDetail.Destination.Name,
			})
		}
		response := findAllResponse{
			Id:            route.Id,
			Username:      route.User.Username,
			Kota:          route.City.Name,
			NamaRute:      route.Name,
			Destinasi:     destinasi,
			EstimasiBiaya: route.Price,
		}
		responses = append(responses, response)
	}
	return &responses
}

func ToFindByIdRouteResponse(route *entities.Route) *findAllResponse {
	var destinasi []Destinasi
	for _, routeDetail := range route.RouteDetail {
		destinasi = append(destinasi, Destinasi{
			Id:            routeDetail.DestinationId,
			NamaDestinasi: routeDetail.Destination.Name,
		})
	}
	response := findAllResponse{
		Id:            route.Id,
		Username:      route.User.Username,
		Kota:          route.City.Name,
		NamaRute:      route.Name,
		Destinasi:     destinasi,
		EstimasiBiaya: route.Price,
	}
	return &response
}
