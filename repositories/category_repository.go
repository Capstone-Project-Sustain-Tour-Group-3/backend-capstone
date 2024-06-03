package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICategoriesRepository interface {
	Create(category *entities.Category) error
	FindAll() ([]entities.Category, error)
	FindById(id uuid.UUID) (*entities.Category, error)
	Update(category *entities.Category) error
	Delete(category *entities.Category) error
}

type CategoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{db}
}

func (r *CategoriesRepository) Create(category *entities.Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriesRepository) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoriesRepository) FindById(id uuid.UUID) (*entities.Category, error) {
	var category *entities.Category
	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoriesRepository) Update(category *entities.Category) error {
	if err := r.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriesRepository) Delete(category *entities.Category) error {
	if err := r.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
