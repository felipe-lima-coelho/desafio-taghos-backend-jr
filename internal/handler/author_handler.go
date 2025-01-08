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

func (h *AuthorHandler) FindAuthorByName(c *gin.Context) {
	authorName := c.Param("name")
	if authorName == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_REQUEST_PARAMETER",
					"message": "invalid request parameter",
					"details": "Author name is required",
				},
			},
		)
		return
	}

	author, err := h.authorService.FindAuthorByName(authorName)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "FIND_AUTHOR_BY_NAME_ERROR",
					"message": "error while finding author by name",
					"details": errors.New("Error while finding author by name: " + err.Error()),
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

func (h *AuthorHandler) FindAllAuthors(c *gin.Context) {
	authors, err := h.authorService.FindAllAuthors()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "FIND_ALL_AUTHORS_ERROR",
					"message": "error while finding all authors",
					"details": errors.New("Error while finding all authors: " + err.Error()),
				},
			},
		)
		return
	}

	if len(authors) == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "AUTHORS_NOT_FOUND",
					"message": "authors not found",
					"details": "Authors not found in the database",
				},
			},
		)
		return
	}

	var authorsResponse []*authorResponse
	for _, author := range authors {
		authorsResponse = append(authorsResponse, h.formatAuthorResponse(author))
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"authors": authorsResponse,
			},
		},
	)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	var author domain.Author
	if err := c.ShouldBindJSON(&author); err != nil {
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

	if err := h.authorService.UpdateAuthor(&author); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "UPDATE_AUTHOR_ERROR",
					"message": "error while updating author",
					"details": errors.New("Error while updating author: " + err.Error()),
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusNoContent,
		gin.H{
			"data": gin.H{
				"message": "Author updated successfully",
			},
		},
	)
}

func (h *AuthorHandler) DeleteAuthorByID(c *gin.Context) {
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

	if err := h.authorService.DeleteAuthorByID(authorID); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "DELETE_AUTHOR_BY_ID_ERROR",
					"message": "error while deleting author by ID",
					"details": errors.New("Error while deleting author by ID: " + err.Error()),
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

func (h *AuthorHandler) formatAuthorResponse(author *domain.Author) *authorResponse {
	return &authorResponse{
		ID:   author.ID,
		Name: author.Name,
	}
}
