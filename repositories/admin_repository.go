package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*entities.Admin, error)
	GetUserByRefreshToken(refreshToken string) (*entities.Admin, error)
	Update(admin *entities.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByUsername(username string) (*entities.Admin, error) {
	var admin *entities.Admin
	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) GetUserByRefreshToken(refreshToken string) (*entities.Admin, error) {
	var admin *entities.Admin
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) Update(admin *entities.Admin) error {
	if err := r.db.Save(admin).Error; err != nil {
		return err
	}
	return nil
}
