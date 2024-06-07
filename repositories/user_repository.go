package repositories

import (
	"strings"

	"capstone/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(page, limit int, searchQuery string) (*[]entities.User, *int64, error)
	FindByEmail(email string) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindByUsername(name string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(user *entities.User) error
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

func (r *userRepository) FindAll(page, limit int, searchQuery string) (*[]entities.User, *int64, error) {
	var users []entities.User
	var total int64
	offset := (page - 1) * limit
	db := r.db.Model(&entities.User{})
	if searchQuery != "" {
		db = db.Where("LOWER(username) LIKE ? OR LOWER(fullname) LIKE ?",
			"%"+strings.ToLower(searchQuery)+"%", "%"+strings.ToLower(searchQuery)+"%")
	}

	if err := db.Debug().
		Offset(offset).Limit(limit).
		Find(&users).
		Error; err != nil {
		return nil, nil, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return &users, &total, nil
}

func (r *userRepository) FindById(id uuid.UUID) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
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

func (r *userRepository) Create(user *entities.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user *entities.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(user *entities.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
