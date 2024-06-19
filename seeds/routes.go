package seeds

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedRoutes() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	routes := generateRoutes(s, 120)
	routeDetails := generateRouteDetails(s, routes, 3)

	saveRoutes(s, routes)
	saveRouteDetails(s, routeDetails)
}

func generateRoutes(s Seed, count int) []entities.Route {
	var routes []entities.Route

	for i := 0; i < count; i++ {
		user, err := getRandomUser(s)
		if err != nil {
			log.Fatalf("failed to get user: %v", err)
		}

		destination, err := getRandomDestination(s)
		if err != nil {
			log.Fatalf("failed to get destinations: %v", err)
		}

		if destination == nil || destination.DestinationAddress == nil {
			log.Fatalf("destination or destination address is nil")
		}

		route := entities.Route{
			Id:             uuid.New(),
			UserId:         user.Id,
			CityId:         destination.DestinationAddress.CityId,
			Name:           "Route ke " + strconv.Itoa(i),
			StartLocation:  "Hotel " + strconv.Itoa(i),
			StartLongitude: getRandomLongitude(),
			StartLatitude:  getRandomLatitude(),
			Price:          getRandomPrice(100000, 500000),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		routes = append(routes, route)
	}

	return routes
}

func generateRouteDetails(s Seed, routes []entities.Route, detailCount int) []entities.RouteDetail {
	var routeDetails []entities.RouteDetail

	for _, route := range routes {
		for j := 1; j <= detailCount; j++ {
			destination, err := getRandomDestination(s)
			if err != nil {
				log.Fatalf("failed to get destinations: %v", err)
			}

			if destination == nil || destination.DestinationAddress == nil {
				log.Fatalf("destination or destination address is nil")
			}

			routedetail := entities.RouteDetail{
				Id:            uuid.New(),
				DestinationId: destination.Id,
				RouteId:       route.Id,
				Longitude:     destination.Longitude,
				Latitude:      destination.Latitude,
				Duration:      60 * j, // Durasi yang berbeda untuk setiap detail rute
				Order:         j,
				VisitStart:    []byte("08:00"),
				VisitEnd:      []byte("09:00"),
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}

			routeDetails = append(routeDetails, routedetail)
		}
	}

	return routeDetails
}

func saveRoutes(s Seed, routes []entities.Route) {
	for _, route := range routes {
		if err := s.db.Save(&route).Error; err != nil {
			log.Fatalf("failed to create route: %v", err)
		}
	}
}

func saveRouteDetails(s Seed, routeDetails []entities.RouteDetail) {
	for _, routeDetail := range routeDetails {
		if err := s.db.Save(&routeDetail).Error; err != nil {
			log.Fatalf("failed to create route detail: %v", err)
		}
	}
}

func getRandomDestination(s Seed) (*entities.Destination, error) {
	var destination entities.Destination

	if err := s.db.Order("RAND()").Preload("DestinationAddress").First(&destination).Error; err != nil {
		return nil, err
	}

	return &destination, nil
}

func getRandomUser(s Seed) (*entities.User, error) {
	var user entities.User

	if err := s.db.Order("RAND()").First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func getRandomLongitude() float64 {
	return 110.0 + rand.Float64()*(120.0-110.0)
}

func getRandomLatitude() float64 {
	return -10.0 + rand.Float64()*(10.0-(-10.0))
}

func getRandomPrice(min, max int) float64 {
	return float64(((rand.Intn((max-min)/1000+1) * 1000) + min))
}
