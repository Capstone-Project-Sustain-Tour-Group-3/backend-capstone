package repositories

import (
	"strings"

	"capstone/dto"
	"capstone/entities"
	"capstone/errorHandlers"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IDestinationRepository interface {
	FindById(id uuid.UUID) (*entities.Destination, error)
	FindAll(page, limit int, searchQuery, sortQuery string) (*int64, []entities.Destination, error)
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

func (r *DestinationRepository) FindAll(page, limit int, searchQuery, sortQuery string) (*int64, []entities.Destination, error) {
	var destinations []entities.Destination
	var total int64
	offset := (page - 1) * limit

	db := r.db.Model(&entities.Destination{})

	if searchQuery != "" {
		db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(searchQuery)+"%")
	}

	switch {
	case dto.DestinationSort(sortQuery) == dto.Terbaru:
		db = db.Order("created_at DESC")
	case dto.DestinationSort(sortQuery) == dto.Terlama:
		db = db.Order("created_at ASC")
	case dto.DestinationSort(sortQuery) == dto.Populer:
		db = db.Order("visit_count DESC")
	}

	if err := db.Debug().
		Offset(offset).Limit(limit).
		Preload("DestinationMedias", "type = ?", "image").
		Preload("Categories").
		Preload("Facilities").
		Preload("DestinationAddress").
		Preload("DestinationAddress.Province").
		Find(&destinations).
		Error; err != nil {
		return nil, nil, err
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return &total, destinations, nil
}
