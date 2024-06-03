package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinations() {
	openTime := "08:00"
	closeTime := "17:00"

	destinations := []entities.Destination{
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec6"), Name: "Goa Pindul", Description: "Gua tempat Joko terbentur tersebut dinamai Gua Pindul yang berasal dari kata dalam bahasa Jawa pipi gebendul yang berarti pipi yang terbentur.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 25000, Longitude: 123.456, Latitude: 456.789, VisitCount: 0}, //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec1"), Name: "Candi Borobudur", Description: "Candi Borobudur adalah sebuah candi Buddha yang terletak di Magelang, Jawa Tengah.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 50000, Longitude: 110.2038, Latitude: -7.6079, VisitCount: 0},                                                       //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec2"), Name: "Pantai Kuta", Description: "Pantai Kuta adalah salah satu pantai yang terkenal di Bali, Indonesia.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 0, Longitude: 115.1675, Latitude: -8.7174, VisitCount: 0},                                                                           //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec3"), Name: "Danau Toba", Description: "Danau Toba adalah danau terbesar di Indonesia dan Asia Tenggara.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 0, Longitude: 99.0852, Latitude: 2.6696, VisitCount: 0},                                                                                    //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec4"), Name: "Taman Mini Indonesia Indah", Description: "Taman Mini Indonesia Indah adalah sebuah taman wisata bertema budaya Indonesia yang terletak di Jakarta Timur.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 20000, Longitude: 106.8956, Latitude: -6.3027, VisitCount: 0},                //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec5"), Name: "Gunung Bromo", Description: "Gunung Bromo adalah sebuah gunung berapi aktif yang terletak di Jawa Timur, Indonesia.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 30000, Longitude: 112.9528, Latitude: -7.9425, VisitCount: 0},                                                      //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec7"), Name: "Pulau Komodo", Description: "Pulau Komodo adalah sebuah pulau yang terletak di Kepulauan Nusa Tenggara, Indonesia.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 150000, Longitude: 119.4986, Latitude: -8.5833, VisitCount: 0},                                                      //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec8"), Name: "Raja Ampat", Description: "Raja Ampat adalah kepulauan yang terletak di bagian barat Papua, Indonesia.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 500000, Longitude: 130.5036, Latitude: -1.0562, VisitCount: 0},                                                                  //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adec9"), Name: "Tanah Lot", Description: "Tanah Lot adalah sebuah formasi batuan di lepas pantai pulau Bali, Indonesia.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 20000, Longitude: 115.0865, Latitude: -8.6211, VisitCount: 0},                                                                  //nolint:lll
		{Id: uuid.MustParse("306d305e-3359-4884-8d38-89c04e8adea6"), Name: "Kawah Ijen", Description: "Kawah Ijen adalah sebuah kompleks gunung berapi di Banyuwangi, Jawa Timur, Indonesia.", OpenTime: openTime, CloseTime: closeTime, EntryPrice: 100000, Longitude: 114.2423, Latitude: -8.0582, VisitCount: 0},                                                        //nolint:lll
	}

	for _, destination := range destinations {
		category, err := getRandomCategory(s)
		if err != nil {
			log.Fatalf("failed to get random category: %v", err)
		}

		destination.CategoryId = category.Id
		if err := s.db.Where(entities.Destination{Name: destination.Name}).FirstOrCreate(&destination).Error; err != nil {
			log.Fatalf("failed to create destination: %v", err)
		}
	}
}

func getAllDestinations(s Seed) ([]entities.Destination, error) {
	var destinations []entities.Destination

	if err := s.db.Find(&destinations).Error; err != nil {
		return nil, err
	}

	return destinations, nil
}
