package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/externals/redis"
	"capstone/repositories"

	"github.com/google/uuid"
)

type HomepageUsecase interface {
	GetHomepageData(req *dto.HomepageRequest, uid uuid.UUID) (*dto.HomepageResponse, error)
}

type homepageUsecase struct {
	destinationRepo repositories.IDestinationRepository
	routeRepo       repositories.RouteRepository
	redisClient     redis.RedisClient
}

func NewHomepageUsecase(
	destinationRepo repositories.IDestinationRepository,
	routeRepo repositories.RouteRepository,
	redisClient redis.RedisClient,
) *homepageUsecase {
	return &homepageUsecase{
		destinationRepo: destinationRepo,
		routeRepo:       routeRepo,
		redisClient:     redisClient,
	}
}

func (uc *homepageUsecase) GetHomepageData(req *dto.HomepageRequest, uid uuid.UUID) (*dto.HomepageResponse, error) {
	nearbyDestination, err := uc.destinationRepo.FindByProvinceName(req.CityName)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi sekitar"}
	}

	popularDestinationsMedia, err := uc.destinationRepo.FindPopularDestinationVideos()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan destinasi populer"}
	}

	// subquery := uc.routeRepo.FindVisitedByUserSubquery(uid)

	// ids, err := uc.redisClient.GetRecommendedDestinationsIds(
	// 	fmt.Sprint("rec_", helpers.EncodeRedisKey(uid.String())),
	// )
	// if err != nil {
	// 	return nil, &errorHandlers.InternalServerError{Message: "Gagal mendapatkan rekomendasi"}
	// }

	// if len(*ids) == 0 {
	// 	// TODO : get destionation that has been visited by user

	// 	// TODO: get recommendation from open ai

	// 	// TODO: set recommendation
	// }

	res, err := dto.ToHomepageResponse(
		nearbyDestination,
		popularDestinationsMedia,
	)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "Gagal melakukan mapping data"}
	}

	return res, nil
}
