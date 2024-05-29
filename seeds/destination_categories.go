package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinationCategories() {
	destinations, _ := getAllDestinations(s)

	destinationCategories := []entities.DestinationCategory{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac41")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac42")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac43")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac44")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac45")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac46")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac47")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac48")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac49")},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac10")},
	}

	for idx, destination := range destinations {
		category, _ := getRandomCategory(s)
		destinationCategory := destinationCategories[idx]
		destinationCategory.DestinationId = destination.Id
		destinationCategory.CategoryId = category.Id

		if err := s.db.Where(entities.DestinationCategory{Id: destinationCategory.Id}).
			FirstOrCreate(&destinationCategory).Error; err != nil {
			log.Fatalf("failed to create destination category: %v", err)
		}
	}
}
