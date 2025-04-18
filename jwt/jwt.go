package jwt

import (
	"context"
	"rogers-software/calipso/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GeneroJWT(ctx context.Context, t models.Usuario) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	miClave := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechanacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioweb":        t.SitioWeb,
		"id":              t.ID,
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
