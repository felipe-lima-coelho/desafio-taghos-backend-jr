package service

import (
	"fmt"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/repository"
)

type AuthorService interface {
	CreateAuthor(author *domain.Author) error
	FindAuthorByID(id string) (*domain.Author, error)
	FindAuthorByName(name string) (*domain.Author, error)
	FindAllAuthors() ([]*domain.Author, error)
	UpdateAuthor(author *domain.Author) error
	DeleteAuthorByID(id string) error
}

type authorService struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorService(authorRepo repository.AuthorRepository) AuthorService {
	return &authorService{authorRepo}
}

func (s *authorService) CreateAuthor(author *domain.Author) error {
	authorName := author.Name

	if authorName == "" {
		return fmt.Errorf("author name is required")
	}

	// Check if the author already exists
	if _, err := s.FindAuthorByName(authorName); err == nil {
		return fmt.Errorf("author already exists")
	}

	return s.authorRepo.Create(author)
}

func (s *authorService) FindAuthorByID(id string) (*domain.Author, error) {
	if id == "" {
		return nil, fmt.Errorf("author ID is required")
	}

	return s.authorRepo.FindByID(id)
}

func (s *authorService) FindAuthorByName(name string) (*domain.Author, error) {
	if name == "" {
		return nil, fmt.Errorf("author name is required")
	}

	return s.authorRepo.FindByName(name)
}

func (s *authorService) FindAllAuthors() ([]*domain.Author, error) {
	return s.authorRepo.FindAll()
}

func (s *authorService) UpdateAuthor(author *domain.Author) error {
	authorID := author.ID
	newAuthorName := author.Name

	if authorID == "" {
		return fmt.Errorf("author ID is required")
	}
	if newAuthorName == "" {
		return fmt.Errorf("author name is required")
	}

	authorOnDB, err := s.FindAuthorByID(authorID)
	if err != nil {
		return fmt.Errorf("error while trying to find the author by ID: %v", err)
	}

	var isNameChanged bool
	if newAuthorName != authorOnDB.Name {
		authorOnDB.Name = newAuthorName
		isNameChanged = true
	}

	if !isNameChanged {
		// If the name is not changed, there is no need to update the category
		return nil
	}

	return s.authorRepo.Update(authorOnDB)
}
