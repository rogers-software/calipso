package routers

import (
	"encoding/json"
	"fmt"
	"rogers-software/calipso/database"
	"rogers-software/calipso/models"

	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.ResApi {
	var r models.ResApi
	r.Status = 400

	fmt.Println("Entre en VerPerfil()")

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}
	perfil, err := database.BuscoPerfil(database.DB, ID)
	if err != nil {
		r.Message = "Ocurrio un error al intentar buscar registro" + err.Error()
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los daroa de usuarios JSON" + err.Error()
		return r
	}
	r.Status = 200
	r.Message = string(respJson)
	return r
}
