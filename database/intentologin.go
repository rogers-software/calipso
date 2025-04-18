package database

import (
	"fmt"
	"rogers-software/calipso/models"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ExisteUsuario(DB, email)

	fmt.Println("Existe Usuario ->", encontrado)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	fmt.Println("Password bytes error->", err)
	fmt.Println("Password dbase error->", err)
	fmt.Println("bcrypt error->", err)

	if err != nil {
		return usu, false
	}

	return usu, true

}
