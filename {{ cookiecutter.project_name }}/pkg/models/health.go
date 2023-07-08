package models

type HealthStatus struct {
	Message         string `json:"message"`
	ServerRunning   bool   `json:"serverRunning"`
	DatabaseRunning bool   `json:"databaseRunning"`
}
