package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type UserPersonalizationRepository interface {
	Create(userPersonalization []*entities.UserPersonalization) error
}

type userPersonalizationRepository struct {
	db *gorm.DB
}

func NewUserPersonalizationRepository(db *gorm.DB) *userPersonalizationRepository {
	return &userPersonalizationRepository{db}
}

func (r *userPersonalizationRepository) Create(userPersonalization []*entities.UserPersonalization) error {
	return r.db.Create(userPersonalization).Error
}
