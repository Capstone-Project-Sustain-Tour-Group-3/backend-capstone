package usecases

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type RouteInterface interface {
	FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error)
	FindById(id uuid.UUID) (*entities.Route, error)
	DeleteRoute(id uuid.UUID) error
	SaveRoute(request *dto.SaveRouteRequest) error
}

type routeUsecase struct {
	routeRepo       repositories.RouteRepository
	routeDetailRepo repositories.RouteDetailRepository
}

func NewRouteUsecase(routeRepo repositories.RouteRepository, routeDetail repositories.RouteDetailRepository) *routeUsecase {
	return &routeUsecase{routeRepo: routeRepo, routeDetailRepo: routeDetail}
}

func (uc *routeUsecase) FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error) {
	routes, total, err := uc.routeRepo.FindAll(page, limit, searchQuery)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data rute"}
	}
	return routes, total, nil
}

func (uc *routeUsecase) FindById(id uuid.UUID) (*entities.Route, error) {
	route, err := uc.routeRepo.FindById(id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"}
	}
	return route, nil
}

func (uc *routeUsecase) DeleteRoute(id uuid.UUID) error {
	route, err := uc.routeRepo.FindById(id)
	if err != nil {
		return &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"}
	}
	if err = uc.routeRepo.Delete(route); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menghapus rute"}
	}

	if err = uc.routeDetailRepo.DeleteMany(&route.RouteDetail); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menghapus detail rute"}
	}
	return nil
}

func (uc *routeUsecase) SaveRoute(request *dto.SaveRouteRequest) error {
	route := &entities.Route{
		Id:             uuid.New(),
		UserId:         request.UserId,
		CityId:         request.CityId,
		Name:           request.Name,
		StartLocation:  request.StartLocation,
		StartLongitude: request.StartLongitude,
		StartLatitude:  request.StartLatitude,
		Price:          request.Price,
	}

	if request.RouteDetails == nil {
		return &errorHandlers.BadRequestError{Message: "data detail rute tidak valid"}
	}

	if err := uc.routeRepo.Create(route); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menyimpan rute"}
	}

	for _, routeDetail := range request.RouteDetails {
		routeDetail := &entities.RouteDetail{
			Id:            uuid.New(),
			RouteId:       route.Id,
			DestinationId: routeDetail.DestinationId,
			Longitude:     routeDetail.Longitude,
			Latitude:      routeDetail.Latitude,
			Duration:      routeDetail.Duration,
			Order:         routeDetail.Order,
			VisitStart:    []byte(routeDetail.VisitStart),
			VisitEnd:      []byte(routeDetail.VisitEnd),
		}

		if err := uc.routeDetailRepo.Create(routeDetail); err != nil {
			return &errorHandlers.InternalServerError{Message: "Gagal menyimpan detail rute"}
		}
	}

	return nil
}
