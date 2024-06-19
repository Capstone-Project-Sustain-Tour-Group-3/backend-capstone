package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalizationRepository interface {
	Create(userPersonalization *entities.UserPersonalization) error
	FindByUserId(uid uuid.UUID) (*entities.UserPersonalization, error)
}

type userPersonalizationRepository struct {
	db *gorm.DB
}

func NewUserPersonalizationRepository(db *gorm.DB) *userPersonalizationRepository {
	return &userPersonalizationRepository{db}
}

func (r *userPersonalizationRepository) Create(userPersonalization *entities.UserPersonalization) error {
	return r.db.Create(userPersonalization).Error
}

func (r *userPersonalizationRepository) FindByUserId(uid uuid.UUID) (*entities.UserPersonalization, error) {
	userPersonalization := new(entities.UserPersonalization)
	db := r.db.Begin()
	if err := db.
		Preload("PersonalizationProvinces").
		Preload("PersonalizationCategories").
		First(userPersonalization, "user_id = ?", uid).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	return userPersonalization, nil
}
