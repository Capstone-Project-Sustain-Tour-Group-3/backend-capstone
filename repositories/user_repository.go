package repositories

import (
	"capstone/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindByUsername(name string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *entities.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user *entities.User) (*entities.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(name string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("username = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
