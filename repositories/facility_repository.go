package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	Create(category *entities.Category) error
	FindAll() ([]entities.Category, error)
	FindById(id uuid.UUID) (*entities.Category, error)
	Update(category *entities.Category) error
	Delete(category *entities.Category) error
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(category *entities.Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) FindById(id uuid.UUID) (*entities.Category, error) {
	var category *entities.Category
	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) Update(category *entities.Category) error {
	if err := r.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) Delete(category *entities.Category) error {
	if err := r.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
