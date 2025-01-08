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
