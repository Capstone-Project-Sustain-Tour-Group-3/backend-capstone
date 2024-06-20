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

type RouteSummaryRequest struct {
	CityId         string        `json:"id_kota" validate:"required,len=4"`
	StartLocation  StartLocation `json:"lokasi_awal" validate:"required"`
	DestinationIds []uuid.UUID   `json:"id_destinasi_tujuan" validate:"gte=1,lte=3,dive,required"`
}

type RouteSummaryResponse struct {
	StartLocation  StartLocation `json:"lokasi_awal"`
	RouteDetails   []RouteDetail `json:"detail_rute"`
	EstimationCost Currency      `json:"estimasi_biaya"`
}

type StartLocation struct {
	Name string  `json:"nama" validate:"required"`
	Lat  float64 `json:"latitude" validate:"latitude"`
	Long float64 `json:"longitude" validate:"longitude"`
}

type RouteDetail struct {
	Destination RouteDestination `json:"destinasi"`
	Duration    Duration         `json:"durasi"`
	Order       int              `json:"urutan"`
	VisitStart  string           `json:"waktu_kunjungan"`
	VisitEnd    string           `json:"waktu_selesai"`
}

type RouteDestination struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"nama"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ImageURL  *string   `json:"url_gambar"`
	EntryCost Currency  `json:"biaya_masuk"`
}

type Currency struct {
	Raw    float64 `json:"raw"`
	Format string  `json:"format"`
}

type Duration struct {
	Raw        int    `json:"raw"`
	Simplified string `json:"simple"`
	Full       string `json:"full"`
}

type GetAllRouteByCurrUserResponse struct {
	Id       uuid.UUID `json:"id"`
	NamaRute string    `json:"nama_rute"`
	Biaya    float64   `json:"biaya"`
}

type GetDetailRouteByCurrUserRequest struct {
	NamaRute      string                          `json:"nama_rute"`
	Destinasi     []DetailRouteByCurrUserResponse `json:"destinasi"`
	EstimasiBiaya float64                         `json:"estimasi_biaya"`
}

type DetailRouteByCurrUserResponse struct {
	Id                    uuid.UUID `json:"id"`
	NamaDestinasi         string    `json:"nama_destinasi"`
	WaktuKunjungan        string    `json:"waktu_kunjungan"`
	WaktuSelesaiKunjungan string    `json:"waktu_selesai_kunjungan"`
	Biaya                 float64   `json:"biaya"`
}

func ToDetailRouteByCurrUserResponse(route *entities.Route) *GetDetailRouteByCurrUserRequest {
	var destinasi []DetailRouteByCurrUserResponse
	for _, routeDetail := range route.RouteDetail {
		destinasi = append(destinasi, DetailRouteByCurrUserResponse{
			Id:                    routeDetail.DestinationId,
			NamaDestinasi:         routeDetail.Destination.Name,
			WaktuKunjungan:        string(routeDetail.VisitStart)[:5],
			WaktuSelesaiKunjungan: string(routeDetail.VisitEnd)[:5],
			Biaya:                 routeDetail.Destination.EntryPrice,
		})
	}
	response := GetDetailRouteByCurrUserRequest{
		NamaRute:      route.Name,
		Destinasi:     destinasi,
		EstimasiBiaya: route.Price,
	}
	return &response
}

func ToGetAllRouteByCurrUserResponse(routes *[]entities.Route) *[]GetAllRouteByCurrUserResponse {
	var responses []GetAllRouteByCurrUserResponse
	for _, route := range *routes {
		response := GetAllRouteByCurrUserResponse{
			Id:       route.Id,
			NamaRute: route.Name,
			Biaya:    route.Price,
		}
		responses = append(responses, response)
	}
	return &responses
}

func ToSummaryRouteResponse(startLocation StartLocation, routeDetails []RouteDetail, estimationCost Currency) *RouteSummaryResponse {
	return &RouteSummaryResponse{
		StartLocation:  startLocation,
		RouteDetails:   routeDetails,
		EstimationCost: estimationCost,
	}
}

func ToStartLocation(name string, lat float64, long float64) StartLocation {
	return StartLocation{
		Name: name,
		Lat:  lat,
		Long: long,
	}
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

type DetailRouteRequest struct {
	DestinationId uuid.UUID `json:"destination_id" validate:"required"`
	Longitude     float64   `json:"longitude" validate:"required"`
	Latitude      float64   `json:"latitude" validate:"required"`
	Duration      int       `json:"duration" validate:"required"`
	Order         int       `json:"order" validate:"required"`
	VisitStart    string    `json:"visit_start" validate:"required"`
	VisitEnd      string    `json:"visit_end" validate:"required"`
}

type SaveRouteRequest struct {
	UserId         uuid.UUID            `json:"user_id" validate:"required"`
	CityId         string               `json:"city_id" validate:"required,len=4"`
	Name           string               `json:"name" validate:"required"`
	StartLocation  string               `json:"start_location" validate:"required"`
	StartLongitude float64              `json:"start_longitude" validate:"required"`
	StartLatitude  float64              `json:"start_latitude" validate:"required"`
	Price          float64              `json:"price" validate:"required"`
	RouteDetails   []DetailRouteRequest `json:"route_details" validate:"required,dive"`
}
