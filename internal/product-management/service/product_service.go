package service

import (
	"time"
	"tirigo/internal/dtos"
	"tirigo/internal/product-management/domain"
	"tirigo/internal/product-management/repository"
	util "tirigo/pkg"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(productDto dtos.ProductSaveParameter) (domain.Product, error) {
	product := domain.Product{
		ID:          util.RandomID(),
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       productDto.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := s.repo.Save(product)
	return product, err
}

func (s *ProductService) Get(id string) (domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) Delete(id string) error {
	return s.repo.DeleteByID(id)
}
