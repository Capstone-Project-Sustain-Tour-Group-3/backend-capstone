package dto

import (
	"capstone/entities"
	"github.com/google/uuid"
)

type findAllResponse struct {
	Id            uuid.UUID     `json:"id"`
	NamaPengguna  string        `json:"nama_pengguna"`
	Kota          string        `json:"kota"`
	EstimasiBiaya float64       `json:"estimasi_biaya"`
	RouteDetail   []RouteDetail `json:"route_detail"`
}

type Destinations struct {
	Id            uuid.UUID `json:"id"`
	NamaDestinasi string    `json:"nama_destinasi"`
}

type RouteDetail struct {
	Id        uuid.UUID    `json:"id"`
	Destinasi Destinations `json:"destinasi"`
}

func ToFindAllRouteResponse(routes *[]entities.Route) *[]findAllResponse {
	var response []findAllResponse
	for _, route := range *routes {
		var routeDetails []RouteDetail
		for _, routeDetail := range route.RouteDetail {
			var destinations Destinations
			destinations.Id = routeDetail.DestinationId
			destinations.NamaDestinasi = routeDetail.Destination.Name
			routeDetails = append(routeDetails, RouteDetail{
				Id:        routeDetail.Id,
				Destinasi: destinations,
			})
		}
		response = append(response, findAllResponse{
			Id:            route.Id,
			NamaPengguna:  route.User.Username,
			Kota:          route.City.Name,
			EstimasiBiaya: route.Price,
			RouteDetail:   routeDetails,
		})
	}
	return &response
}

func ToFindByIdRouteResponse(route *entities.Route) *findAllResponse {
	var routeDetails []RouteDetail
	for _, routeDetail := range route.RouteDetail {
		var destinations Destinations
		destinations.Id = routeDetail.DestinationId
		destinations.NamaDestinasi = routeDetail.Destination.Name
		routeDetails = append(routeDetails, RouteDetail{
			Id:        routeDetail.Id,
			Destinasi: destinations,
		})
	}
	response := findAllResponse{
		Id:            route.Id,
		NamaPengguna:  route.User.Username,
		Kota:          route.City.Name,
		EstimasiBiaya: route.Price,
		RouteDetail:   routeDetails,
	}
	return &response
}
