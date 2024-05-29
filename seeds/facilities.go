package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedFacilities() {
	facilities := []entities.Facility{
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd1"), Name: "Kamar Mandi", Url: "https://picsum.photos/300/200"}, //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd2"), Name: "Mushola", Url: "https://picsum.photos/300/200"},
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd3"), Name: "Area Parkir", Url: "https://picsum.photos/300/200"}, //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd4"), Name: "Penginapan", Url: "https://picsum.photos/300/200"},  //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd5"), Name: "Restoran", Url: "https://picsum.photos/300/200"},
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd6"), Name: "Toko Oleh-Oleh", Url: "https://picsum.photos/300/200"},          //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd7"), Name: "Pusat Informasi", Url: "https://picsum.photos/300/200"},         //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd8"), Name: "Area Bermain Anak", Url: "https://picsum.photos/300/200"},       //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efd9"), Name: "Jalur Hiking", Url: "https://picsum.photos/300/200"},            //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efa1"), Name: "Kolam Renang", Url: "https://picsum.photos/300/200"},            //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efa2"), Name: "WiFi Gratis", Url: "https://picsum.photos/300/200"},             //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efa3"), Name: "Pusat Kesehatan", Url: "https://picsum.photos/300/200"},         //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efa4"), Name: "Layanan Penjemputan", Url: "https://picsum.photos/300/200"},     //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efa5"), Name: "Tempat Penitipan Barang", Url: "https://picsum.photos/300/200"}, //nolint:lll
		{Id: uuid.MustParse("515cb2da-3361-48b8-99fd-bd894828efa6"), Name: "Area Piknik", Url: "https://picsum.photos/300/200"},             //nolint:lll
	}

	for _, facility := range facilities {
		if err := s.db.Where(entities.Facility{Name: facility.Name}).
			FirstOrCreate(&facility).Error; err != nil {
			log.Fatalf("failed to create facility: %v", err)
		}
	}
}

func getRandomFacility(s Seed) (*entities.Facility, error) {
	var facility entities.Facility

	if err := s.db.Order("RAND()").First(&facility).Error; err != nil {
		return nil, err
	}

	return &facility, nil
}
