package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedCategories() {
	categories := []entities.Category{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac61"), Name: "Alam", Url: "https://storage.googleapis.com/categories_url/alam.jpeg"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac62"), Name: "Seni dan Budaya", Url: "https://storage.googleapis.com/categories_url/seni-budaya.jpeg"}, //nolint:lll
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac63"), Name: "Sejarah", Url: "https://storage.googleapis.com/categories_url/sejarah.jpeg"},
	}

	for _, category := range categories {
		if err := s.db.Where(entities.Category{Name: category.Name}).
			FirstOrCreate(&category).Error; err != nil {
			log.Fatalf("failed to create category: %v", err)
		}
	}
}
