package repository

import (
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	Create(author *domain.Author) error
	FindByID(id string) (*domain.Author, error)
	FindByName(name string) (*domain.Author, error)
	FindAll() ([]*domain.Author, error)
	Update(author *domain.Author) error
	Delete(id string) error
}

type gormAuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &gormAuthorRepository{db}
}

func (r *gormAuthorRepository) Create(author *domain.Author) error {
	return r.db.Create(author).Error
}

func (r *gormAuthorRepository) FindByID(id string) (*domain.Author, error) {
	var author domain.Author
	if err := r.db.
		First(&author, id).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *gormAuthorRepository) FindByName(name string) (*domain.Author, error) {
	var author domain.Author
	if err := r.db.
		First(&author, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *gormAuthorRepository) FindAll() ([]*domain.Author, error) {
	var authors []*domain.Author
	if err := r.db.
		Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *gormAuthorRepository) Update(author *domain.Author) error {
	return r.db.Save(author).Error
}

func (r *gormAuthorRepository) Delete(id string) error {
	return r.db.Delete(&domain.Author{}, id).Error
}
