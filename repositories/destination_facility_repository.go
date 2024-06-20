package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDestinationFacilityRepository interface {
	Create(destinationFacilities *[]entities.DestinationFacility, tx *gorm.DB) error
	FindAll() ([]entities.DestinationFacility, error)
	FindById(id uuid.UUID) (*entities.DestinationFacility, error)
	Update(destinationFacility *entities.DestinationFacility, tx *gorm.DB) error
	Delete(destinationFacility *entities.DestinationFacility) error
	DeleteMany(destinationFacilities *[]entities.DestinationFacility, tx *gorm.DB) error
}

type DestinationFacilityRepository struct {
	db *gorm.DB
}

func NewDestinationFacilityRepository(db *gorm.DB) *DestinationFacilityRepository {
	return &DestinationFacilityRepository{db}
}

func (r *DestinationFacilityRepository) Create(destinationFacilities *[]entities.DestinationFacility, tx *gorm.DB) error {
	if err := tx.Create(destinationFacilities).Error; err != nil {
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

func (r *DestinationFacilityRepository) Update(destinationFacility *entities.DestinationFacility, tx *gorm.DB) error {
	if err := tx.Save(destinationFacility).Error; err != nil {
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

func (r *DestinationFacilityRepository) DeleteMany(destinationFacilities *[]entities.DestinationFacility, tx *gorm.DB) error {
	if len(*destinationFacilities) == 0 {
		return nil
	}

	var ids []uuid.UUID
	for _, destinationFacility := range *destinationFacilities {
		ids = append(ids, destinationFacility.Id)
	}

	if err := tx.Where("id IN ?", ids).Delete(&entities.DestinationFacility{}).Error; err != nil {
		return err
	}

	return nil
}
