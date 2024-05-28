package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedCategories() {
	categories := []entities.Category{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac41"), Name: "Alam", Url: "https://picsum.photos/300/200"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac42"), Name: "Seni dan Budaya", Url: "https://picsum.photos/300/200"}, //nolint:lll
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac43"), Name: "Sejarah", Url: "https://picsum.photos/300/200"},
	}

	for _, category := range categories {
		if err := s.db.Where(entities.Category{Name: category.Name}).
			FirstOrCreate(&category).Error; err != nil {
			log.Fatalf("failed to create category: %v", err)
		}
	}
}

func getRandomCategory(s Seed) (*entities.Category, error) {
	var category entities.Category

	if err := s.db.Order("RAND()").First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
