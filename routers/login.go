package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/mail"
	"rogers-software/calipso/database"
	"rogers-software/calipso/jwt"
	"rogers-software/calipso/models"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.ResApi {
	var t models.Usuario
	var r models.ResApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = "Usuario y/o contraseña invalidos " + err.Error()
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Email del usuario es requerido "
		return r
	}

	// validar @ .
	if !validEmail(t.Email) {
		r.Message = "Email del usuario es incorrecto " + t.Email
		return r
	}

	userData, existe := database.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Email y/o usuario no registrado en sistema " + t.Email
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)

	if err != nil {
		r.Message = "Ocurrio error al intentar generar token " + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocurrio error al intentar formatear el token a JSON " + err2.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set.Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r

}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
