package main

import (
	"log"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/config"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/repository"
	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db := config.Database()

	// Initialize the repositories
	bookRepo := repository.NewBookRepository(db)
	categoryRepo := repository.NewCategoriesRepository(db)
	authorRepo := repository.NewAuthorRepository(db)

	// Initialize the services
	bookService := service.NewBookService(bookRepo, categoryRepo, authorRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	authorService := service.NewAuthorService(authorRepo)
}
