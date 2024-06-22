package tests

import (
	"fmt"
	"testing"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/mocks/externals"
	"capstone/mocks/repositories"
	"capstone/usecases"

	"github.com/google/uuid"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type nearbyDestinationTestCase struct {
	name          string
	mockSetup     func(destinationRepo *repositories.MockDestinationRepository)
	expectedResp  []dto.NearbyDestination
	expectedError error
}

type popularDestinationMediaTestCase struct {
	name          string
	mockSetup     func(destinationRepo *repositories.MockDestinationRepository)
	expectedResp  []dto.PopularDestination
	expectedError error
}

type recommendatonTestCase struct {
	name          string
	mockSetup     func(redisClient *externals.MockRedisClient, destinationRepo *repositories.MockDestinationRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository, routeRepo *repositories.MockRouteRepository, openAIClient *externals.MockOpenAIClient)
	expectedResp  []dto.RecommendDestination
	expectedError error
}

func TestGetNearbyDestinations(t *testing.T) {
	const limit = 5
	destinations := []entities.Destination{
		{
			Id:   uuid.New(),
			Name: "Destination 1",
			DestinationAddress: &entities.DestinationAddress{
				Province: entities.Province{
					Name: "Jawa Timur",
				},
				City: entities.City{
					Name: "Surabaya",
				},
			},
			DestinationMedias: []entities.DestinationMedia{
				{
					Url: "https://www.google.com",
				},
			},
		},
	}
	req := dto.HomepageRequest{
		CityName: "Surabaya",
	}
	expectedResp := []dto.NearbyDestination{
		{
			Id:   destinations[0].Id,
			Name: destinations[0].Name,
			DestinationAddress: dto.Address{
				Province: dto.Province{
					Name: destinations[0].DestinationAddress.Province.Name,
				},
				City: dto.City{
					Name: destinations[0].DestinationAddress.City.Name,
				},
			},
			ImageUrl: &destinations[0].DestinationMedias[0].Url,
		},
	}

	testCases := []nearbyDestinationTestCase{
		{
			name: "Success get nearby destinations",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository) {
				destinationRepo.On("FindByProvinceName", req.CityName, limit).Return(&destinations, nil)
			},
			expectedResp:  expectedResp,
			expectedError: nil,
		},
		{
			name: "Failed get nearby destinations",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository) {
				destinationRepo.On("FindByProvinceName", req.CityName, limit).Return(nil, assert.AnError)
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi sekitar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			destinationRepo := new(repositories.MockDestinationRepository)
			usecase := usecases.NewHomepageUsecase(destinationRepo, nil, nil, nil, nil)
			tc.mockSetup(destinationRepo)
			resp, err := usecase.GetNearbyDestinations(&req)
			assert.Equal(t, tc.expectedResp, resp)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetPopularDestinations(t *testing.T) {
	destinations := []entities.Destination{
		{
			Id:   uuid.New(),
			Name: "Destination 1",
			DestinationMedias: []entities.DestinationMedia{
				{
					Id:    uuid.New(),
					Title: "vid 1",
					Url:   "https://www.google.com",
				},
			},
		},
	}

	expectedResp := []dto.PopularDestination{
		{
			Id:   destinations[0].Id,
			Name: destinations[0].Name,
			Content: dto.DestinationMedia{
				Id:    destinations[0].DestinationMedias[0].Id,
				Title: destinations[0].DestinationMedias[0].Title,
				Url:   destinations[0].DestinationMedias[0].Url,
			},
		},
	}

	testCases := []popularDestinationMediaTestCase{
		{
			name: "Success get popular destinations",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository) {
				destinationRepo.On("FindPopularDestinationVideos").Return(&destinations, nil)
			},
			expectedResp:  expectedResp,
			expectedError: nil,
		},
		{
			name: "Failed get popular destinations",
			mockSetup: func(destinationRepo *repositories.MockDestinationRepository) {
				destinationRepo.On("FindPopularDestinationVideos").Return(nil, assert.AnError)
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi populer"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			destinationRepo := new(repositories.MockDestinationRepository)
			usecase := usecases.NewHomepageUsecase(destinationRepo, nil, nil, nil, nil)
			tc.mockSetup(destinationRepo)
			resp, err := usecase.GetPopularDestinationMeadias()
			assert.Equal(t, tc.expectedResp, resp)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetRecommendations(t *testing.T) {
	uid := uuid.New()
	encodedKey := helpers.EncodeRedisKey(uid.String())

	destinations := []entities.Destination{
		{
			Id:         uuid.New(),
			Name:       "Destination 1",
			CategoryId: uuid.New(),
			DestinationMedias: []entities.DestinationMedia{
				{
					Id:    uuid.New(),
					Title: "image 1",
					Url:   "https://www.google.com",
				},
			},
			DestinationAddress: &entities.DestinationAddress{
				Province: entities.Province{
					Name: "Province 1",
				},
			},
		},
		{
			Id:   uuid.New(),
			Name: "Destination 2",
			DestinationMedias: []entities.DestinationMedia{
				{
					Id:    uuid.New(),
					Title: "image 2",
					Url:   "https://www.google.com",
				},
			},
			DestinationAddress: &entities.DestinationAddress{
				Province: entities.Province{
					Name: "Province 2",
				},
			},
		},
		{
			Id:         uuid.New(),
			Name:       "Destination 1",
			CategoryId: uuid.New(),
			DestinationMedias: []entities.DestinationMedia{
				{
					Id:    uuid.New(),
					Title: "image 1",
					Url:   "https://www.google.com",
				},
			},
			DestinationAddress: &entities.DestinationAddress{
				Province: entities.Province{
					Name: "Province 1",
				},
			},
		},
		{
			Id:   uuid.New(),
			Name: "Destination 2",
			DestinationMedias: []entities.DestinationMedia{
				{
					Id:    uuid.New(),
					Title: "image 2",
					Url:   "https://www.google.com",
				},
			},
			DestinationAddress: &entities.DestinationAddress{
				Province: entities.Province{
					Name: "Province 2",
				},
			},
		},
		{
			Id:   uuid.New(),
			Name: "Destination 2",
			DestinationMedias: []entities.DestinationMedia{
				{
					Id:    uuid.New(),
					Title: "image 2",
					Url:   "https://www.google.com",
				},
			},
			DestinationAddress: &entities.DestinationAddress{
				Province: entities.Province{
					Name: "Province 2",
				},
			},
		},
	}

	recommendationDestination := []dto.RecommendDestination{
		{
			Id:       destinations[0].Id,
			Name:     destinations[0].Name,
			ImageUrl: &destinations[0].DestinationMedias[0].Url,
			Province: dto.Province{Name: destinations[0].DestinationAddress.Province.Name},
		},
		{
			Id:       destinations[1].Id,
			Name:     destinations[1].Name,
			ImageUrl: &destinations[1].DestinationMedias[0].Url,
			Province: dto.Province{Name: destinations[1].DestinationAddress.Province.Name},
		},
		{
			Id:       destinations[2].Id,
			Name:     destinations[2].Name,
			ImageUrl: &destinations[2].DestinationMedias[0].Url,
			Province: dto.Province{Name: destinations[2].DestinationAddress.Province.Name},
		},
		{
			Id:       destinations[3].Id,
			Name:     destinations[3].Name,
			ImageUrl: &destinations[3].DestinationMedias[0].Url,
			Province: dto.Province{Name: destinations[3].DestinationAddress.Province.Name},
		},
		{
			Id:       destinations[4].Id,
			Name:     destinations[4].Name,
			ImageUrl: &destinations[4].DestinationMedias[0].Url,
			Province: dto.Province{Name: destinations[4].DestinationAddress.Province.Name},
		},
	}
	ids := &[]string{"1", "2", "3", "4", "5"}
	personalizations := entities.UserPersonalization{
		Id: uuid.New(),
		PersonalizationProvinces: []entities.PersonalizationProvince{
			{
				ProvinceId: destinations[0].DestinationAddress.Province.Id,
			},
		},
		PersonalizationCategories: []entities.PersonalizationCategory{
			{
				CategoryId: destinations[0].CategoryId,
			},
		},
	}

	testCases := []recommendatonTestCase{
		{
			name: "Success get recommendations from Redis",
			mockSetup: func(redisClient *externals.MockRedisClient, destinationRepo *repositories.MockDestinationRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository, routeRepo *repositories.MockRouteRepository, openAIClient *externals.MockOpenAIClient) {
				redisClient.On("GetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey)).Return(ids, nil)
				destinationRepo.On("FindByManyIds", *ids).Return(&destinations, nil)
			},
			expectedResp:  recommendationDestination,
			expectedError: nil,
		},
		{
			name: "Failed to get recommendations from Redis",
			mockSetup: func(redisClient *externals.MockRedisClient, destinationRepo *repositories.MockDestinationRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository, routeRepo *repositories.MockRouteRepository, openAIClient *externals.MockOpenAIClient) {
				redisClient.On("GetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey)).Return(nil, assert.AnError)
				destinationRepo.On("FindByManyIds", *ids).Return(&destinations, nil)
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"},
		},
		{
			name: "Error find by many ids",
			mockSetup: func(redisClient *externals.MockRedisClient, destinationRepo *repositories.MockDestinationRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository, routeRepo *repositories.MockRouteRepository, openAIClient *externals.MockOpenAIClient) {
				redisClient.On("GetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey)).Return(ids, nil)
				destinationRepo.On("FindByManyIds", *ids).Return(nil, assert.AnError)
			},
			expectedResp:  nil,
			expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"},
		},
		{
			name: "Success get recommendations using personalization data",
			mockSetup: func(redisClient *externals.MockRedisClient, destinationRepo *repositories.MockDestinationRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository, routeRepo *repositories.MockRouteRepository, openAIClient *externals.MockOpenAIClient) {
				subquery := new(gorm.DB)

				redisClient.On("GetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey)).Return(nil, goredis.Nil)
				userPersonalizationRepo.On("FindByUserId", uid).Return(&personalizations, nil)
				routeRepo.On("FindVisitedByUserSubquery", uid).Return(subquery)
				destinationRepo.On("FindBySubqueryRoute", mock.Anything, true, mock.Anything, mock.Anything).Return(&[]entities.Destination{}, nil)
				destinationRepo.On("FindBySubqueryRoute", mock.Anything, false, mock.Anything, mock.Anything).Return(&destinations, nil)
				openAIClient.On("GenerateAnswer", mock.Anything, mock.Anything).Return("1\n2\n3\n4\n5", nil)
				redisClient.On("SetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey), mock.Anything).Return(nil)
				destinationRepo.On("FindByManyIds", *ids).Return(&destinations, nil)
			},
			expectedResp:  recommendationDestination,
			expectedError: nil,
		},
		// {
		// 	name: "Failed to get personalization data",
		// 	mockSetup: func(redisClient *externals.MockRedisClient, destinationRepo *repositories.MockDestinationRepository, userPersonalizationRepo *repositories.MockUserPersonalizationRepository, routeRepo *repositories.MockRouteRepository, openAIClient *externals.MockOpenAIClient) {
		// 		redisClient.On("GetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey)).Return(nil, redis.Nil)
		// 		userPersonalizationRepo.On("FindByUserId", uid).Return(nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data personalisasi"})
		// 	},
		// 	expectedResp:  nil,
		// 	expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data personalisasi"},
		// },
		// {
		// 	name: "Failed to generate recommendations from OpenAI",
		// 	mockSetup: func(redisClient *externals.MockRedisClient, destinationRepo *externals.MockDestinationRepo, userPersonalizationRepo *externals.MockUserPersonalizationRepo, routeRepo *externals.MockRouteRepo, openAIClient *externals.MockOpenAIClient) {
		// 		redisClient.On("GetRecommendedDestinationsIds", fmt.Sprintf("rec_%s", encodedKey)).Return(nil, redis.Nil)
		// 		userPersonalizationRepo.On("FindByUserId", uid).Return(&entities.UserPersonalization{}, nil)
		// 		destinationRepo.On("FindBySubqueryRoute", mock.Anything, true, mock.Anything, mock.Anything).Return(&mockRecommendations, nil)
		// 		destinationRepo.On("FindBySubqueryRoute", mock.Anything, false, mock.Anything, mock.Anything).Return(&mockRecommendations, nil)
		// 		openAIClient.On("GenerateAnswer", mock.Anything, mock.Anything).Return("", &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"})
		// 	},
		// 	expectedResp:  nil,
		// 	expectedError: &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redisClient := new(externals.MockRedisClient)
			destinationRepo := new(repositories.MockDestinationRepository)
			userPersonalizationRepo := new(repositories.MockUserPersonalizationRepository)
			routeRepo := new(repositories.MockRouteRepository)
			openAIClient := new(externals.MockOpenAIClient)

			tc.mockSetup(redisClient, destinationRepo, userPersonalizationRepo, routeRepo, openAIClient)

			uc := usecases.NewHomepageUsecase(destinationRepo, routeRepo, userPersonalizationRepo, redisClient, openAIClient)
			resp, err := uc.GetRecommendations(uid)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResp, resp)
			}
		})
	}
}
