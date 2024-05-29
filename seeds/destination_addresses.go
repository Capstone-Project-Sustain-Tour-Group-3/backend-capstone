package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationAddresses() {
	destinations, _ := getAllDestinations(s)

	destinationAddresses := []entities.DestinationAddress{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac71"), City: "Jakarta", Regency: "Jakarta Pusat", PostalCode: "10110", StreetName: "Jalan MH Thamrin"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac72"), City: "Bandung", Regency: "Bandung Barat", PostalCode: "40181", StreetName: "Jalan Dago"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac73"), City: "Surabaya", Regency: "Surabaya Timur", PostalCode: "60231", StreetName: "Jalan Tunjungan"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac74"), City: "Yogyakarta", Regency: "Yogyakarta Utara", PostalCode: "55281", StreetName: "Jalan Malioboro"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac75"), City: "Denpasar", Regency: "Denpasar Selatan", PostalCode: "80228", StreetName: "Jalan Sunset Road"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac76"), City: "Medan", Regency: "Medan Kota", PostalCode: "20151", StreetName: "Jalan Gatot Subroto"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac77"), City: "Semarang", Regency: "Semarang Barat", PostalCode: "50149", StreetName: "Jalan Pandanaran"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac78"), City: "Makassar", Regency: "Makassar Timur", PostalCode: "90114", StreetName: "Jalan Pettarani"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac79"), City: "Balikpapan", Regency: "Balikpapan Utara", PostalCode: "76125", StreetName: "Jalan Sudirman"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac70"), City: "Manado", Regency: "Manado Selatan", PostalCode: "95111", StreetName: "Jalan Sam Ratulangi"},
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
