package handler

import (
	"errors"
	"net/http"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/domain"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/service"
	"github.com/gin-gonic/gin"
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

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var request authorRequest
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

	var author domain.Author
	author.Name = request.Name

	if err := h.authorService.CreateAuthor(&author); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "CREATE_AUTHOR_ERROR",
					"message": "error while creating author",
					"details": errors.New("Error while creating author: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": gin.H{
				"message": "Author created successfully",
			},
		},
	)
}

func (h *AuthorHandler) FindAuthorByID(c *gin.Context) {
	authorID := c.Param("id")
	if authorID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_REQUEST_PARAMETER",
					"message": "invalid request parameter",
					"details": errors.New("Author ID is required"),
				},
			},
		)
		return
	}

	author, err := h.authorService.FindAuthorByID(authorID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "FIND_AUTHOR_BY_ID_ERROR",
					"message": "error while finding author by ID",
					"details": errors.New("Error while finding author by ID: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"author": h.formatAuthorResponse(author),
			},
		},
	)
}

func (h *AuthorHandler) formatAuthorResponse(author *domain.Author) *authorResponse {
	return &authorResponse{
		ID:   author.ID,
		Name: author.Name,
	}
}
