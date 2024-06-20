package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type ICityRepository interface {
	Create(city *entities.City) error
	FindAll() ([]entities.City, error)
	FindById(id string) (*entities.City, error)
	Update(city *entities.City) error
	Delete(city *entities.City) error
	GetCitiesWithDestinations() ([]entities.City, error)
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

func (r *CityRepository) FindById(id string) (*entities.City, error) {
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

func (r *CityRepository) GetCitiesWithDestinations() ([]entities.City, error) {
	var cities []entities.City

	err := r.db.Preload("Province").
		Joins("JOIN destination_addresses da ON da.city_id = cities.id").
		Joins("JOIN destinations d ON d.id = da.destination_id").
		Where("d.id IS NOT NULL").
		Group("cities.id").
		Order("cities.name ASC").
		Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}
