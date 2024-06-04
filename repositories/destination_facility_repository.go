package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDestinationFacilityRepository interface {
	Create(destinationFacilities *[]entities.DestinationFacility) error
	FindAll() ([]entities.DestinationFacility, error)
	FindById(id uuid.UUID) (*entities.DestinationFacility, error)
	Update(destinationFacility *entities.DestinationFacility) error
	Delete(destinationFacility *entities.DestinationFacility) error
}

type DestinationFacilityRepository struct {
	db *gorm.DB
}

func NewDestinationFacilityRepository(db *gorm.DB) *DestinationFacilityRepository {
	return &DestinationFacilityRepository{db}
}

func (r *DestinationFacilityRepository) Create(destinationFacilities *[]entities.DestinationFacility) error {
	if err := r.db.Create(destinationFacilities).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationFacilityRepository) FindAll() ([]entities.DestinationFacility, error) {
	var destinationFacilities []entities.DestinationFacility
	if err := r.db.Find(&destinationFacilities).Error; err != nil {
		return nil, err
	}
	return destinationFacilities, nil
}

func (r *DestinationFacilityRepository) FindById(id uuid.UUID) (*entities.DestinationFacility, error) {
	var destinationFacility *entities.DestinationFacility
	if err := r.db.Where("id = ?", id).First(&destinationFacility).Error; err != nil {
		return nil, err
	}
	return destinationFacility, nil
}

func (r *DestinationFacilityRepository) Update(destinationFacility *entities.DestinationFacility) error {
	if err := r.db.Save(destinationFacility).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationFacilityRepository) Delete(destinationFacility *entities.DestinationFacility) error {
	if err := r.db.Delete(destinationFacility).Error; err != nil {
		return err
	}
	return nil
}
