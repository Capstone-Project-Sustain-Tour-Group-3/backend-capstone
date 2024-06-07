package repositories

import (
	"time"

	"capstone/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll(offset, limit int, search string) (*[]entities.Admin, *int64, error)
	FindById(id uuid.UUID) (*entities.Admin, error)
	FindByUsername(username string) (*entities.Admin, error)
	GetUserByRefreshToken(refreshToken string) (*entities.Admin, error)
	Create(admin *entities.Admin) error
	Update(admin *entities.Admin) error
	Delete(admin *entities.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindAll(offset, limit int, search string) (*[]entities.Admin, *int64, error) {
	admins := new([]entities.Admin)
	count := new(int64)

	db := r.db.Model(entities.Admin{}).
		Where("role = ? AND username LIKE ?", "admin", "%"+search+"%").
		Session(&gorm.Session{})

	res := db.
		Limit(limit).
		Offset(offset).
		Find(admins)
	if res.Error != nil {
		return nil, nil, res.Error
	}

	err := db.Count(count).Error
	if err != nil {
		return nil, nil, err
	}

	return admins, count, nil
}

func (r *adminRepository) FindById(id uuid.UUID) (*entities.Admin, error) {
	admin := new(entities.Admin)
	if err := r.db.First(&admin, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return admin, nil
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
	return r.db.Save(admin).Error
}

func (r *adminRepository) Create(admin *entities.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) Delete(admin *entities.Admin) error {
	tx := r.db.Begin()

	err := tx.Model(admin).Update("delete_milli", time.Now().UnixMilli()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(admin).Update("deleted_at", time.Now()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
