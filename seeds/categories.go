package seeds

import (
	"log"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedCategories() {
	categories := []entities.Category{
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac61"), Name: "Alam", Url: "https://res.cloudinary.com/alta-minpro/image/upload/v1717500654/nqecwzfmydqvdznqr3iq.jpg"},
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac62"), Name: "Seni dan Budaya", Url: "https://res.cloudinary.com/alta-minpro/image/upload/v1717500659/hyvolbf45qcc04snsdlx.jpg"}, //nolint:lll
		{Id: uuid.MustParse("eb77b590-b255-4ea1-b11a-d445a259ac63"), Name: "Sejarah", Url: "https://res.cloudinary.com/alta-minpro/image/upload/v1717500656/s8demedae0h85rcdpqfw.jpg"},
	}

	for _, category := range categories {
		if err := s.db.Where(entities.Category{Name: category.Name}).
			FirstOrCreate(&category).Error; err != nil {
			log.Fatalf("failed to create category: %v", err)
		}
	}
}
