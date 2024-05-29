package seeds

import (
	"log"

	"github.com/google/uuid"

	"capstone/entities"
)

func (s Seed) SeedProvinces() {
	provinces := []entities.Province{
		{Id: uuid.MustParse("c67a8a7e-999b-4b5a-bd71-f1f1b68c58f1"), Name: "Aceh", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("e5d2b1b8-8b24-4c2d-9bb5-c1f1d1d1e1e1"), Name: "Sumatera Utara", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("f6c3d2a3-1b4b-4c3d-9c1f-d2f1b1c2e2e2"), Name: "Sumatera Barat", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("d7b4c3e2-2c3d-4d5a-8d2c-e3c1f2f1d1d1"), Name: "Riau", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("e8a5d4b1-3c4d-5e6a-9c3d-f1b2d2c3e3e3"), Name: "Kepulauan Riau", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("f9b6c5a2-4d5e-6f7a-8d4e-c2e1d1c2f2f2"), Name: "Jambi", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be181"), Name: "Sumatera Selatan", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be182"), Name: "Bangka Belitung", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be183"), Name: "Bengkulu", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be184"), Name: "Lampung", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be185"), Name: "DKI Jakarta", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be186"), Name: "Jawa Barat", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be187"), Name: "Jawa Tengah", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be188"), Name: "DI Yogyakarta", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be189"), Name: "Jawa Timur", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be111"), Name: "Banten", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be121"), Name: "Bali", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be131"), Name: "Nusa Tenggara Barat", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be141"), Name: "Nusa Tenggara Timur", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be151"), Name: "Kalimantan Barat", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be161"), Name: "Kalimantan Tengah", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be171"), Name: "Kalimantan Selatan", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be191"), Name: "Kalimantan Timur", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be281"), Name: "Kalimantan Utara", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be381"), Name: "Sulawesi Utara", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be481"), Name: "Sulawesi Tengah", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be581"), Name: "Sulawesi Selatan", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be681"), Name: "Sulawesi Tenggara", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be781"), Name: "Gorontalo", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be881"), Name: "Sulawesi Barat", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be981"), Name: "Maluku", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be211"), Name: "Maluku Utara", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be221"), Name: "Papua", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be231"), Name: "Papua Barat", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be241"), Name: "Papua Tengah", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be351"), Name: "Papua Pegunungan", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be561"), Name: "Papua Selatan", Url: "https://picsum.photos/250/250"},
		{Id: uuid.MustParse("2546cc42-90ad-4838-9216-0bcfce0be689"), Name: "Papua Barat Daya", Url: "https://picsum.photos/250/250"},
	}

	for _, province := range provinces {
		if err := s.db.Where(entities.Province{Name: province.Name}).
			FirstOrCreate(&province).Error; err != nil {
			log.Fatalf("failed to create province: %v", err)
		}
	}
}
