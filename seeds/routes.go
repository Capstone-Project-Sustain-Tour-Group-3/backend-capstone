package seeds

import (
	"log"
	"strconv"
	"time"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedRoutes() {
	destination, err := getRandomDestination(s)
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

	routeIds := [15]uuid.UUID{
		uuid.MustParse("179665db-dfba-4d10-8371-29d2ce169cc2"),
		uuid.MustParse("ef234ca4-f0d4-4ee1-a7ae-1199f72a2d7f"),
		uuid.MustParse("3fbff38a-5e74-48a7-bad1-65fc58035436"),
		uuid.MustParse("da386bd6-e1b7-4129-b3a6-e75d24594c1a"),
		uuid.MustParse("2d128871-2825-4e44-90dc-92567df931f3"),
		uuid.MustParse("e32f9896-3cd7-423b-8f36-0df92f4dcde4"),
		uuid.MustParse("cdf5c8ef-0401-4e31-926a-c102b88ea1f9"),
		uuid.MustParse("ed23be94-c0c5-40bf-a017-2ee831ab919e"),
		uuid.MustParse("f1d60ba6-e689-4041-a11c-1a5a8cf0c793"),
		uuid.MustParse("2c1fe7b4-9074-4729-afe7-4b474d6b274d"),
		uuid.MustParse("e55eac90-521d-4cc5-879a-d208870f0275"),
		uuid.MustParse("12d9132d-487f-456c-8903-583969dc6a3b"),
		uuid.MustParse("0308d6d2-d28b-42d0-ad28-e64fd96bb2fd"),
		uuid.MustParse("3e9009a5-8edd-421f-9c2c-0740bfde93f8"),
		uuid.MustParse("f8fad44d-e5b1-4a3b-939e-1955bf9f7c67"),
	}

	routeDetailIds := [15]uuid.UUID{
		uuid.MustParse("92b83aa1-a21c-4aae-9c05-db37202bc473"),
		uuid.MustParse("57ba59a7-f151-4a4f-b1ea-60ba9c464944"),
		uuid.MustParse("d4f67fdc-0f83-412a-8a3d-29d806de3b20"),
		uuid.MustParse("b5633440-68b7-4384-acf6-8b0e6dd48164"),
		uuid.MustParse("4706fb68-cd5c-4254-9d42-a29376a8b8c0"),
		uuid.MustParse("0fd4b579-711a-4811-a8ec-0fcefe43504e"),
		uuid.MustParse("f4807627-2a0f-4510-b6be-99d9765db504"),
		uuid.MustParse("e95d3985-3508-4405-a6d7-9e9d5106b0a6"),
		uuid.MustParse("d37a5542-0300-4027-b006-4e28d04f4205"),
		uuid.MustParse("457ad80a-f4ff-48cd-9dfb-d8220612ec33"),
		uuid.MustParse("3edc771b-a2fc-485d-9462-501b6dad15c3"),
		uuid.MustParse("dbfe7fcd-b57f-4a03-b448-f875b0f300b9"),
		uuid.MustParse("9bc4a709-4806-46d1-a994-31cc5615e56c"),
		uuid.MustParse("43025319-7917-4bd9-8a6a-bf94fbccb2f0"),
		uuid.MustParse("f4f9cfed-0a58-4cf9-b29d-fcdf9716a126"),
	}

	for i := range [15]int{} {
		route := entities.Route{
			Id:             routeIds[i],
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
			Id:            routeDetailIds[i],
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
