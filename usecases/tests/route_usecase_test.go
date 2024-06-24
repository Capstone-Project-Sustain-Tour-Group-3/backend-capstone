package tests

import (
	"errors"
	"testing"

	"capstone/dto"
	"capstone/helpers"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/externals"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type summarizeRouteTestCase struct {
	name          string
	mockSetup     func(destinationRepo *repositories.MockDestinationRepository, openAIClient *externals.MockOpenAIClient)
	expectedResp  *dto.RouteSummaryResponse
	expectedError error
}

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

func TestFindAllByCurrentUser(t *testing.T) {
	mockRouteRepo := new(repositories.MockRouteRepository)

	uc := usecases.NewRouteUsecase(nil, nil, mockRouteRepo, nil, nil)

	testCases := []struct {
		name               string
		userId             uuid.UUID
		mockRouteRepoSetup func()
		expectedResponse   *[]entities.Route
		expectedError      error
	}{
		{
			name:   "Success get routes by current user",
			userId: uuid.New(),
			mockRouteRepoSetup: func() {
				mockRoutes := []entities.Route{
					{Id: uuid.New(), UserId: uuid.New(), Name: "Route 1"},
					{Id: uuid.New(), UserId: uuid.New(), Name: "Route 2"},
				}

				mockRouteRepo.On("FindAllByCurrentUser", mock.Anything).Return(&mockRoutes, nil).Once()
			},
			expectedResponse: &[]entities.Route{
				{Id: uuid.New(), UserId: uuid.New(), Name: "Route 1"},
				{Id: uuid.New(), UserId: uuid.New(), Name: "Route 2"},
			},
			expectedError: nil,
		},
		{
			name:   "No routes found",
			userId: uuid.New(),
			mockRouteRepoSetup: func() {
				mockRouteRepo.On("FindAllByCurrentUser", mock.Anything).Return(nil, &errorHandlers.NotFoundError{Message: "Gagal mendapatkan data rute"}).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.NotFoundError{Message: "Gagal mendapatkan data rute"},
		},
		{
			name:   "Internal server error",
			userId: uuid.New(),
			mockRouteRepoSetup: func() {
				mockRouteRepo.On("FindAllByCurrentUser", mock.Anything).Return(nil, errors.New("internal error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.NotFoundError{Message: "Gagal mendapatkan data rute"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockRouteRepoSetup()

			response, err := uc.FindAllByCurrentUser(tc.userId)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, response)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			}

			mockRouteRepo.AssertExpectations(t)
		})
	}
}

func TestFindByIdCurrentUser(t *testing.T) {
	mockRouteRepo := new(repositories.MockRouteRepository)

	uc := usecases.NewRouteUsecase(nil, nil, mockRouteRepo, nil, nil)

	testCases := []struct {
		name               string
		userId, routeId    uuid.UUID
		mockRouteRepoSetup func()
		expectedResponse   *entities.Route
		expectedError      error
	}{
		{
			name:    "Success get route by current user",
			userId:  uuid.New(),
			routeId: uuid.New(),
			mockRouteRepoSetup: func() {
				mockRoute := &entities.Route{Id: uuid.New(), UserId: uuid.New(), Name: "Route 1"}

				mockRouteRepo.On("FindByIdCurrentUser", mock.Anything, mock.Anything).Return(mockRoute, nil).Once()
			},
			expectedResponse: &entities.Route{Id: uuid.New(), UserId: uuid.New(), Name: "Route 1"},
			expectedError:    nil,
		},
		{
			name:    "Route not found",
			userId:  uuid.New(),
			routeId: uuid.New(),
			mockRouteRepoSetup: func() {
				mockRouteRepo.On("FindByIdCurrentUser", mock.Anything, mock.Anything).Return(nil, &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"}).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"},
		},
		{
			name:    "Internal server error",
			userId:  uuid.New(),
			routeId: uuid.New(),
			mockRouteRepoSetup: func() {
				mockRouteRepo.On("FindByIdCurrentUser", mock.Anything, mock.Anything).Return(nil, &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"}).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.NotFoundError{Message: "Rute tidak ditemukan"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockRouteRepoSetup()

			response, err := uc.FindByIdCurrentUser(tc.userId, tc.routeId)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, response)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			}

			mockRouteRepo.AssertExpectations(t)
		})
	}
}

func TestSaveRoute(t *testing.T) {
	mockRouteRepo := new(repositories.MockRouteRepository)
	mockRouteDetailRepo := new(repositories.MockRouteDetailRepository)

	uc := usecases.NewRouteUsecase(nil, nil, mockRouteRepo, mockRouteDetailRepo, nil)

	testCases := []struct {
		name          string
		request       *dto.SaveRouteRequest
		mockRepoSetup func()
		expectedError error
	}{
		{
			name: "Success save route and route details",
			request: &dto.SaveRouteRequest{
				UserId:         uuid.New(),
				CityId:         "ABC",
				Name:           "Route 1",
				StartLocation:  "Start Location",
				StartLongitude: 100.0,
				StartLatitude:  10.0,
				Price:          50000,
				RouteDetails: []dto.DetailRouteRequest{
					{
						DestinationId: uuid.New(),
						Longitude:     100.1,
						Latitude:      10.1,
						Duration:      30,
						Order:         1,
						VisitStart:    "2024-06-22T08:00:00Z",
						VisitEnd:      "2024-06-22T09:00:00Z",
					},
				},
			},
			mockRepoSetup: func() {
				mockRouteRepo.On("Create", mock.Anything).Return(nil).Once()
				mockRouteDetailRepo.On("Create", mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Invalid route details",
			request: &dto.SaveRouteRequest{
				UserId:         uuid.New(),
				CityId:         "ABC",
				Name:           "Route 1",
				StartLocation:  "Start Location",
				StartLongitude: 100.0,
				StartLatitude:  10.0,
				Price:          50000,
				RouteDetails:   nil,
			},
			mockRepoSetup: func() {},
			expectedError: &errorHandlers.BadRequestError{Message: "data detail rute tidak valid"},
		},
		{
			name: "Failed to save route",
			request: &dto.SaveRouteRequest{
				UserId:         uuid.New(),
				CityId:         "ABC",
				Name:           "Route 1",
				StartLocation:  "Start Location",
				StartLongitude: 100.0,
				StartLatitude:  10.0,
				Price:          50000,
				RouteDetails: []dto.DetailRouteRequest{
					{
						DestinationId: uuid.New(),
						Longitude:     100.1,
						Latitude:      10.1,
						Duration:      30,
						Order:         1,
						VisitStart:    "2024-06-22T08:00:00Z",
						VisitEnd:      "2024-06-22T09:00:00Z",
					},
				},
			},
			mockRepoSetup: func() {
				mockRouteRepo.On("Create", mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal menyimpan detail rute"}).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menyimpan rute"},
		},
		{
			name: "Failed to save route detail",
			request: &dto.SaveRouteRequest{
				UserId:         uuid.New(),
				CityId:         "ABC",
				Name:           "Route 1",
				StartLocation:  "Start Location",
				StartLongitude: 100.0,
				StartLatitude:  10.0,
				Price:          50000,
				RouteDetails: []dto.DetailRouteRequest{
					{
						DestinationId: uuid.New(),
						Longitude:     100.1,
						Latitude:      10.1,
						Duration:      30,
						Order:         1,
						VisitStart:    "2024-06-22T08:00:00Z",
						VisitEnd:      "2024-06-22T09:00:00Z",
					},
				},
			},
			mockRepoSetup: func() {
				mockRouteRepo.On("Create", mock.Anything).Return(nil).Once()
				mockRouteDetailRepo.On("Create", mock.Anything).Return(&errorHandlers.InternalServerError{Message: "Gagal menyimpan detail rute"}).Once()
			},
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal menyimpan detail rute"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockRepoSetup()

			err := uc.SaveRoute(tc.request)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockRouteRepo.AssertExpectations(t)
			mockRouteDetailRepo.AssertExpectations(t)
		})
	}
}

func TestSummarizeRoute(t *testing.T) {
	destinationId := uuid.New()
	req := &dto.RouteSummaryRequest{
		CityId: "1001",
		StartLocation: dto.StartLocation{
			Lat:  10.0,
			Long: 100.0,
			Name: "Start Location",
		},
		DestinationIds: []uuid.UUID{destinationId},
	}
	destinations := []entities.Destination{
		{
			Id:         destinationId,
			Name:       "Destination 1",
			Longitude:  100.0,
			Latitude:   10.0,
			OpenTime:   "08:00",
			CloseTime:  "17:00",
			EntryPrice: 50000,
		},
	}
	destination := destinations[0]
	recommendationRoute := "1,100.0,10.0,Destination 1,08:00,17:00\nTotal Biaya:50000"
	routeDetails, estimationCost := helpers.ExtractRouteInformation(
		recommendationRoute,
		destinations,
	)
	expectedResp := dto.ToSummaryRouteResponse(
		dto.ToStartLocation(req.StartLocation.Name, req.StartLocation.Lat, req.StartLocation.Long),
		routeDetails,
		estimationCost,
	)

	testCases := []summarizeRouteTestCase{
		{
			name: "Success summarize route",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository, openAIClient *externals.MockOpenAIClient) {
				destinationRepo.On("FindByIdInCityId", mock.Anything, mock.Anything).Return(&destination, nil).Once()
				openAIClient.On("GenerateAnswer", mock.Anything, mock.Anything).Return(recommendationRoute, nil).Once()
			},
			expectedResp:  expectedResp,
			expectedError: nil,
		},
		{
			name: "Destination not found",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository, openAIClient *externals.MockOpenAIClient) {
				destinationRepo.On("FindByIdInCityId", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.NotFoundError{Message: "Destinasi ke-1 tidak ditemukan"},
		},
		{
			name: "Failed get destination detail",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository, openAIClient *externals.MockOpenAIClient) {
				destinationRepo.On("FindByIdInCityId", mock.Anything, mock.Anything).Return(nil, assert.AnError).Once()
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.NotFoundError{Message: "Gagal mendapatkan detail destinasi"},
		},
		{
			name: "Failed get summarization",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository, openAIClient *externals.MockOpenAIClient) {
				destinationRepo.On("FindByIdInCityId", mock.Anything, mock.Anything).Return(&destination, nil).Once()
				openAIClient.On("GenerateAnswer", mock.Anything, mock.Anything).Return("", assert.AnError).Once()
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi rute"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDestinationRepo := new(repositories.MockDestinationRepository)
			mockOpenAIClient := new(externals.MockOpenAIClient)

			uc := usecases.NewRouteUsecase(
				nil,
				mockDestinationRepo,
				nil,
				nil,
				mockOpenAIClient,
			)

			tc.mockSetup(mockDestinationRepo, mockOpenAIClient)

			resp, err := uc.SummarizeRoute(req)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResp, resp)
			}
		})
	}
}
