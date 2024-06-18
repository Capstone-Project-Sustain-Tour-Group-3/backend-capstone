package tests

import (
	"errors"
	"testing"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/mocks/externals"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSearchDestinations(t *testing.T) {
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockDestinationFacilityRepo := new(repositories.MockDestinationFacilityRepository)
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationAddressRepo := new(repositories.MockDestinationAddressRepository)
	mockCategoryRepo := new(repositories.MockCategoryRepository)
	mockFacilityRepo := new(repositories.MockFacilityRepository)
	mockProvinceRepo := new(repositories.MockProvinceRepository)
	mockCityRepo := new(repositories.MockCityRepository)
	mockSubdistrictRepo := new(repositories.MockSubdistrictRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationUsecase(
		mockDestinationRepo,
		mockDestinationFacilityRepo,
		mockDestinationMediaRepo,
		mockDestinationAddressRepo,
		mockCategoryRepo,
		mockFacilityRepo,
		mockProvinceRepo,
		mockCityRepo,
		mockSubdistrictRepo,
		mockCloudinaryClient,
	)

	testCases := []struct {
		name                 string
		page                 int
		limit                int
		searchQuery          string
		sortQuery            string
		filterQuery          string
		mockSetup            func()
		expectedError        error
		expectedResponse     *[]dto.SearchDestination
		expectedTotal        *int64
		expectedCategoryName string
	}{
		{
			name:        "Success search destinations",
			page:        1,
			limit:       10,
			searchQuery: "danau",
			sortQuery:   "populer",
			filterQuery: "ABC123",
			mockSetup: func() {
				mockDestinationRepo.On("FindAll", 1, 10, "danau", "populer", "ABC123").
					Return("Alam", func() *int64 { i := int64(5); return &i }(), []entities.Destination{
						{
							Id:       uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							Name:     "Beach 1",
							Category: entities.Category{},
							DestinationAddress: &entities.DestinationAddress{
								Province: entities.Province{},
								City:     entities.City{},
							},
							DestinationMedias: []entities.DestinationMedia{},
						},
					}, nil).Once()
			},
			expectedError: nil,
			expectedResponse: &[]dto.SearchDestination{
				{
					Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
					Name: "Beach 1",
				},
			},
			expectedTotal:        func() *int64 { i := int64(5); return &i }(),
			expectedCategoryName: "Alam",
		},
		{
			name:        "Error internal server",
			page:        1,
			limit:       10,
			searchQuery: "danau",
			sortQuery:   "populer",
			filterQuery: "ABC123",
			mockSetup: func() {
				mockDestinationRepo.On("FindAll", 1, 10, "danau", "populer", "ABC123").
					Return("", nil, nil, errors.New("internal error")).Once()
			},
			expectedError:        &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data destinasi"},
			expectedResponse:     nil,
			expectedTotal:        nil,
			expectedCategoryName: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			categoryName, total, response, err := uc.SearchDestinations(tc.page, tc.limit, tc.searchQuery, tc.sortQuery, tc.filterQuery)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedCategoryName, categoryName)
				require.Equal(t, tc.expectedTotal, total)
				require.Equal(t, tc.expectedResponse, response)
			}

			mockDestinationRepo.AssertExpectations(t)
		})
	}
}

func TestDetailDestination(t *testing.T) {
	mockDestinationRepo := new(repositories.MockDestinationRepository)
	mockDestinationFacilityRepo := new(repositories.MockDestinationFacilityRepository)
	mockDestinationMediaRepo := new(repositories.MockDestinationMediaRepository)
	mockDestinationAddressRepo := new(repositories.MockDestinationAddressRepository)
	mockCategoryRepo := new(repositories.MockCategoryRepository)
	mockFacilityRepo := new(repositories.MockFacilityRepository)
	mockProvinceRepo := new(repositories.MockProvinceRepository)
	mockCityRepo := new(repositories.MockCityRepository)
	mockSubdistrictRepo := new(repositories.MockSubdistrictRepository)
	mockCloudinaryClient := new(externals.MockCloudinaryClient)

	uc := usecases.NewDestinationUsecase(
		mockDestinationRepo,
		mockDestinationFacilityRepo,
		mockDestinationMediaRepo,
		mockDestinationAddressRepo,
		mockCategoryRepo,
		mockFacilityRepo,
		mockProvinceRepo,
		mockCityRepo,
		mockSubdistrictRepo,
		mockCloudinaryClient,
	)

	url := "https://www.google.com"

	testCases := []struct {
		name                        string
		id                          uuid.UUID
		mockDestinationRepoSetup    func()
		mockSimilarDestinationSetup func()
		expectedError               error
		expectedResponse            *dto.DetailDestinationResponse
	}{
		{
			name: "Success get detail destination",
			id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindById", mock.Anything).Return(&entities.Destination{
					Id:          uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
					Name:        "Destination 1",
					OpenTime:    "08:00",
					CloseTime:   "17:00",
					EntryPrice:  50000,
					VisitCount:  100,
					Description: "Description 1",
					DestinationAddress: &entities.DestinationAddress{
						Province:    entities.Province{Name: "Province 1"},
						City:        entities.City{Name: "City 1"},
						Subdistrict: entities.Subdistrict{Name: "Subdistrict 1"},
						StreetName:  "Street 1",
						PostalCode:  "12345",
					},
					Category: entities.Category{
						Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Name: "Alam",
					},
					DestinationMedias: []entities.DestinationMedia{
						{
							Id:            uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							Url:           "https://www.google.com",
							Type:          "video",
							Title:         "video",
						},
						{
							Id:            uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							Url:           "https://www.google.com",
							Type:          "image",
							Title:         "image",
						},
					},
					DestinationFacilities: &[]entities.DestinationFacility{
						{
							Id:            uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							FacilityId:    uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							Facility: entities.Facility{
								Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
								Name: "Parkir",
								Url:  "https://www.google.com",
							},
						},
					},
				}, nil).Once()
			},
			mockSimilarDestinationSetup: func() {
				mockDestinationRepo.On("FindByCategoryId", mock.Anything).Return([]entities.Destination{
					{
						Id:         uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
						Name:       "Pantai",
						CategoryId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Category:   entities.Category{Id: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"), Name: "Alam"},
						DestinationAddress: &entities.DestinationAddress{
							Province:    entities.Province{Name: "Province 2"},
							City:        entities.City{Name: "City 2"},
							Subdistrict: entities.Subdistrict{Name: "Subdistrict 2"},
							StreetName:  "Street 2",
							PostalCode:  "54321",
						},
						DestinationMedias: []entities.DestinationMedia{
							{
								Id:            uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
								DestinationId: uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
								Destination:   entities.Destination{},
								Url:           "https://www.google.com",
								Type:          "image",
								Title:         "image",
							},
						},
						VisitCount: 100,
					},
				}, nil).Once()
			},
			expectedError: nil,
			expectedResponse: &dto.DetailDestinationResponse{
				Id:          uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
				Name:        "Destination 1",
				OpenTime:    "08:00",
				CloseTime:   "17:00",
				EntryPrice:  50000,
				VisitCount:  100,
				Description: "Description 1",
				DestinationAddress: &dto.DestinationAddress{
					Province:    "Province 1",
					City:        "City 1",
					Subdistrict: "Subdistrict 1",
					StreetName:  "Street 1",
					PostalCode:  "12345",
				},
				UrlImages: &[]dto.UrlImage{
					{
						Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Url:  "https://www.google.com",
						Type: "image",
					},
				},
				UrlVideos: &[]dto.UrlVideo{
					{
						Id:    uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Url:   "https://www.google.com",
						Type:  "video",
						Title: "video",
					},
				},
				Categories: &dto.Category{
					Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
					Name: "Alam",
				},
				Facilities: &[]dto.Facility{
					{
						Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Name: "Parkir",
						Url:  "https://www.google.com",
					},
				},
				SimilarDestination: &[]dto.SearchDestination{
					{
						Id:         uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Name:       "Pantai",
						Url:        &url,
						City:       "City 2",
						VisitCount: 100,
						Category: dto.Category{
							Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
							Name: "Alam",
						},
					},
				},
			},
		},
		{
			name: "Error destination not found",
			id:   uuid.MustParse("cc88d401-47bc-4a73-9e7c-b7d84168954c"),
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindById", mock.Anything).Return(nil, errors.New("not found")).Once()
			},
			mockSimilarDestinationSetup: func() {},
			expectedError: &errorHandlers.NotFoundError{
				Message: "Destinasi tidak ditemukan",
			},
			expectedResponse: nil,
		},
		{
			name: "Error similar destination not found",
			id:   uuid.MustParse("dd99d401-47bc-4a73-9e7c-b7d84168954d"),
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindById", mock.Anything).Return(&entities.Destination{
					Id:          uuid.MustParse("dd99d401-47bc-4a73-9e7c-b7d84168954d"),
					Name:        "Destination 3",
					OpenTime:    "10:00",
					CloseTime:   "19:00",
					EntryPrice:  100000,
					VisitCount:  120,
					Description: "Description 3",
					DestinationAddress: &entities.DestinationAddress{
						Province:    entities.Province{Name: "Province 3"},
						City:        entities.City{Name: "City 3"},
						Subdistrict: entities.Subdistrict{Name: "Subdistrict 3"},
						StreetName:  "Street 3",
						PostalCode:  "67890",
					},
					Category: entities.Category{Name: "Alam"},
				}, nil).Once()
				mockDestinationRepo.On("FindByCategoryId", mock.Anything).Return(nil, errors.New("not found")).Once()
			},
			mockSimilarDestinationSetup: func() {},
			expectedError: &errorHandlers.InternalServerError{
				Message: "Gagal mendapatkan destinasi serupa",
			},
			expectedResponse: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockDestinationRepoSetup()
			tc.mockSimilarDestinationSetup()

			response, err := uc.DetailDestination(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			}

			mockDestinationRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllDestinations(t *testing.T) {
	mockDestinationRepo := new(repositories.MockDestinationRepository)

	uc := usecases.NewDestinationUsecase(
		mockDestinationRepo,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	)

	mockDestinations := []entities.Destination{
		{
			Id:         uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
			Name:       "Destination 1",
			OpenTime:   "08:00",
			CloseTime:  "17:00",
			VisitCount: 100,
			EntryPrice: 50000,
			DestinationAddress: &entities.DestinationAddress{
				Province:    entities.Province{Name: "Province 1"},
				City:        entities.City{Name: "City 1"},
				Subdistrict: entities.Subdistrict{Name: "Subdistrict 1"},
				StreetName:  "Street 1",
				PostalCode:  "12345",
			},
			CategoryId: uuid.New(),
			Category:   entities.Category{Name: "Alam"},
		},
		{
			Id:         uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
			Name:       "Destination 2",
			OpenTime:   "09:00",
			CloseTime:  "18:00",
			VisitCount: 80,
			EntryPrice: 75000,
			DestinationAddress: &entities.DestinationAddress{
				Province:    entities.Province{Name: "Province 2"},
				City:        entities.City{Name: "City 2"},
				Subdistrict: entities.Subdistrict{Name: "Subdistrict 2"},
				StreetName:  "Street 2",
				PostalCode:  "54321",
			},
			CategoryId: uuid.New(),
			Category:   entities.Category{Name: "Budaya"},
		},
	}
	total := int64(2)

	testCases := []struct {
		name                     string
		page                     int
		limit                    int
		searchQuery              string
		mockDestinationRepoSetup func()
		expectedError            error
		expectedResponse         *[]dto.GetAllDestination
		expectedTotal            int64
	}{
		{
			name:        "Success get all destinations",
			page:        1,
			limit:       10,
			searchQuery: "Destination",
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindAll", 1, 10, "Destination", "", "").Return("Alam", &total, mockDestinations, nil).Once()
			},
			expectedError: nil,
			expectedResponse: &[]dto.GetAllDestination{
				{
					Id:         uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
					Name:       "Destination 1",
					OpenTime:   "08:00",
					CloseTime:  "17:00",
					VisitCount: 100,
					EntryPrice: 50000,
					DestinationAddress: &dto.DestinationAddress{
						Province:    "Province 1",
						City:        "City 1",
						Subdistrict: "Subdistrict 1",
						StreetName:  "Street 1",
						PostalCode:  "12345",
					},
					Category: dto.Category{
						Id:   uuid.New(),
						Name: "Alam",
					},
				},
				{
					Id:         uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
					Name:       "Destination 2",
					OpenTime:   "09:00",
					CloseTime:  "18:00",
					VisitCount: 80,
					EntryPrice: 75000,
					DestinationAddress: &dto.DestinationAddress{
						Province:    "Province 2",
						City:        "City 2",
						Subdistrict: "Subdistrict 2",
						StreetName:  "Street 2",
						PostalCode:  "54321",
					},
					Category: dto.Category{
						Id:   uuid.New(),
						Name: "Budaya",
					},
				},
			},
			expectedTotal: 2,
		},
		{
			name:        "Error get all destinations",
			page:        1,
			limit:       10,
			searchQuery: "Destination",
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindAll", 1, 10, "Destination", "", "").Return("", int64(0), nil, errors.New("internal error")).Once()
			},
			expectedError:    errors.New("internal error"),
			expectedResponse: nil,
			expectedTotal:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockDestinationRepoSetup()

			totalRes, response, err := uc.GetAllDestinations(tc.page, tc.limit, tc.searchQuery)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, response)
				require.Equal(t, tc.expectedTotal, *totalRes)
			}

			mockDestinationRepo.AssertExpectations(t)
		})
	}
}

func TestGetDestinationById(t *testing.T) {
	mockDestinationRepo := new(repositories.MockDestinationRepository)

	uc := usecases.NewDestinationUsecase(
		mockDestinationRepo,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	)

	testCases := []struct {
		name                     string
		id                       uuid.UUID
		mockDestinationRepoSetup func()
		expectedError            error
		expectedResponse         *dto.GetByIdDestinationResponse
	}{
		{
			name: "Success get destination by id",
			id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
			mockDestinationRepoSetup: func() {
				mockDestination := entities.Destination{
					Id:          uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
					Name:        "Destination 1",
					OpenTime:    "08:00",
					CloseTime:   "17:00",
					VisitCount:  100,
					EntryPrice:  50000,
					Description: "Beautiful place",
					DestinationAddress: &entities.DestinationAddress{
						Province:    entities.Province{Name: "Province 1"},
						City:        entities.City{Name: "City 1"},
						Subdistrict: entities.Subdistrict{Name: "Subdistrict 1"},
						StreetName:  "Street 1",
						PostalCode:  "12345",
					},
					DestinationMedias: []entities.DestinationMedia{
						{Id: uuid.New(), Url: "http://example.com/image1.jpg", Type: "image"},
						{Id: uuid.New(), Url: "http://example.com/image2.jpg", Type: "image"},
					},
					CategoryId: uuid.New(),
					Category:   entities.Category{Name: "Alam"},
					DestinationFacilities: &[]entities.DestinationFacility{
						{FacilityId: uuid.New(), Facility: entities.Facility{Name: "WiFi", Url: "http://example.com/wifi.jpg"}},
						{FacilityId: uuid.New(), Facility: entities.Facility{Name: "Parking", Url: "http://example.com/parking.jpg"}},
					},
				}

				mockDestinationRepo.On("FindById", mockDestination.Id).Return(&mockDestination, nil).Once()
			},
			expectedError: nil,
			expectedResponse: &dto.GetByIdDestinationResponse{
				Id:          uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
				Name:        "Destination 1",
				OpenTime:    "08:00",
				CloseTime:   "17:00",
				VisitCount:  100,
				EntryPrice:  50000,
				Description: "Beautiful place",
				DestinationAddress: &dto.DestinationAddress{
					Province:    "Province 1",
					City:        "City 1",
					Subdistrict: "Subdistrict 1",
					StreetName:  "Street 1",
					PostalCode:  "12345",
				},
				UrlImages: &[]dto.UrlImage{
					{Id: uuid.New(), Url: "http://example.com/image1.jpg", Type: "image"},
					{Id: uuid.New(), Url: "http://example.com/image2.jpg", Type: "image"},
				},
				Categories: &dto.Category{
					Id:   uuid.New(),
					Name: "Alam",
				},
				Facilities: &[]dto.Facility{
					{Id: uuid.New(), Name: "WiFi", Url: "http://example.com/wifi.jpg"},
					{Id: uuid.New(), Name: "Parking", Url: "http://example.com/parking.jpg"},
				},
			},
		},
		{
			name: "Destination not found",
			id:   uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindById", uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b")).Return(nil, errors.New("not found")).Once()
			},
			expectedError:    &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"},
			expectedResponse: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockDestinationRepoSetup()

			response, err := uc.GetDestinationById(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			}

			mockDestinationRepo.AssertExpectations(t)
		})
	}
}

func TestIncrementVisitCount(t *testing.T) {
	mockDestinationRepo := new(repositories.MockDestinationRepository)

	uc := usecases.NewDestinationUsecase(
		mockDestinationRepo,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	)

	testCases := []struct {
		name                     string
		id                       uuid.UUID
		mockDestinationRepoSetup func()
		expectedError            error
	}{
		{
			name: "Success increment visit count",
			id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
			mockDestinationRepoSetup: func() {
				mockDestination := entities.Destination{
					Id:         uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
					Name:       "Destination 1",
					VisitCount: 100,
				}

				mockDestinationRepo.On("FindById", mockDestination.Id).Return(&mockDestination, nil).Once()
				mockDestinationRepo.On("Update", mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Destination not found",
			id:   uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindById", uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b")).Return(nil, errors.New("not found")).Once()
			},
			expectedError: errors.New("not found"),
		},
		{
			name: "Update error",
			id:   uuid.MustParse("cc88d401-47bc-4a73-9e7c-b7d84168954c"),
			mockDestinationRepoSetup: func() {
				mockDestination := entities.Destination{
					Id:         uuid.MustParse("cc88d401-47bc-4a73-9e7c-b7d84168954c"),
					Name:       "Destination 3",
					VisitCount: 50,
				}

				mockDestinationRepo.On("FindById", mockDestination.Id).Return(&mockDestination, nil).Once()
				mockDestinationRepo.On("Update", mock.Anything).Return(errors.New("update error")).Once()
			},
			expectedError: errors.New("update error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockDestinationRepoSetup()

			err := uc.IncrementVisitCount(tc.id)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			mockDestinationRepo.AssertExpectations(t)
		})
	}
}

func TestGetCitiesWithDestinations(t *testing.T) {
	mockCityRepo := new(repositories.MockCityRepository)

	uc := usecases.NewDestinationUsecase(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		mockCityRepo,
		nil,
		nil,
	)

	testCases := []struct {
		name              string
		mockCityRepoSetup func()
		expectedResponse  *[]dto.Cities
		expectedError     error
	}{
		{
			name: "Success get cities with destinations",
			mockCityRepoSetup: func() {
				mockCities := []entities.City{
					{
						Id:   "ABC1",
						Name: "City 1",
					},
					{
						Id:   "ABC2",
						Name: "City 2",
					},
				}

				mockCityRepo.On("GetCitiesWithDestinations").Return(mockCities, nil).Once()
			},
			expectedResponse: &[]dto.Cities{
				{
					Id:   "ABC1",
					Nama: "City 1",
				},
				{
					Id:   "ABC2",
					Nama: "City 2",
				},
			},
			expectedError: nil,
		},
		{
			name: "Internal server error",
			mockCityRepoSetup: func() {
				mockCityRepo.On("GetCitiesWithDestinations").Return(nil, errors.New("internal error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data kota"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCityRepoSetup()

			response, err := uc.GetCitiesWithDestinations()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, response)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResponse, response)
			}

			mockCityRepo.AssertExpectations(t)
		})
	}
}

func TestGetDestinationByCityId(t *testing.T) {
	mockDestinationRepo := new(repositories.MockDestinationRepository)

	uc := usecases.NewDestinationUsecase(
		mockDestinationRepo,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	)

	testCases := []struct {
		name                     string
		cityId                   string
		mockDestinationRepoSetup func()
		expectedResponse         *[]dto.DestinationsByCity
		expectedError            error
	}{
		{
			name:   "Success get destinations by city ID",
			cityId: "1",
			mockDestinationRepoSetup: func() {
				mockDestinations := []entities.Destination{
					{
						Id:   uuid.MustParse("aa66d401-47bc-4a73-9e7c-b7d84168954a"),
						Name: "Destination 1",
						DestinationAddress: &entities.DestinationAddress{
							StreetName:  "Street 1",
							Subdistrict: entities.Subdistrict{Name: "Subdistrict 1"},
							City:        entities.City{Name: "City 1"},
							Province:    entities.Province{Name: "Province 1"},
						},
					},
					{
						Id:   uuid.MustParse("bb77d401-47bc-4a73-9e7c-b7d84168954b"),
						Name: "Destination 2",
						DestinationAddress: &entities.DestinationAddress{
							StreetName:  "Street 2",
							Subdistrict: entities.Subdistrict{Name: "Subdistrict 2"},
							City:        entities.City{Name: "City 2"},
							Province:    entities.Province{Name: "Province 2"},
						},
					},
				}

				mockDestinationRepo.On("FindDestinationByCityId", "1").Return(mockDestinations, nil).Once()
			},
			expectedResponse: &[]dto.DestinationsByCity{
				{
					Id:            "aa66d401-47bc-4a73-9e7c-b7d84168954a",
					NamaDestinasi: "Destination 1",
					NamaJalan:     "Street 1",
					Kecamatan:     "Subdistrict 1",
					Kota:          "City 1",
					Provinsi:      "Province 1",
				},
				{
					Id:            "bb77d401-47bc-4a73-9e7c-b7d84168954b",
					NamaDestinasi: "Destination 2",
					NamaJalan:     "Street 2",
					Kecamatan:     "Subdistrict 2",
					Kota:          "City 2",
					Provinsi:      "Province 2",
				},
			},
			expectedError: nil,
		},
		{
			name:   "No destinations found",
			cityId: "2",
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindDestinationByCityId", "2").Return(nil, nil).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.NotFoundError{Message: "Destinasi tidak ditemukan"},
		},
		{
			name:   "Internal server error",
			cityId: "3",
			mockDestinationRepoSetup: func() {
				mockDestinationRepo.On("FindDestinationByCityId", "3").Return(nil, errors.New("internal error")).Once()
			},
			expectedResponse: nil,
			expectedError:    &errorHandlers.InternalServerError{Message: "Gagal untuk menemukan data destinasi"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockDestinationRepoSetup()

			response, err := uc.GetDestinationByCityId(tc.cityId)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, response)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResponse, response)
			}

			mockDestinationRepo.AssertExpectations(t)
		})
	}
}
