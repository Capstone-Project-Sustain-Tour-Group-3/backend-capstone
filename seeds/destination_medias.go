package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationMedias() {
	destinationMedias := []entities.DestinationMedia{
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e971"), Type: "jpg", Url: "https://picsum.photos/300/200", Title: "Destinasi 1"},                                                                        //nolint:lll
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e972"), Type: "jpg", Url: "https://picsum.photos/300/200", Title: "Destinasi 2"},                                                                        //nolint:lll
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e973"), Type: "jpg", Url: "https://picsum.photos/300/200", Title: "Destinasi 3"},                                                                        //nolint:lll
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e974"), Type: "mp4", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.mp4", Title: "Tears of Steel"},                 //nolint:lll
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e975"), Type: "mp4", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/VolkswagenGTIReview.mp4", Title: "Volkswagen GTI Review"},   //nolint:lll
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e976"), Type: "mp4", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/WeAreGoingOnBullrun.mp4", Title: "We Are Going On Bullrun"}, //nolint:lll
	}

	for _, destinationMedia := range destinationMedias {
		// get random destination
		destination, _ := getRandomDestination(s)
		destinationMedia.DestinationId = destination.Id
		if err := s.db.Where(entities.DestinationMedia{Title: destinationMedia.Title}).
			FirstOrCreate(&destinationMedia).Error; err != nil {
			log.Fatalf("failed to create destination media: %v", err)
		}
	}
}
