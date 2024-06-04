package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICityRepository interface {
	Create(city *entities.City) error
	FindAll() ([]entities.City, error)
	FindById(id uuid.UUID) (*entities.City, error)
	Update(city *entities.City) error
	Delete(city *entities.City) error
}

type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db}
}

func (r *CityRepository) Create(city *entities.City) error {
	if err := r.db.Create(city).Error; err != nil {
		return err
	}
	return nil
}

func (r *CityRepository) FindAll() ([]entities.City, error) {
	var cities []entities.City
	if err := r.db.Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *CityRepository) FindById(id uuid.UUID) (*entities.City, error) {
	var city *entities.City
	if err := r.db.Where("id = ?", id).First(&city).Error; err != nil {
		return nil, err
	}
	return city, nil
}

func (r *CityRepository) Update(city *entities.City) error {
	if err := r.db.Save(city).Error; err != nil {
		return err
	}
	return nil
}

func (r *CityRepository) Delete(city *entities.City) error {
	if err := r.db.Delete(city).Error; err != nil {
		return err
	}
	return nil
}
