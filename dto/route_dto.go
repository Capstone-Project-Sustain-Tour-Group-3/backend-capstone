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

type DetailRouteRequest struct {
	DestinationId uuid.UUID `json:"destination_id" validate:"required"`
	Longitude     float64   `json:"longitude" validate:"required"`
	Latitude      float64   `json:"latitude" validate:"required"`
	Duration      int       `json:"duration" validate:"required"`
	Order         int       `json:"order" validate:"required"`
	VisitStart    string    `json:"visit_start" validate:""`
	VisitEnd      string    `json:"visit_end" validate:""`
}

type SaveRouteRequest struct {
	UserId         uuid.UUID            `json:"user_id" validate:"required"`
	CityId         string               `json:"city_id" validate:"required"`
	Name           string               `json:"name" validate:"required"`
	StartLocation  string               `json:"start_location" validate:"required"`
	StartLongitude float64              `json:"start_longitude" validate:"required"`
	StartLatitude  float64              `json:"start_latitude" validate:"required"`
	Price          float64              `json:"price" validate:"required"`
	RouteDetails   []DetailRouteRequest `json:"route_details" validate:"required"`
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
