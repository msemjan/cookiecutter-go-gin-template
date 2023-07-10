package models

type {{ cookiecutter.main_endpoint.capitalize() }} struct {
	Id		string `json:"id"`
	Name	string `json:"name"`
	// Modify to add desired fields
}
