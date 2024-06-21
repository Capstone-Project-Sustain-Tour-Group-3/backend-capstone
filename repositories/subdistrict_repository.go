package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type ISubdistrictRepository interface {
	Create(subdistrict *entities.Subdistrict) error
	FindAll(cityId string) ([]entities.Subdistrict, error)
	FindById(id string) (*entities.Subdistrict, error)
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

func (r *SubdistrictRepository) FindAll(cityId string) ([]entities.Subdistrict, error) {
	var subdistricts []entities.Subdistrict

	db := r.db.Model(&entities.Subdistrict{})

	if cityId != "" {
		db = db.Where("city_id = ?", cityId)
	}

	if err := db.Find(&subdistricts).Error; err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func (r *SubdistrictRepository) FindById(id string) (*entities.Subdistrict, error) {
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
