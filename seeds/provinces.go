package seeds

import (
	"log"

	"capstone/entities"
)

func (s Seed) SeedProvinces() {
	provinces := []entities.Province{
		{
			Id:   "11",
			Name: "ACEH",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "12",
			Name: "SUMATERA UTARA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "13",
			Name: "SUMATERA BARAT",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "14",
			Name: "RIAU",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "15",
			Name: "JAMBI",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "16",
			Name: "SUMATERA SELATAN",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "17",
			Name: "BENGKULU",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "18",
			Name: "LAMPUNG",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "19",
			Name: "KEPULAUAN BANGKA BELITUNG",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "21",
			Name: "KEPULAUAN RIAU",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "31",
			Name: "DKI JAKARTA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "32",
			Name: "JAWA BARAT",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "33",
			Name: "JAWA TENGAH",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "34",
			Name: "DAERAH ISTIMEWA YOGYAKARTA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "35",
			Name: "JAWA TIMUR",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "36",
			Name: "BANTEN",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "51",
			Name: "BALI",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "52",
			Name: "NUSA TENGGARA BARAT",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "53",
			Name: "NUSA TENGGARA TIMUR",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "61",
			Name: "KALIMANTAN BARAT",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "62",
			Name: "KALIMANTAN TENGAH",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "63",
			Name: "KALIMANTAN SELATAN",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "64",
			Name: "KALIMANTAN TIMUR",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "65",
			Name: "KALIMANTAN UTARA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "71",
			Name: "SULAWESI UTARA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "72",
			Name: "SULAWESI TENGAH",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "73",
			Name: "SULAWESI SELATAN",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "74",
			Name: "SULAWESI TENGGARA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "75",
			Name: "GORONTALO",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "76",
			Name: "SULAWESI BARAT",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "81",
			Name: "MALUKU",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "82",
			Name: "MALUKU UTARA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "91",
			Name: "PAPUA",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "92",
			Name: "PAPUA BARAT",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "93",
			Name: "PAPUA SELATAN",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "94",
			Name: "PAPUA TENGAH",
			Url:  "https://picsum.photos/250/250",
		},
		{
			Id:   "95",
			Name: "PAPUA PEGUNUNGAN",
			Url:  "https://picsum.photos/250/250",
		},
	}

	for _, province := range provinces {
		if err := s.db.Where(entities.Province{Name: province.Name}).
			FirstOrCreate(&province).Error; err != nil {
			log.Fatalf("failed to create province: %v", err)
		}
	}
}

func getRandomProvince(s Seed) (*entities.Province, error) {
	var province entities.Province

	if err := s.db.Order("RAND()").First(&province).Error; err != nil {
		return nil, err
	}

	return &province, nil
}
