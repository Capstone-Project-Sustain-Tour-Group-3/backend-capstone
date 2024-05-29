package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationAddresses() {
	destinations, _ := getAllDestinations(s)

	destinationAddresses := []entities.DestinationAddress{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac71")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac72")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac73")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac74")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac75")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac76")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac77")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac78")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac79")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac70")},
	}

	for idx, destination := range destinations {
		province, _ := getRandomProvince(s)
		destinationAddress := destinationAddresses[idx]
		destinationAddress.DestinationId = destination.Id
		destinationAddress.ProvinceId = province.Id

		if err := s.db.Where(entities.DestinationAddress{Id: destinationAddress.Id}).
			FirstOrCreate(&destinationAddress).Error; err != nil {
			log.Fatalf("failed to create destination address: %v", err)
		}
	}
}
