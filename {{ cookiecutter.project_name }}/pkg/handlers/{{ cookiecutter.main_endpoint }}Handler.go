package handlers

import (
	"log"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"{{ cookiecutter.package_name }}/pkg/repository"
)

type {{ cookiecutter.main_endpoint.capitalize() }}Handler struct {
	repository. repository.{{ cookiecutter.main_endpoint.capitalize() }}Repository
}

func New{{ cookiecutter.main_endpoint.capitalize() }}Handler(repo *repository.{{ cookiecutter.main_endpoint.capitalize() }}Repository) *{{ cookiecutter.main_endpoint.capitalize() }}Handler {
	return &{{ cookiecutter.main_endpoint.capitalize() }}Handler{
		repository: repo,
	}
}

func (handler *{{ cookiecutter.main_endpoint.capitalize() }}Handler) GetById(c *gin.Context) {
	// Re-implement the logic as needed

	log.Println("{{ cookiecutter.main_endpoint.capitalize() }}Handler.GetById hit")
	id := c.Param("id")

	record, err := handler.repository.GetById(id)
	if err != nil {
		// Handle error
		log.Printf("Error while retrieving a {{ cookiecutter.main_endpoint }} with id='%v'. Error: %v", id, err)
		return
	}

	c.JSON(http.StatusOK, record)
}

// Add more handler functions as needed
