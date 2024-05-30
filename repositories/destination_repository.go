package repositories

import (
	"strings"

	"capstone/dto"
	"capstone/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IDestinationRepository interface {
	FindById(id uuid.UUID) (*entities.Destination, error)
	FindAll(page, limit int, searchQuery, sortQuery, filterQuery string) (string, *int64, []entities.Destination, error)
	FindByCategoryIds(ids []uuid.UUID) ([]entities.Destination, error)
}

type DestinationRepository struct {
	db *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) *DestinationRepository {
	return &DestinationRepository{db}
}

func (r *DestinationRepository) FindById(id uuid.UUID) (*entities.Destination, error) {
	var destination *entities.Destination
	if err := r.db.Where("id = ?", id).
		Preload("DestinationMedias").
		Preload("DestinationFacilities.Facility").
		Preload("DestinationCategories.Category").
		Preload("DestinationAddress").
		Preload("DestinationAddress.Province").
		First(&destination).Error; err != nil {
		return nil, err
	}

	return destination, nil
}

func (r *DestinationRepository) FindAll(page, limit int, searchQuery, sortQuery, filterQuery string) (string, *int64, []entities.Destination, error) {
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

	var category entities.Category
	var filterName string

	if filterQuery != "" {
		db = db.Joins("JOIN destination_categories ON destination_categories.destination_id = destinations.id").
			Where("destination_categories.category_id = ?", filterQuery)
		// Group("destinations.id")

		if err := r.db.Model(&entities.Category{}).
			Select("name").
			Where("id = ?", filterQuery).
			First(&category).Error; err != nil {
			return filterName, nil, nil, err
		}

		filterName = category.Name
	}

	if err := db.Debug().
		Offset(offset).Limit(limit).
		Preload("DestinationMedias", "type = ?", "image").
		Preload("DestinationAddress").
		Preload("DestinationAddress.Province").
		Preload("DestinationCategories.Category").
		Find(&destinations).
		Error; err != nil {
		return filterName, nil, nil, err
	}

	if err := db.Count(&total).Error; err != nil {
		return filterName, nil, nil, err
	}
	return filterName, &total, destinations, nil
}

func (r *DestinationRepository) FindByCategoryIds(ids []uuid.UUID) ([]entities.Destination, error) {
	var destinations []entities.Destination

	if err := r.db.Model(&entities.Destination{}).
		Limit(5).
		Joins("JOIN destination_categories ON destination_categories.destination_id = destinations.id").
		Where("destination_categories.category_id IN (?)", ids).
		Preload("DestinationMedias", "type = ?", "image").
		Preload("DestinationAddress").
		Preload("DestinationAddress.Province").
		Preload("DestinationCategories.Category").
		Find(&destinations).
		Error; err != nil {
		return nil, err
	}
	return destinations, nil
}
