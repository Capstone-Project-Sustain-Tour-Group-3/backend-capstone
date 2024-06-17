package usecases

import (
	"errors"
	"fmt"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/externals/openai"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RouteInterface interface {
	FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error)
	FindById(id uuid.UUID) (*entities.Route, error)
	DeleteRoute(id uuid.UUID) error
	SummarizeRoute(req *dto.RouteSummaryRequest) (*dto.RouteSummaryResponse, error)
}

type routeUsecase struct {
	cityRepo        repositories.ICityRepository
	destinationRepo repositories.IDestinationRepository
	routeRepo       repositories.RouteRepository
	routeDetailRepo repositories.RouteDetailRepository

	openAIClient openai.OpenAIClient
}

func NewRouteUsecase(
	cityRepo repositories.ICityRepository,
	destinationRepo repositories.IDestinationRepository,
	routeRepo repositories.RouteRepository,
	routeDetail repositories.RouteDetailRepository,
	openAIClient openai.OpenAIClient,
) *routeUsecase {
	return &routeUsecase{
		cityRepo:        cityRepo,
		destinationRepo: destinationRepo,
		routeRepo:       routeRepo,
		routeDetailRepo: routeDetail,
		openAIClient:    openAIClient,
	}
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

func (uc *routeUsecase) SummarizeRoute(req *dto.RouteSummaryRequest) (*dto.RouteSummaryResponse, error) {
	var (
		destinations []entities.Destination
		err          error
	)

	for idx, id := range req.DestinationIds {
		var destination *entities.Destination
		destination, err = uc.destinationRepo.FindByIdInCityId(id, req.CityId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, &errorHandlers.NotFoundError{Message: fmt.Sprintf("Destinasi ke-%d tidak ditemukan", idx+1)}
			}

			return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan detail destinasi"}
		}
		destinations = append(destinations, *destination)
	}

	prompt := helpers.ToRoutePrompt(&destinations, req.StartLocation.Lat, req.StartLocation.Long)
	recommendationRoute, err := uc.openAIClient.GenerateAnswer(
		prompt,
		helpers.GetRouteSystemInstruction(),
	)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi rute"}
	}

	routeDetails, estimationCost := helpers.ExtractRouteInformation(
		recommendationRoute,
		destinations,
	)

	res := dto.ToSummaryRouteResponse(
		dto.ToStartLocation(req.StartLocation.Name, req.StartLocation.Lat, req.StartLocation.Long),
		routeDetails,
		estimationCost,
	)

	return res, nil
}
