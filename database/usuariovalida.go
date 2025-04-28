package database

import (
	"database/sql"
	"fmt"
	"rogers-software/calipso/models"
	"time"
)

func ExisteUsuario(db *sql.DB, email string) (models.Usuario, bool, int) {

	var id int
	var nombre string
	var apellidos string
	var fechanacimiento time.Time
	var password string
	var avatar string
	var banner string
	var biografia string
	var ubicacion string
	var sitioweb string

	var resultado models.Usuario

	fmt.Println("EXISTEUSUARIO:email ->", email)

	query := "SELECT id, nombre, apellidos, fechanacimiento, email, password, avatar, banner, biografia, ubicacion, sitioweb FROM usuarios WHERE email = $1"

	err := db.QueryRow(query, email).Scan(&id, &nombre, &apellidos, &fechanacimiento, &email, &password, &avatar, &banner, &biografia, &ubicacion, &sitioweb)

	if err != nil {
		fmt.Println("EXISTEUSUARIO:mail error ->", err)
		return resultado, false, 0
	}

	fmt.Println("id ->", id)
	fmt.Println("Nombre ->", nombre)
	fmt.Println("Apellidos ->", apellidos)

	resultado.ID = id
	resultado.Nombre = nombre
	resultado.Apellidos = apellidos
	resultado.FechaNacimiento = fechanacimiento
	resultado.Email = email
	// resultado.Password = password
	resultado.Avatar = avatar
	resultado.Banner = banner
	resultado.Biografia = biografia
	resultado.Ubicacion = ubicacion
	resultado.SitioWeb = sitioweb

	return resultado, true, id
}
