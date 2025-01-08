package service

import (
	"fmt"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/repository"
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
