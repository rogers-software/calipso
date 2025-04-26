package database

import (
	"database/sql"
	"fmt"
	"rogers-software/calipso/models"
	"strconv"
	"time"
)

func BuscoPerfil(db *sql.DB, ID string) (models.Usuario, error) {
	var id int
	var nombre string
	var apellidos string
	var fechanacimiento time.Time
	var email string
	var password string
	var avatar string
	var banner string
	var biografia string
	var ubicacion string
	var sitioweb string

	var perfil models.Usuario
	id, errInt := strconv.Atoi(ID)
	if errInt != nil {
		fmt.Println("id error ->", errInt)
		return perfil, errInt
	}

	fmt.Println("procesando ID ->", id)

	query := "SELECT nombre, apellidos, fechanacimiento, email, password, avatar, banner, biografia, ubicacion, sitioweb FROM usuarios WHERE Id = $1"

	err := db.QueryRow(query, id).Scan(&nombre, &apellidos, &fechanacimiento, &email, &password, &avatar, &banner, &biografia, &ubicacion, &sitioweb)

	if err != nil {
		fmt.Println("error al buscar dtatabase ->", err)
		return perfil, err
	}
	perfil.ID = id
	perfil.Nombre = nombre
	perfil.Apellidos = apellidos
	perfil.FechaNacimiento = fechanacimiento
	perfil.Email = email
	//perfil.Password = password
	perfil.Avatar = avatar
	perfil.Banner = banner
	perfil.Biografia = biografia
	perfil.Ubicacion = ubicacion
	perfil.SitioWeb = sitioweb

	return perfil, nil
}
