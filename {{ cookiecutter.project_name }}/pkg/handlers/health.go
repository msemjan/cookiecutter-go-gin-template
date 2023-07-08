package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"{{ cookiecutter.package_name }}/pkg/models"
)

func HealthHandler(c *gin.Context) {
	log.Println("Health check triggered. The server is up and running.")

	status := models.HealthStatus{
		Message:         "Server running",
		ServerRunning:   true,
		DatabaseRunning: true,
	}

	c.JSON(http.StatusOK, status)
}
