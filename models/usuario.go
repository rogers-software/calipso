package models

import (
	"time"
)

type Usuario struct {
	ID              int       `json:"id"`
	Nombre          string    `json:"nombre"`
	Apellidos       string    `json:"apellidos"`
	FechaNacimiento time.Time `json:"fechanacimiento"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Avatar          string    `json:"avatar"`
	Banner          string    `json:"banner"`
	Biografia       string    `json:"biografia"`
	Ubicacion       string    `json:"ubicacion"`
	SitioWeb        string    `json:"sitioweb"`
}
