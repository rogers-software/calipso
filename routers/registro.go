package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"rogers-software/calipso/database"
	"rogers-software/calipso/models"
)

func Registro(ctx context.Context) models.ResApi {
	var t models.Usuario
	var r models.ResApi
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = err.Error()
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Debe especifical al menos 6 caracteres en password"
		fmt.Println(r.Message)
		return r
	}

	db := database.GetConnection()

	_, encontrado, _ := database.ExisteUsuario(db, t.Email)

	if encontrado {
		r.Message = "Ya existe un usuario registrado con este email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := database.InsertoRegistro(db, t)

	if err != nil {
		r.Message = "Ocurrio un error al intentar relializar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Regisro ok"
	fmt.Println(r.Message)
	return r
}
