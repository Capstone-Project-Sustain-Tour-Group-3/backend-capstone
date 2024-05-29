package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindByUsername(name string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByEmail(email string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) Create(user *entities.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) Update(user *entities.User) (*entities.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) FindByUsername(name string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("username = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
