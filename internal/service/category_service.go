package service

import (
	"fmt"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *domain.Category) error
	FindCategoryByID(id string) (*domain.Category, error)
	FindCategoryByName(name string) (*domain.Category, error)
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

func (s *categoryService) CreateCategory(category *domain.Category) error {
	categoryName := category.Name

	if categoryName == "" {
		return fmt.Errorf("category name is required")
	}

	// Check if the category already exists
	if _, err := s.categoryRepo.FindByName(categoryName); err == nil {
		return fmt.Errorf("category already exists")
	}

	return s.categoryRepo.Create(category)
}

func (s *categoryService) FindCategoryByID(id string) (*domain.Category, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}

	return s.categoryRepo.FindByID(id)
}

func (s *categoryService) FindCategoryByName(name string) (*domain.Category, error) {
	if name == "" {
		return nil, fmt.Errorf("category name is required")
	}

	return s.categoryRepo.FindByName(name)
}

func (s *categoryService) FindAllCategories() ([]*domain.Category, error) {
	return s.categoryRepo.FindAll()
}
