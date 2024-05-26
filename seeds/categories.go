package seeds

import (
	"errors"
	"fmt"
	"log"

	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s Seed) SeedCategories() {
	categories := []entities.Category{
		{Id: uuid.New(), Name: "Alam", Url: "https://picsum.photos/300/200"},
		{Id: uuid.New(), Name: "Seni dan Budaya", Url: "https://picsum.photos/300/200"},
		{Id: uuid.New(), Name: "Sejarah", Url: "https://picsum.photos/300/200"},
	}

	for _, category := range categories {
		var existingCategory entities.Category
		if err := s.db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err = s.db.Create(&category).Error; err != nil {
					log.Fatalf("failed to create category: %v", err)
				}
				fmt.Printf("Category %s created successfully.\n", category.Name)
			} else {
				log.Fatalf("failed to check if category exists: %v", err)
			}
		}
	}

	fmt.Println("[+] Categories seeding succeeded")
}
