package repository

import (
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *domain.Book) error
	FindByID(id string) (*domain.Book, error)
	FindByTitle(title string) (*domain.Book, error)
	FindAll() ([]*domain.Book, error)
	Update(book *domain.Book) error
	Delete(id string) error
}

type gormBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &gormBookRepository{db}
}

func (r *gormBookRepository) Create(book *domain.Book) error {
	return r.db.Create(book).Error
}

func (r *gormBookRepository) FindByID(id string) (*domain.Book, error) {
	var book domain.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *gormBookRepository) FindByTitle(title string) (*domain.Book, error) {
	var book domain.Book
	if err := r.db.First(&book, "title = ?", title).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
