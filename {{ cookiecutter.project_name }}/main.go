package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	
	"github.com/gin-gonic/gin"
	"{{ cookiecutter.package_name }}/pkg/handlers"
	"{{ cookiecutter.package_name }}/pkg/repository"
)

func main() {
	// Initialize the database connection
	db, err := sql.Open("sqlite3", "myDatabase.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Create a Gin router
	r := gin.Default()

	// Create your repository instances here and pass the DB reference
	{{ cookiecutter.main_endpoint }}Repository := repository.New{{ cookiecutter.main_endpoint.capitalize() }}Repository(db)

	// Create handler instances here
	{{ cookiecutter.main_endpoint }}Hander := handlers.New{{ cookiecutter.main_endpoint.capitalize() }}Handler({{ cookiecutter.main_endpoint }}Repository)


	// Register your endpoints
	r.GET("/health", handlers.HealthHandler)
	r.GET("/{{ cookiecutter.main_endpoint }}/:id", {{ cookiecutter.main_endpoint }}Handler.GetById)

	// Add more endpoint registrations as needed

	// Start the server
	log.Println("Server started on http://localhost:8080")
	r.Run()
}
