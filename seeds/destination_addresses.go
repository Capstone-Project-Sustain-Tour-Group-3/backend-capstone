package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationAddresses() {
	destinations, _ := getAllDestinations(s)

	destinationAddresses := []entities.DestinationAddress{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac71"), PostalCode: "10110", StreetName: "Jalan MH Thamrin"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac72"), PostalCode: "40181", StreetName: "Jalan Dago"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac73"), PostalCode: "60231", StreetName: "Jalan Tunjungan"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac74"), PostalCode: "55281", StreetName: "Jalan Malioboro"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac75"), PostalCode: "80228", StreetName: "Jalan Sunset Road"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac76"), PostalCode: "20151", StreetName: "Jalan Gatot Subroto"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac77"), PostalCode: "50149", StreetName: "Jalan Pandanaran"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac78"), PostalCode: "90114", StreetName: "Jalan Pettarani"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac79"), PostalCode: "76125", StreetName: "Jalan Sudirman"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac70"), PostalCode: "95111", StreetName: "Jalan Sam Ratulangi"},
	}

	for idx, destination := range destinations {
		subdistrict, _ := getRandomSubdistrict(s)

		destinationAddress := destinationAddresses[idx]
		destinationAddress.DestinationId = destination.Id
		destinationAddress.SubdistrictId = subdistrict.Id
		destinationAddress.CityId = subdistrict.City.Id
		destinationAddress.ProvinceId = subdistrict.City.ProvinceId

		if err := s.db.Where(entities.DestinationAddress{Id: destinationAddress.Id}).
			FirstOrCreate(&destinationAddress).Error; err != nil {
			log.Fatalf("failed to create destination address: %v", err)
		}
	}
}
