package seeds

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedRoutes() {
	destination, err := getRandomDestination(s)
	fmt.Println("Destination: ", destination)
	if err != nil {
		log.Fatalf("failed to get destinations: %v", err)
	}

	if destination == nil || destination.DestinationAddress == nil {
		log.Fatalf("destination or destination address is nil")
	}

	user, err := getRandomUser(s)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	var routes []entities.Route
	var routeDetails []entities.RouteDetail

	for i := range [15]int{} {
		route := entities.Route{
			Id:             uuid.New(),
			UserId:         user.Id,
			CityId:         destination.DestinationAddress.CityId,
			Name:           "Route ke " + strconv.Itoa(i),
			StartLocation:  "Hotel " + strconv.Itoa(i),
			StartLongitude: 12.01,
			StartLatitude:  12.01,
			Price:          100000,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		routedetail := entities.RouteDetail{
			Id:            uuid.New(),
			DestinationId: destination.Id,
			RouteId:       route.Id,
			Longitude:     destination.Longitude,
			Latitude:      destination.Latitude,
			Duration:      60,
			Order:         1,
			VisitStart:    []byte("08:00"),
			VisitEnd:      []byte("09:00"),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		routes = append(routes, route)
		routeDetails = append(routeDetails, routedetail)
	}

	for _, route := range routes {
		if err = s.db.Where(entities.Route{Id: route.Id}).
			FirstOrCreate(&route).Error; err != nil {
			log.Fatalf("failed to create route: %v", err)
		}
	}

	for _, routeDetail := range routeDetails {
		if err = s.db.Where(entities.RouteDetail{Id: routeDetail.Id}).
			FirstOrCreate(&routeDetail).Error; err != nil {
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
