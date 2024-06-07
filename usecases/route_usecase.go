package usecases

import (
	"capstone/entities"
	"capstone/errorHandlers"
	"github.com/google/uuid"
)

type RouteInterface interface {
	FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error)
	FindById(id uuid.UUID) (*entities.Route, error)
}

type routeUsecase struct {
	repository RouteInterface
}

func NewRouteUsecase(repository RouteInterface) *routeUsecase {
	return &routeUsecase{repository: repository}
}

func (uc *routeUsecase) FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error) {
	routes, total, err := uc.repository.FindAll(page, limit, searchQuery)
	if err != nil {
		return nil, nil, err
	}
	return routes, total, nil
}

func (uc *routeUsecase) FindById(id uuid.UUID) (*entities.Route, error) {
	route, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.ConflictError{Message: "Rute tidak ditemukan"}
	}
	return route, nil
}
