package jwt

import (
	"errors"
	"fmt"
	"rogers-software/calipso/database"
	"rogers-software/calipso/models"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario int

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, int, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, 0, errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		// Rutina que cheque con la BD
		fmt.Println("email ->", claims.Email)
		_, encontrado, _ := database.ExisteUsuario(database.DB, claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID
		}
		return &claims, encontrado, IDUsuario, nil

	}

	if !tkn.Valid {
		return &claims, false, 0, errors.New("token invalido")
	}

	return &claims, false, 0, err

}
