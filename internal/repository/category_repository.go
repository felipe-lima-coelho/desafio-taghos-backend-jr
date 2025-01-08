package repository

import (
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	FindByID(id string) (*domain.Category, error)
	FindByName(name string) (*domain.Category, error)
	FindAll() ([]*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id string) error
}

type gormCategoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoriesRepository{db}
}

func (r *gormCategoriesRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *gormCategoriesRepository) FindByID(id string) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.
		First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *gormCategoriesRepository) FindByName(name string) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.
		First(&category, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *gormCategoriesRepository) FindAll() ([]*domain.Category, error) {
	var categories []*domain.Category
	if err := r.db.
		Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *gormCategoriesRepository) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *gormCategoriesRepository) Delete(id string) error {
	return r.db.Delete(&domain.Category{}, id).Error
}
