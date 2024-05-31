package repository

import (
	"tirigo/internal/product-management/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product domain.Product) error
	FindByID(id string) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	DeleteByID(id string) error
}

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) Save(product domain.Product) error {
	return r.db.Create(product).Error
}

func (r *GormProductRepository) FindByID(id string) (domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, "id = ?", id)
	return product, result.Error
}

func (r *GormProductRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *GormProductRepository) DeleteByID(id string) error {
	return r.db.Delete(&domain.Product{}, "id = ?", id).Error
}
