package repositories

import (
	"capstone/entities"
	"capstone/errorHandlers"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IDestinationRepository interface {
	FindById(id uuid.UUID) (*entities.Destination, error)
	FindAll() ([]entities.Destination, error)
}

type DestinationRepository struct {
	db *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) *DestinationRepository {
	return &DestinationRepository{db}
}

func (r *DestinationRepository) FindById(id uuid.UUID) (*entities.Destination, error) {
	var destination *entities.Destination
	if err := r.db.Where("id = ?", id).First(&destination).Error; err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "gagal mencari tempat wisata"}
	}
	return destination, nil
}

func (r *DestinationRepository) FindAll() ([]entities.Destination, error) {
	var destinations []entities.Destination
	if err := r.db.
		Preload("DestinationMedias", "type = ?", "image").
		Preload("Categories").
		Preload("Facilities").
		Preload("DestinationAddress").
		Preload("DestinationAddress.Province").
		Find(&destinations).
		Error; err != nil {
		return nil, &errorHandlers.InternalServerError{Message: "gagal untuk melihat detail wisata"}
	}
	return destinations, nil
}
