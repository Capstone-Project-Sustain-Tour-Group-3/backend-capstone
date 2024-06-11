package seeds

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/google/uuid"

	"capstone/entities"
)

func (s Seed) SeedFacilities() {
	file, err := os.Open("storage/data_url_facilities.csv")
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

	var facilities []entities.Facility
	for _, record := range records {
		if err != nil {
			log.Fatalf("failed to parse UUID: %v", err)
		}
		facilities = append(facilities, entities.Facility{
			Id:   uuid.New(),
			Name: record[0],
			Url:  record[1],
		})
	}

	for _, facility := range facilities {
		var temp entities.Facility
		if err = s.db.FirstOrInit(&temp, entities.Facility{Name: facility.Name}).Error; err != nil {
			log.Fatalf("failed to init facility: %v", err)
		}

		if temp.Id == uuid.Nil {
			if err = s.db.Create(&facility).Error; err != nil {
				log.Fatalf("failed to create facility: %v", err)
			}
		}
	}
}
