package handlers

import (
	"context"
	"fmt"
	"rogers-software/calipso/jwt"
	"rogers-software/calipso/models"
	"rogers-software/calipso/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.ResApi
	r.Status = 400

	isOk, statusCode, msg, _ := validoAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}
	fmt.Println("method->", ctx.Value(models.Key("method")).(string))
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		case "login":
			return routers.Login(ctx)
		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verperfil":
			return routers.VerPerfil(request)
		}
		//
	case "PUT":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	}
	r.Message = "Method Invalid"
	return r

}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	fmt.Println("path->", path)
	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token Requerido", models.Claim{}
	}

	claim, todoOK, _, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, "Error Token", models.Claim{}
		} else {
			fmt.Println("Error en el token... ")
			return false, 401, "Error token...", models.Claim{}
		}
	}
	fmt.Println("Token ok ")
	return true, 200, "Token ok", *claim

}
