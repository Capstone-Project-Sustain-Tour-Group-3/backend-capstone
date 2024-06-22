package seeds

import (
	"encoding/csv"
	"log"
	"os"

	"capstone/entities"
)

func (s Seed) SeedProvinces() {
	file, err := os.Open("storage/data_url_provinces.csv")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	reader := csv.NewReader(file)
	// Skip header row
	if _, err = reader.Read(); err != nil {
		log.Fatalf("failed to read header row: %v", err)
	}
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV file: %v", err)
	}

	provinces := []entities.Province{
		{
			Id:   "11",
			Name: records[0][0],
			Url:  records[0][1],
		},
		{
			Id:   "12",
			Name: records[1][0],
			Url:  records[1][1],
		},
		{
			Id:   "13",
			Name: records[2][0],
			Url:  records[2][1],
		},
		{
			Id:   "14",
			Name: records[3][0],
			Url:  records[3][1],
		},
		{
			Id:   "21",
			Name: records[4][0],
			Url:  records[4][1],
		},
		{
			Id:   "15",
			Name: records[5][0],
			Url:  records[5][1],
		},
		{
			Id:   "16",
			Name: records[6][0],
			Url:  records[6][1],
		},
		{
			Id:   "19",
			Name: records[7][0],
			Url:  records[7][1],
		},
		{
			Id:   "17",
			Name: records[8][0],
			Url:  records[8][1],
		},
		{
			Id:   "18",
			Name: records[9][0],
			Url:  records[9][1],
		},
		{
			Id:   "31",
			Name: records[10][0],
			Url:  records[10][1],
		},
		{
			Id:   "32",
			Name: records[11][0],
			Url:  records[11][1],
		},
		{
			Id:   "33",
			Name: records[12][0],
			Url:  records[12][1],
		},
		{
			Id:   "34",
			Name: records[13][0],
			Url:  records[13][1],
		},
		{
			Id:   "35",
			Name: records[14][0],
			Url:  records[14][1],
		},
		{
			Id:   "36",
			Name: records[15][0],
			Url:  records[15][1],
		},
		{
			Id:   "51",
			Name: records[16][0],
			Url:  records[16][1],
		},
		{
			Id:   "52",
			Name: records[17][0],
			Url:  records[17][1],
		},
		{
			Id:   "53",
			Name: records[18][0],
			Url:  records[18][1],
		},
		{
			Id:   "61",
			Name: records[19][0],
			Url:  records[19][1],
		},
		{
			Id:   "62",
			Name: records[20][0],
			Url:  records[20][1],
		},
		{
			Id:   "63",
			Name: records[21][0],
			Url:  records[21][1],
		},
		{
			Id:   "64",
			Name: records[22][0],
			Url:  records[22][1],
		},
		{
			Id:   "65",
			Name: records[23][0],
			Url:  records[23][1],
		},
		{
			Id:   "71",
			Name: records[24][0],
			Url:  records[24][1],
		},
		{
			Id:   "72",
			Name: records[25][0],
			Url:  records[25][1],
		},
		{
			Id:   "73",
			Name: records[26][0],
			Url:  records[26][1],
		},
		{
			Id:   "74",
			Name: records[27][0],
			Url:  records[27][1],
		},
		{
			Id:   "75",
			Name: records[28][0],
			Url:  records[28][1],
		},
		{
			Id:   "76",
			Name: records[29][0],
			Url:  records[29][1],
		},
		{
			Id:   "81",
			Name: records[30][0],
			Url:  records[30][1],
		},
		{
			Id:   "82",
			Name: records[31][0],
			Url:  records[31][1],
		},
		{
			Id:   "91",
			Name: records[32][0],
			Url:  records[32][1],
		},
		{
			Id:   "92",
			Name: records[33][0],
			Url:  records[33][1],
		},
		{
			Id:   "93",
			Name: records[34][0],
			Url:  records[34][1],
		},
		{
			Id:   "94",
			Name: records[35][0],
			Url:  records[35][1],
		},
		{
			Id:   "95",
			Name: records[36][0],
			Url:  records[36][1],
		},
		{
			Id:   "96",
			Name: records[37][0],
			Url:  records[37][1],
		},
	}

	for _, province := range provinces {
		if err = s.db.Where(entities.Province{Name: province.Name}).
			FirstOrCreate(&province).Error; err != nil {
			log.Fatalf("failed to create province: %v", err)
		}
	}
}

// func getRandomProvince(s Seed) (*entities.Province, error) {
//	var province entities.Province
//
//	if err := s.db.Order("RAND()").First(&province).Error; err != nil {
//		return nil, err
//	}
//
//	return &province, nil
//}
