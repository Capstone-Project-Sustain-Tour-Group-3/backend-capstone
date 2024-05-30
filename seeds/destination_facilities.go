package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationFacilities() {
	destinations, _ := getAllDestinations(s)

	destinationFacilities := []entities.DestinationFacility{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac81")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac82")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac83")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac84")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac85")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac86")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac87")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac88")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac89")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac80")},
	}

	for idx, destination := range destinations {
		facility, _ := getRandomFacility(s)
		destinationFacility := destinationFacilities[idx]
		destinationFacility.DestinationId = destination.Id
		destinationFacility.FacilityId = facility.Id

		if err := s.db.Where(entities.DestinationFacility{Id: destinationFacility.Id}).
			FirstOrCreate(&destinationFacility).Error; err != nil {
			log.Fatalf("failed to create destination facility: %v", err)
		}
	}
}
