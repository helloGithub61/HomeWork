package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}
