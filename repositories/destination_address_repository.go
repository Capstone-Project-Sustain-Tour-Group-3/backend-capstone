package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDestinationAddressRepository interface {
	Create(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error
	FindAll() ([]entities.DestinationAddress, error)
	FindById(id uuid.UUID) (*entities.DestinationAddress, error)
	Update(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error
	Delete(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error
}

type DestinationAddressRepository struct {
	db *gorm.DB
}

func NewDestinationAddressRepository(db *gorm.DB) *DestinationAddressRepository {
	return &DestinationAddressRepository{db}
}

func (r *DestinationAddressRepository) Create(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error {
	if err := tx.Create(destinationAddress).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationAddressRepository) FindAll() ([]entities.DestinationAddress, error) {
	var destinationAddresses []entities.DestinationAddress
	if err := r.db.Find(&destinationAddresses).Error; err != nil {
		return nil, err
	}
	return destinationAddresses, nil
}

func (r *DestinationAddressRepository) FindById(id uuid.UUID) (*entities.DestinationAddress, error) {
	var destinationAddress *entities.DestinationAddress
	if err := r.db.Where("id = ?", id).First(&destinationAddress).Error; err != nil {
		return nil, err
	}
	return destinationAddress, nil
}

func (r *DestinationAddressRepository) Update(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error {
	if err := tx.Save(destinationAddress).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationAddressRepository) Delete(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error {
	if err := tx.Delete(destinationAddress).Error; err != nil {
		return err
	}
	return nil
}
