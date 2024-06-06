package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type IProvinceRepository interface {
	Create(province *entities.Province) error
	FindAll() ([]entities.Province, error)
	FindById(id string) (*entities.Province, error)
	Update(province *entities.Province) error
	Delete(province *entities.Province) error
}

type ProvinceRepository struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) *ProvinceRepository {
	return &ProvinceRepository{db}
}

func (r *ProvinceRepository) Create(province *entities.Province) error {
	if err := r.db.Create(province).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProvinceRepository) FindAll() ([]entities.Province, error) {
	var provinces []entities.Province
	if err := r.db.Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *ProvinceRepository) FindById(id string) (*entities.Province, error) {
	var province *entities.Province
	if err := r.db.Where("id = ?", id).First(&province).Error; err != nil {
		return nil, err
	}
	return province, nil
}

func (r *ProvinceRepository) Update(province *entities.Province) error {
	if err := r.db.Save(province).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProvinceRepository) Delete(province *entities.Province) error {
	if err := r.db.Delete(province).Error; err != nil {
		return err
	}
	return nil
}
