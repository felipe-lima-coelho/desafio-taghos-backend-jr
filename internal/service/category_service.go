package service

import (
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *domain.Category) error
	FindCategoryByID(id string) (*domain.Category, error)
	FindAllCategories() ([]*domain.Category, error)
	UpdateCategory(category *domain.Category) error
	DeleteCategoryByID(id string) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}
