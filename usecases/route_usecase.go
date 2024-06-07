package usecases

import (
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/repositories"

	"github.com/google/uuid"
)

type RouteInterface interface {
	FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error)
	FindById(id uuid.UUID) (*entities.Route, error)
	DeleteRoute(id uuid.UUID) error
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
		return nil, &errorHandlers.ConflictError{Message: "Rute tidak ditemukan"}
	}
	return route, nil
}

func (uc *routeUsecase) DeleteRoute(id uuid.UUID) error {
	route, err := uc.routeRepo.FindById(id)
	if err != nil {
		return &errorHandlers.ConflictError{Message: "Rute tidak ditemukan"}
	}
	if err = uc.routeRepo.Delete(route); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menghapus rute"}
	}

	if err = uc.routeDetailRepo.DeleteMany(&route.RouteDetail); err != nil {
		return &errorHandlers.InternalServerError{Message: "Gagal menghapus detail rute"}
	}
	return nil
}
