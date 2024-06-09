package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type IDestinationMediaRepository interface {
	Create(destinationMedia *entities.DestinationMedia, tx *gorm.DB) error
	FindAll(page, limit int, searchQuery string) (*int64, []entities.DestinationMedia, error)
	FindById(id uuid.UUID) (*entities.DestinationMedia, error)
	Update(destinationMedia *entities.DestinationMedia) error
	Delete(destinationMedia *entities.DestinationMedia) error
	DeleteMany(destinationMedias *[]entities.DestinationMedia) error
}

type DestinationMediaRepository struct {
	db *gorm.DB
}

func NewDestinationMediaRepository(db *gorm.DB) *DestinationMediaRepository {
	return &DestinationMediaRepository{db}
}

func (r *DestinationMediaRepository) Create(destinationMedia *entities.DestinationMedia, tx *gorm.DB) error {
	var db *gorm.DB

	if tx != nil {
		db = tx
	} else {
		db = r.db
	}

	if err := db.Create(destinationMedia).Error; err != nil {
		return err
	}
	return nil
}

func (r *DestinationMediaRepository) FindAll(page, limit int, searchQuery string) (*int64, []entities.DestinationMedia, error) {
	var destinationMedias []entities.DestinationMedia
	var total int64

	offset := (page - 1) * limit

	db := r.db.Model(&entities.DestinationMedia{}).Joins("JOIN destinations ON destinations.id = destination_media.destination_id")

	if searchQuery != "" {
		db = db.Where("LOWER(destinations.name) LIKE ?", "%"+strings.ToLower(searchQuery)+"%")
	}

	if err := db.Debug().
		Limit(limit).
		Offset(offset).
		Preload("Destination").
		Find(&destinationMedias).Error; err != nil {
		return nil, nil, err
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	return &total, destinationMedias, nil
}

func (r *DestinationMediaRepository) FindById(id uuid.UUID) (*entities.DestinationMedia, error) {
	var destinationMedia *entities.DestinationMedia
	if err := r.db.Where("id = ?", id).
		Preload("Destination").
		First(&destinationMedia).Error; err != nil {
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

func (r *DestinationMediaRepository) DeleteMany(destinationMedias *[]entities.DestinationMedia) error {
	if len(*destinationMedias) == 0 {
		return nil
	}

	var ids []uuid.UUID

	for _, destinationMedia := range *destinationMedias {
		ids = append(ids, destinationMedia.Id)
	}

	if err := r.db.Where("id IN ?", ids).Delete(&entities.DestinationMedia{}).Error; err != nil {
		return err
	}
	return nil
}
