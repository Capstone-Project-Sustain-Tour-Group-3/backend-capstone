package usecases

import (
	"errors"
	"fmt"
	"strings"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"
	"capstone/externals/openai"
	exredis "capstone/externals/redis"
	"capstone/helpers"
	"capstone/repositories"

	"github.com/redis/go-redis/v9"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HomepageUsecase interface {
	GetNearbyDestinations(req *dto.HomepageRequest) ([]dto.NearbyDestination, error)
	GetPopularDestinationMeadias() ([]dto.PopularDestination, error)
	GetRecommendations(uid uuid.UUID) ([]dto.RecommendDestination, error)
}

type homepageUsecase struct {
	destinationRepo         repositories.IDestinationRepository
	routeRepo               repositories.RouteRepository
	userPersonalizationRepo repositories.UserPersonalizationRepository
	redisClient             exredis.RedisClient
	openAIClient            openai.OpenAIClient
}

func NewHomepageUsecase(
	destinationRepo repositories.IDestinationRepository,
	routeRepo repositories.RouteRepository,
	userPersonalizationRepo repositories.UserPersonalizationRepository,
	redisClient exredis.RedisClient,
	openAIClient openai.OpenAIClient,
) *homepageUsecase {
	return &homepageUsecase{
		destinationRepo:         destinationRepo,
		routeRepo:               routeRepo,
		userPersonalizationRepo: userPersonalizationRepo,
		redisClient:             redisClient,
		openAIClient:            openAIClient,
	}
}

func (uc *homepageUsecase) GetNearbyDestinations(req *dto.HomepageRequest) ([]dto.NearbyDestination, error) {
	const limit = 5

	nearbyDestination, err := uc.destinationRepo.FindByProvinceName(req.CityName, limit)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi sekitar"}
	}

	return dto.ToNearbyDestinations(nearbyDestination), nil
}

func (uc *homepageUsecase) GetPopularDestinationMeadias() ([]dto.PopularDestination, error) {
	popularDestinations, err := uc.destinationRepo.FindPopularDestinationVideos()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi populer"}
	}

	return dto.ToPopularDestinations(popularDestinations), nil
}

func (uc *homepageUsecase) GetRecommendations(uid uuid.UUID) ([]dto.RecommendDestination, error) {
	var recommendDestinations *[]entities.Destination

	encodedKey := helpers.EncodeRedisKey(uid.String())

	ids, err := uc.redisClient.GetRecommendedDestinationsIds(fmt.Sprint("rec_", encodedKey))
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"}
	}

	if ids != nil {
		recommendDestinations, err = uc.destinationRepo.FindByManyIds(*ids)
		if err != nil {
			return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"}
		}

		return dto.ToRecommendDestinations(recommendDestinations), nil
	}

	personalization, err := uc.userPersonalizationRepo.FindByUserId(uid)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, &errorHandlers.NotFoundError{Message: "Personalisasi belum dibuat"}
		default:
			return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data personalisasi"}
		}
	}

	provinceIds := dto.ToPersonalizationProvinceIds(&personalization.PersonalizationProvinces)
	categoryIds := dto.ToCategoryIds(&personalization.PersonalizationCategories)

	var visitedDestinations, unVisitedDestinations *[]entities.Destination
	subquery := uc.routeRepo.FindVisitedByUserSubquery(uid)
	visitedDestinations, err = uc.destinationRepo.FindBySubqueryRoute(
		subquery,
		true,
		provinceIds,
		categoryIds,
	)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data personalisasi"}
	}

	unVisitedDestinations, err = uc.destinationRepo.FindBySubqueryRoute(
		subquery,
		false,
		provinceIds,
		categoryIds,
	)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan data personalisasi"}
	}

	if len(*unVisitedDestinations) < 5 {
		ids = helpers.SelectRecommendationIds(visitedDestinations, unVisitedDestinations)
	} else {
		var recommendations string
		prompt := helpers.ToRecommendationPrompt(visitedDestinations, unVisitedDestinations)
		recommendations, err = uc.openAIClient.GenerateAnswer(
			prompt,
			helpers.GetRecommendationSystemInstruction(),
		)
		if err != nil {
			return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"}
		}

		splittedIds := strings.Split(recommendations, "\n")
		ids = &splittedIds
	}

	err = uc.redisClient.SetRecommendedDestinationsIds(fmt.Sprint("rec_", encodedKey), *ids)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal menyimpan rekomendasi destinasi"}
	}

	recommendedDestinations, err := uc.destinationRepo.FindByManyIds(*ids)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi destinasi"}
	}

	return dto.ToRecommendDestinations(recommendedDestinations), nil
}
