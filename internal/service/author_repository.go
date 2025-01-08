package service

import (
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
