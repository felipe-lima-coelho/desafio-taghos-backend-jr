package handler

import (
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/service"
)

type AuthorHandler struct {
	authorService service.AuthorService
}

type authorRequest struct {
	Name string `binding:"required"`
}

type authorResponse struct {
	ID   string
	Name string
}

func NewAuthorHandler(authorService service.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService}
}

func (h *AuthorHandler) formatAuthorResponse(author *authorResponse) *authorResponse {
	return &authorResponse{
		ID:   author.ID,
		Name: author.Name,
	}
}
