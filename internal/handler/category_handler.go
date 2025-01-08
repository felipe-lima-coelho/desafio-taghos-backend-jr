package handler

import (
	"errors"
	"net/http"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_REQUEST_BODY",
					"message": "invalid request body",
					"details": errors.New("Error while binding JSON: " + err.Error()),
				},
			},
		)
		return
	}

	if err := h.categoryService.CreateCategory(&category); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "CREATE_CATEGORY_ERROR",
					"message": "error while creating category",
					"details": errors.New("Error while creating category: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": gin.H{
				"category": category,
			},
		},
	)
}
