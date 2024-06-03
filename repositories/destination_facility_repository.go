package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDestinationCategoryRepository interface {
	Create(destinationCategory *entities.DestinationCategory) error
	FindAll() ([]entities.DestinationCategory, error)
	FindById(id uuid.UUID) (*entities.DestinationCategory, error)
	Update(destinationCategory *entities.DestinationCategory) error
	Delete(destinationCategory *entities.DestinationCategory) error
}

type DestinationCategoryRepository struct {
	db *gorm.DB
}

func NewDestinationCategoryRepository(db *gorm.DB) *DestinationCategoryRepository {
	return &DestinationCategoryRepository{db}
}

func (r *DestinationCategoryRepository) Create(destinationCategory *entities.DestinationCategory) error {
	if err := r.db.Create(destinationCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationCategoryRepository) FindAll() ([]entities.DestinationCategory, error) {
	var destinationCategories []entities.DestinationCategory
	if err := r.db.Find(&destinationCategories).Error; err != nil {
		return nil, err
	}
	return destinationCategories, nil
}

func (r *DestinationCategoryRepository) FindById(id uuid.UUID) (*entities.DestinationCategory, error) {
	var destinationCategory *entities.DestinationCategory
	if err := r.db.Where("id = ?", id).First(&destinationCategory).Error; err != nil {
		return nil, err
	}
	return destinationCategory, nil
}

func (r *DestinationCategoryRepository) Update(destinationCategory *entities.DestinationCategory) error {
	if err := r.db.Save(destinationCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationCategoryRepository) Delete(destinationCategory *entities.DestinationCategory) error {
	if err := r.db.Delete(destinationCategory).Error; err != nil {
		return err
	}
	return nil
}
