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
	SaveRoute(req *dto.SaveRouteRequest) error
	FindAllByCurrentUser(user_id uuid.UUID) (*[]entities.Route, error)
	FindByIdCurrentUser(user_id, id uuid.UUID) (*entities.Route, error)
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

func (uc *routeUsecase) FindAllByCurrentUser(user_id uuid.UUID) (*[]entities.Route, error) {
	routes, err := uc.routeRepo.FindAllByCurrentUser(user_id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "Gagal mendapatkan data rute"}
	}
	return routes, nil
}

func (uc *routeUsecase) FindByIdCurrentUser(user_id, id uuid.UUID) (*entities.Route, error) {
	route, err := uc.routeRepo.FindByIdCurrentUser(user_id, id)
	if err != nil {
		return nil, &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"}
	}
	return route, nil
}
