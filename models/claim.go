package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}
