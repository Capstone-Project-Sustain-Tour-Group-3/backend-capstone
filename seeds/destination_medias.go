package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationMedias() {
	destinationImages := []entities.DestinationMedia{
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e941"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 1"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e942"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 2"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e943"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 3"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e944"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 4"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e945"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 5"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e946"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 6"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e947"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 7"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e948"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 8"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e949"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 9"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e940"), Type: "image", Url: "https://picsum.photos/300/200", Title: "Destinasi 10"},
	}

	destinationVideos := []entities.DestinationMedia{
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e931"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.video", Title: "Tears of Steel 1"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e932"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/VolkswagenGTIReview.video", Title: "Volkswagen GTI Review 2"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e933"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/WeAreGoingOnBullrun.video", Title: "We Are Going On Bullrun 3"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e934"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.video", Title: "Tears of Steel 4"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e935"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/VolkswagenGTIReview.video", Title: "Volkswagen GTI Review 5"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e936"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/WeAreGoingOnBullrun.video", Title: "We Are Going On Bullrun 6"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e937"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.video", Title: "Tears of Steel 7"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e938"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/VolkswagenGTIReview.video", Title: "Volkswagen GTI Review 8"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e939"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/WeAreGoingOnBullrun.video", Title: "We Are Going On Bullrun 9"},
		{Id: uuid.MustParse("1bde58e3-ef19-4daa-9df7-084ba5d3e930"), Type: "video", Url: "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.video", Title: "Tears of Steel 10"},
	}

	destinations, _ := getAllDestinations(s)

	for idx, destination := range destinations {
		destinationImage := destinationImages[idx]
		destinationImage.DestinationId = destination.Id

		if err := s.db.Where(entities.DestinationMedia{Id: destinationImage.Id}).
			FirstOrCreate(&destinationImage).Error; err != nil {
			log.Fatalf("failed to create destination image: %v", err)
		}
	}

	for idx, destination := range destinations {
		destinationVideo := destinationVideos[idx]
		destinationVideo.DestinationId = destination.Id

		if err := s.db.Where(entities.DestinationMedia{Id: destinationVideo.Id}).
			FirstOrCreate(&destinationVideo).Error; err != nil {
			log.Fatalf("failed to create destination image: %v", err)
		}
	}
}
