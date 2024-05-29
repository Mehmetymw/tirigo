package repository

import (
	"errors"
	"tirigo/internal/user-management/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (domain.User, error)
	FindByUsername(username string) (domain.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user domain.User) error {
	return r.db.Create(&user).Error
}

func (r *GormUserRepository) FindByID(id string) (domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return domain.User{}, errors.New("user not found by id")
	}
	return user, result.Error
}

func (r *GormUserRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "username = ?", username)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return domain.User{}, errors.New("user not found by username")
	}
	return user, result.Error
}
