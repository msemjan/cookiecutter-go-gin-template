package models

type {{ cookiecutter.main_endpoint.capitalize() }} struct {
	BaseModel
	Name	string `json:"name"`
	// Modify to add desired fields
}
