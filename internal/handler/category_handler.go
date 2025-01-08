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

type categoryRequest struct {
	Name string `binding:"required"`
}

type categoryResponse struct {
	ID   string
	Name string
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var request categoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
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

	var category domain.Category
	category.Name = request.Name

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
				"category": h.formatCategoryDataReturn(&category),
			},
		},
	)
}

func (h *CategoryHandler) FindCategoryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_REQUEST_PARAMETER",
					"message": "invalid request parameter",
					"details": errors.New("ID is required"),
				},
			},
		)
		return
	}

	category, err := h.categoryService.FindCategoryByID(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "CATEGORY_NOT_FOUND",
					"message": "error while finding category by ID",
					"details": errors.New("Error while finding category by ID: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"category": h.formatCategoryDataReturn(category),
			},
		},
	)
}

func (h *CategoryHandler) FindCategoryByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_REQUEST_PARAMETER",
					"message": "invalid request parameter",
					"details": errors.New("Name is required"),
				},
			},
		)
		return
	}

	category, err := h.categoryService.FindCategoryByName(name)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "CATEGORY_NOT_FOUND",
					"message": "error while finding category by name",
					"details": errors.New("Error while finding category by name: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"category": h.formatCategoryDataReturn(category),
			},
		},
	)
}

func (h *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := h.categoryService.FindAllCategories()
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "CATEGORIES_NOT_FOUND",
					"message": "error while finding all categories",
					"details": errors.New("Error while finding all categories: " + err.Error()),
				},
			},
		)
		return
	}

	catFormated := []categoryResponse{}
	for _, category := range categories {
		catFormated = append(catFormated, h.formatCategoryDataReturn(category))
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"categories": catFormated,
			},
		},
	)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
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

	if err := h.categoryService.UpdateCategory(&category); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "UPDATE_CATEGORY_ERROR",
					"message": "error while updating category",
					"details": errors.New("Error while updating category: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"category": h.formatCategoryDataReturn(&category),
			},
		},
	)
}

func (h *CategoryHandler) DeleteCategoryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_REQUEST_PARAMETER",
					"message": "invalid request parameter",
					"details": errors.New("ID is required"),
				},
			},
		)
		return
	}

	if err := h.categoryService.DeleteCategoryByID(id); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "DELETE_CATEGORY_ERROR",
					"message": "error while deleting category",
					"details": errors.New("Error while deleting category: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusNoContent,
		gin.H{},
	)
}

func (h *CategoryHandler) formatCategoryDataReturn(category *domain.Category) categoryResponse {
	return categoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
