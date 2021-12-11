package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Type string `json:"type"`
	Role string `json:"role"`
	jwt.StandardClaims `json:"standard_claims"`
}

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}