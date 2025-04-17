package database

import (
	"database/sql"
	"fmt"
	"rogers-software/calipso/models"
)

func InsertoRegistro(db *sql.DB, u models.Usuario) (int, bool, error) {
	var err error
	u.Password, err = EncriptarPassword(u.Password)
	query := `INSERT INTO usuarios (nombre, apellidos, fechanacimiento, email, password, avatar, banner, biografia, ubicacion, sitioweb)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	var pk int
	fmt.Println("Nombre->", u.Nombre)
	errInsert := db.QueryRow(query, u.Nombre, u.Apellidos, u.FechaNacimiento, u.Email, u.Password, u.Avatar, u.Banner, u.Biografia, u.Ubicacion, u.SitioWeb).Scan(&pk)
	if errInsert != nil {
		fmt.Println("Error", err)
		return 0, false, err
	}

	return pk, true, nil
}
