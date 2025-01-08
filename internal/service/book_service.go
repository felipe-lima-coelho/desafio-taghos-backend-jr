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

	book, _, err = s.handleCategory(book)
	if err != nil {
		return fmt.Errorf("error in book_services while handling category: %v", err)
	}

	book, _, err = s.handleAuthor(book)
	if err != nil {
		return fmt.Errorf("error in book_services while handling author: %v", err)
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

func (s *bookService) handleCategory(book *domain.Book) (*domain.Book, *bool, error) {
	trueValue := true
	falseValue := false
	isCategoryCreated := &falseValue
	catService := NewCategoryService(s.categoryRepo)

	for _, category := range book.Categories {
		// Check if the category already exists
		// If not, create it
		_, err := catService.FindCategoryByName(category.Name)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil, fmt.Errorf("error in book_services while trying to find the category by name: %v", err)
			}

			if err := catService.CreateCategory(&category); err != nil {
				return nil, nil, fmt.Errorf("error in book_services while trying to create the category: %v", err)
			}

			isCategoryCreated = &trueValue
		}
	}

	return book, isCategoryCreated, nil
}

func (s *bookService) handleAuthor(book *domain.Book) (*domain.Book, *bool, error) {
	trueValue := true
	falseValue := false
	isAuthorCreated := &falseValue
	authorService := NewAuthorService(s.authorRepo)

	for _, author := range book.Authors {
		// Check if the author already exists
		// If not, create it
		_, err := authorService.FindAuthorByName(author.Name)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil, fmt.Errorf("error in book_services while trying to find the author by name: %v", err)
			}

			if err := authorService.CreateAuthor(&author); err != nil {
				return nil, nil, fmt.Errorf("error in book_services while trying to create the author: %v", err)
			}

			isAuthorCreated = &trueValue
		}
	}

	return book, isAuthorCreated, nil
}

func (s *bookService) FindBookByID(id string) (*domain.Book, error) {
	if id == "" {
		return nil, fmt.Errorf("book ID is required")
	}

	book, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error in book_services while trying to find the book by ID: %v", err)
	}

	return book, nil
}

func (s *bookService) FindBookByTitle(title string) (*domain.Book, error) {
	if title == "" {
		return nil, fmt.Errorf("book title is required")
	}

	book, err := s.bookRepo.FindByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("error in book_services while trying to find the book by title: %v", err)
	}

	return book, nil
}

func (s *bookService) FindAllBooks() ([]*domain.Book, error) {
	return s.bookRepo.FindAll()
}

func (s *bookService) UpdateBook(book *domain.Book) error {
	bookID := book.ID

	if bookID == "" {
		return fmt.Errorf("book ID is required")
	}

	bookOnDB, err := s.FindBookByID(bookID)
	if err != nil {
		return fmt.Errorf("error in book_services while trying to find the book by ID: %v", err)
	}

	bookOnDB, isCatCreated, err := s.handleCategory(bookOnDB)
	if err != nil {
		return fmt.Errorf("error in book_services while handling category: %v", err)
	}

	bookOnDB, isAutCreated, err := s.handleAuthor(bookOnDB)
	if err != nil {
		return fmt.Errorf("error in book_services while handling author: %v", err)
	}

	var isTitleChanged, isSynopsisChanged bool
	if book.Title != bookOnDB.Title {
		bookOnDB.Title = book.Title
		isTitleChanged = true
	}
	if book.Synopsis != bookOnDB.Synopsis {
		bookOnDB.Synopsis = book.Synopsis
		isSynopsisChanged = true
	}

	if !*isCatCreated && !*isAutCreated && !isTitleChanged && !isSynopsisChanged {
		// If the title, synopsis, category and author
		// are not changed, there is no need to update the book
		return nil
	}

	err = s.bookRepo.Update(bookOnDB)
	if err != nil {
		return fmt.Errorf("error in book_services while trying to update the book: %v", err)
	}

	return nil
}
