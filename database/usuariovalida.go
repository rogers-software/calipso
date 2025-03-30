package database

import (
	"database/sql"
	"rogers-software/calipso/models"
)

func ExisteUsuario(db *sql.DB, email string) (models.Usuario, bool, int) {

	var id int
	var nombre string
	var apellidos string

	var resultado models.Usuario

	query := "SELECT id, nombre FROM product WHERE email = $1"

	err := db.QueryRow(query, email).Scan(&id, &nombre, &apellidos)

	if err != nil {
		return resultado, false, 0
	}

	resultado.ID = id
	resultado.Nombre = nombre
	resultado.Apellidos = apellidos

	return resultado, true, id
}
