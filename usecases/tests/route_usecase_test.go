package tests

import (
	"testing"

	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/externals"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindAllRoute(t *testing.T) {
	page := 1
	limit := 10
	searchQuery := "test"
	totalExpected := int64(2)

	routeLists := []entities.Route{
		{Id: uuid.New(), Name: "Route 1"},
		{Id: uuid.New(), Name: "Route 2"},
	}
	testCases := []struct {
		name           string
		page           int
		limit          int
		searchQuery    string
		mockSetup      func(routeRepo *repositories.MockRouteRepository)
		expectedRoutes *[]entities.Route
		expectedTotal  *int64
		expectedError  error
	}{
		{
			name:        "Success find all routes",
			page:        page,
			limit:       limit,
			searchQuery: searchQuery,
			mockSetup: func(routeRepo *repositories.MockRouteRepository) {
				routeRepo.On("FindAll", page, limit, searchQuery).Return(&routeLists, &totalExpected, nil)
			},
			expectedRoutes: &routeLists,
			expectedTotal:  &totalExpected,
			expectedError:  nil,
		},
		{
			name:        "Failed find all routes",
			page:        page,
			limit:       limit,
			searchQuery: searchQuery,
			mockSetup: func(routeRepo *repositories.MockRouteRepository) {
				routeRepo.On("FindAll", page, limit, searchQuery).Return(nil, nil, assert.AnError)
			},
			expectedRoutes: nil,
			expectedTotal:  nil,
			expectedError:  &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data rute"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cityRepo := new(repositories.MockCityRepository)
			destinationRepo := new(repositories.MockDestinationRepository)
			routeRepo := new(repositories.MockRouteRepository)
			routeDet := new(repositories.MockRouteDetailRepository)
			openAIClient := new(externals.MockOpenAIClient)
			uc := usecases.NewRouteUsecase(cityRepo, destinationRepo, routeRepo, routeDet, openAIClient)

			tc.mockSetup(routeRepo)

			routes, total, err := uc.FindAll(tc.page, tc.limit, tc.searchQuery)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedRoutes, routes)
				require.Equal(t, tc.expectedTotal, total)
			}

			routeRepo.AssertExpectations(t)
		})
	}
}

func TestFindRouteById(t *testing.T) {
	id := uuid.New()
	routeExpected := entities.Route{
		Id:   id,
		Name: "Route 1",
	}
	testCases := []struct {
		name          string
		id            uuid.UUID
		mockSetup     func(routeRepo *repositories.MockRouteRepository)
		expectedRoute *entities.Route
		expectedError error
	}{
		{
			name: "Success find route by id",
			id:   id,
			mockSetup: func(routeRepo *repositories.MockRouteRepository) {
				routeRepo.On("FindById", id).Return(&routeExpected, nil)
			},
			expectedRoute: &routeExpected,
			expectedError: nil,
		},
		{
			name: "Failed find route by id",
			id:   id,
			mockSetup: func(routeRepo *repositories.MockRouteRepository) {
				routeRepo.On("FindById", id).Return(nil, assert.AnError)
			},
			expectedRoute: nil,
			expectedError: &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cityRepo := new(repositories.MockCityRepository)
			destinationRepo := new(repositories.MockDestinationRepository)
			routeRepo := new(repositories.MockRouteRepository)
			routeDet := new(repositories.MockRouteDetailRepository)
			openAIClient := new(externals.MockOpenAIClient)
			uc := usecases.NewRouteUsecase(cityRepo, destinationRepo, routeRepo, routeDet, openAIClient)

			tc.mockSetup(routeRepo)

			route, err := uc.FindById(tc.id)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedRoute, route)
			}

			routeRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteRoute(t *testing.T) {
	id := uuid.New()
	routeExpected := entities.Route{
		Id:   id,
		Name: "Route 1",
	}
	testCases := []struct {
		name          string
		id            uuid.UUID
		mockSetup     func(routeRepo *repositories.MockRouteRepository, routeDet *repositories.MockRouteDetailRepository)
		expectedError error
	}{
		{
			name: "Success delete route",
			id:   id,
			mockSetup: func(routeRepo *repositories.MockRouteRepository, routeDet *repositories.MockRouteDetailRepository) {
				routeRepo.On("FindById", id).Return(&routeExpected, nil)
				routeRepo.On("Delete", &routeExpected).Return(nil)
				routeDet.On("DeleteMany", &routeExpected.RouteDetail).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Failed delete route",
			id:   id,
			mockSetup: func(routeRepo *repositories.MockRouteRepository, routeDet *repositories.MockRouteDetailRepository) {
				routeRepo.On("FindById", id).Return(nil, assert.AnError)
			},
			expectedError: &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"},
		},
		{
			name: "Failed delete route",
			id:   id,
			mockSetup: func(routeRepo *repositories.MockRouteRepository, routeDet *repositories.MockRouteDetailRepository) {
				routeRepo.On("FindById", id).Return(&routeExpected, nil)
				routeRepo.On("Delete", &routeExpected).Return(assert.AnError)
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menghapus rute"},
		},
		{
			name: "Failed delete many route detail",
			id:   id,
			mockSetup: func(routeRepo *repositories.MockRouteRepository, routeDet *repositories.MockRouteDetailRepository) {
				routeRepo.On("FindById", id).Return(&routeExpected, nil)
				routeRepo.On("Delete", &routeExpected).Return(nil)
				routeDet.On("DeleteMany", &routeExpected.RouteDetail).Return(assert.AnError)
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menghapus detail rute"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cityRepo := new(repositories.MockCityRepository)
			destinationRepo := new(repositories.MockDestinationRepository)
			routeRepo := new(repositories.MockRouteRepository)
			routeDet := new(repositories.MockRouteDetailRepository)
			openAIClient := new(externals.MockOpenAIClient)
			uc := usecases.NewRouteUsecase(cityRepo, destinationRepo, routeRepo, routeDet, openAIClient)

			tc.mockSetup(routeRepo, routeDet)

			err := uc.DeleteRoute(tc.id)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			routeRepo.AssertExpectations(t)
		})
	}
}
