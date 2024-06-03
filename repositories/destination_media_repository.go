package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDestinationMediaRepository interface {
	Create(destinationMedia *entities.DestinationMedia) error
	FindAll() ([]entities.DestinationMedia, error)
	FindById(id uuid.UUID) (*entities.DestinationMedia, error)
	Update(destinationMedia *entities.DestinationMedia) error
	Delete(destinationMedia *entities.DestinationMedia) error
}

type DestinationMediaRepository struct {
	db *gorm.DB
}

func NewDestinationMediaRepository(db *gorm.DB) *DestinationMediaRepository {
	return &DestinationMediaRepository{db}
}

func (r *DestinationMediaRepository) Create(destinationMedia *entities.DestinationMedia) error {
	if err := r.db.Create(destinationMedia).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationMediaRepository) FindAll() ([]entities.DestinationMedia, error) {
	var destinationMedias []entities.DestinationMedia
	if err := r.db.Find(&destinationMedias).Error; err != nil {
		return nil, err
	}
	return destinationMedias, nil
}

func (r *DestinationMediaRepository) FindById(id uuid.UUID) (*entities.DestinationMedia, error) {
	var destinationMedia *entities.DestinationMedia
	if err := r.db.Where("id = ?", id).First(&destinationMedia).Error; err != nil {
		return nil, err
	}
	return destinationMedia, nil
}

func (r *DestinationMediaRepository) Update(destinationMedia *entities.DestinationMedia) error {
	if err := r.db.Save(destinationMedia).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationMediaRepository) Delete(destinationMedia *entities.DestinationMedia) error {
	if err := r.db.Delete(destinationMedia).Error; err != nil {
		return err
	}
	return nil
}
