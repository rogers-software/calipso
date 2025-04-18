package database

import (
	"database/sql"
	"fmt"
	"rogers-software/calipso/models"
)

func ExisteUsuario(db *sql.DB, email string) (models.Usuario, bool, int) {

	var id int
	var nombre string
	var apellidos string

	var resultado models.Usuario

	fmt.Println("existe email ->", email)

	query := "SELECT id, nombre, apellidos FROM usuarios WHERE email = $1"

	err := db.QueryRow(query, email).Scan(&id, &nombre, &apellidos)

	if err != nil {
		fmt.Println("existe mail error ->", err)
		return resultado, false, 0
	}

	fmt.Println("id ->", id)
	fmt.Println("Nombre ->", nombre)
	fmt.Println("Apellidos ->", apellidos)

	resultado.ID = id
	resultado.Nombre = nombre
	resultado.Apellidos = apellidos

	return resultado, true, id
}
