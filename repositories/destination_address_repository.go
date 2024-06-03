package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IFacilityRepository interface {
	Create(facility *entities.Facility) error
	FindAll() ([]entities.Facility, error)
	FindById(id uuid.UUID) (*entities.Facility, error)
	Update(facility *entities.Facility) error
	Delete(facility *entities.Facility) error
}

type FacilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) *FacilityRepository {
	return &FacilityRepository{db}
}

func (r *FacilityRepository) Create(facility *entities.Facility) error {
	if err := r.db.Create(facility).Error; err != nil {
		return err
	}
	return nil
}

func (r *FacilityRepository) FindAll() ([]entities.Facility, error) {
	var facilities []entities.Facility
	if err := r.db.Find(&facilities).Error; err != nil {
		return nil, err
	}
	return facilities, nil
}

func (r *FacilityRepository) FindById(id uuid.UUID) (*entities.Facility, error) {
	var facility *entities.Facility
	if err := r.db.Where("id = ?", id).First(&facility).Error; err != nil {
		return nil, err
	}
	return facility, nil
}

func (r *FacilityRepository) Update(facility *entities.Facility) error {
	if err := r.db.Save(facility).Error; err != nil {
		return err
	}
	return nil
}

func (r *FacilityRepository) Delete(facility *entities.Facility) error {
	if err := r.db.Delete(facility).Error; err != nil {
		return err
	}
	return nil
}
