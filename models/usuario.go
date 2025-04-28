package models

import (
	"time"
)

type Usuario struct {
	ID              int       `json:"id,omitempty"`
	Nombre          string    `json:"nombre,omitempty"`
	Apellidos       string    `json:"apellidos,omitempty"`
	FechaNacimiento time.Time `json:"fechanacimiento,omitempty"`
	Email           string    `json:"email,omitempty"`
	Password        string    `json:"password,omitempty"`
	Avatar          string    `json:"avatar,omitempty"`
	Banner          string    `json:"banner,omitempty"`
	Biografia       string    `json:"biografia,omitempty"`
	Ubicacion       string    `json:"ubicacion,omitempty"`
	SitioWeb        string    `json:"sitioweb,omitempty"`
}
