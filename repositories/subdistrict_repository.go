package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISubdistrictRepository interface {
	Create(subdistrict *entities.Subdistrict) error
	FindAll() ([]entities.Subdistrict, error)
	FindById(id uuid.UUID) (*entities.Subdistrict, error)
	Update(subdistrict *entities.Subdistrict) error
	Delete(subdistrict *entities.Subdistrict) error
}

type SubdistrictRepository struct {
	db *gorm.DB
}

func NewSubdistrictRepository(db *gorm.DB) *SubdistrictRepository {
	return &SubdistrictRepository{db}
}

func (r *SubdistrictRepository) Create(subdistrict *entities.Subdistrict) error {
	if err := r.db.Create(subdistrict).Error; err != nil {
		return err
	}
	return nil
}

func (r *SubdistrictRepository) FindAll() ([]entities.Subdistrict, error) {
	var subdistricts []entities.Subdistrict
	if err := r.db.Find(&subdistricts).Error; err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func (r *SubdistrictRepository) FindById(id uuid.UUID) (*entities.Subdistrict, error) {
	var subdistrict *entities.Subdistrict
	if err := r.db.Where("id = ?", id).First(&subdistrict).Error; err != nil {
		return nil, err
	}
	return subdistrict, nil
}

func (r *SubdistrictRepository) Update(subdistrict *entities.Subdistrict) error {
	if err := r.db.Save(subdistrict).Error; err != nil {
		return err
	}
	return nil
}

func (r *SubdistrictRepository) Delete(subdistrict *entities.Subdistrict) error {
	if err := r.db.Delete(subdistrict).Error; err != nil {
		return err
	}
	return nil
}
