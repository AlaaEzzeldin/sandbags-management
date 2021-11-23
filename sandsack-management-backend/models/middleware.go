package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Type string `json:"type"`
	jwt.StandardClaims `json:"standard_claims"`
}