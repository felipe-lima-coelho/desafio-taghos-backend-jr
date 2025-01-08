package handler

import (
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/service"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}
