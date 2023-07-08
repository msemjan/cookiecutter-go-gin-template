package repository

import (
	"database/sql"
	"errors"
	"log"

	"{{ cookiecutter.package_name }}/pkg/models"
)

type {{ cookiecutter.main_endpoint.capitalize() }}Repository struct {
	db *sql.DB,
}

func New{{ cookiecutter.main_endpoint.capitalize() }}Repository(db *sql.DB) *{{ cookiecutter.main_endpoint.capitalize() }}Repository {
	return &{{ cookiecutter.main_endpoint.capitalize() }}Repository{
		db: db,
	}
}

func (r *{{ cookiecutter.main_endpoint.capitalize() }}Repository) Save({{ cookiecutter.main_endpoint }} *models.{{ cookiecutter.main_endpoint.capitalize() }}) error {
	// Implement the logic to save the user to the SQLite3 database
	// You can use the "db" field to interact with the database
	log.Println("Saving {{ cookiecutter.main_endpoint.capitalize() }} into database")

	// Return any appropriate error or nil if successful
	return errors.New("Save method not implemented")
}

func (r *{{ cookiecutter.main_endpoint.capitalize() }}Repository) GetById(id string) (models.{{ cookiecutter.main_endpoint.capitalize() }}, error) {
	// Implement the logic to save the user to the SQLite3 database
	// You can use the "db" field to interact with the database
	log.Printf("Getting {{ cookiecutter.main_endpoint.capitalize() }} with id='%v' from database\n", id)

	// Return any appropriate error or nil if successful
	return models.{{ cookiecutter.main_endpoint.capitalize() }}{}, errors.New("Save method not implemented")
}

// Add more repository functions as needed
