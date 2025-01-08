package service

import (
	"errors"
	"fmt"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/repository"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBook(book *domain.Book) error
	FindBookByID(id string) (*domain.Book, error)
	FindBookByTitle(title string) (*domain.Book, error)
	FindAllBooks() ([]*domain.Book, error)
	UpdateBook(book *domain.Book) error
	DeleteBookByID(id string) error
}

type bookService struct {
	bookRepo     repository.BookRepository
	categoryRepo repository.CategoryRepository
	authorRepo   repository.AuthorRepository
}

func NewBookService(
	bookRepo repository.BookRepository,
	categoryRepo repository.CategoryRepository,
	authorRepo repository.AuthorRepository,
) BookService {
	return &bookService{bookRepo, categoryRepo, authorRepo}
}

func (s *bookService) CreateBook(book *domain.Book) error {
	ok, err := s.validateBook(book)
	if !ok {
		return fmt.Errorf("invalid book: %v", err)
	}

	book, err := s.handleCategory(book)
	if err != nil {
		return fmt.Errorf("error while handling category: %v", err)
	}

	book, err := s.handleAuthor(book)
	if err != nil {
		return fmt.Errorf("error while handling author: %v", err)
	}

	return s.bookRepo.Create(book)
}

func (s *bookService) validateBook(book *domain.Book) (bool, error) {
	if book.Title == "" {
		return false, fmt.Errorf("title is required")
	}

	if book.Synopsis == "" {
		return false, fmt.Errorf("synopsis is required")
	}

	for _, category := range book.Categories {
		if category.Name == "" {
			return false, fmt.Errorf("category name is required")
		}
	}

	for _, author := range book.Authors {
		if author.Name == "" {
			return false, fmt.Errorf("author name is required")
		}
	}

	return true, nil
}

func (s *bookService) handleCategory(book *domain.Book) (*domain.Book, error) {
	catService := NewCategoryService(s.categoryRepo)

	for _, category := range book.Categories {
		// Check if the category already exists
		// If not, create it
		_, err := catService.FindCategoryByName(category.Name)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("error while trying to find the category by name: %v", err)
			}

			if err := catService.CreateCategory(&category); err != nil {
				return nil, fmt.Errorf("error while trying to create the category: %v", err)
			}
		}
	}

	return book, nil
}
