package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"os"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
	"time"
)

const ACCESS_TOKEN_TTL = time.Minute * 60
const REFRESH_TOKEN_TTL = time.Hour * 24 * 120
const VERIFY_TOKEN_TTL = time.Hour * 24
const RESET_PASSWORD_TOKEN_TTL = time.Hour * 24

var jwtPrivateKey = []byte(os.Getenv("JWT_SECRET"))




func GenerateTokens(db *gorm.DB, email string) (map[string]string, error) {
	user, err := service.GetUserByEmail(db, email)
	if err != nil {
		return nil, err
	}

	role := "user"
	if user.IsSuperUser {
		role = "admin"
	}

	var atClaims = models.CustomClaims{
		Id:    user.Id,
		Email: email,
		Type: "access",
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_TTL).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := token.SignedString(jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	var rtClaims = models.CustomClaims{
		Id:    user.Id,
		Email: email,
		Type: "refresh",
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_TTL).Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := token.SignedString(jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}, nil

}

func VerifyToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Unexpected signing method",)
		}
		return jwtPrivateKey, nil
	})
	return token, err
}

func GetEmail(encodedToken string) (email string, err error) {
	token, err := jwt.ParseWithClaims(encodedToken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Unexpected signing method",)
		}
		return jwtPrivateKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok  := token.Claims.(*models.CustomClaims)
	if !ok {
		return "", errors.New("something went wrong")
	}

	return claims.Email, nil

}

func GetClaims(encodedToken string) (*models.CustomClaims, error){
	token, err := jwt.ParseWithClaims(encodedToken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Unexpected signing method",)
		}
		return jwtPrivateKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok  := token.Claims.(*models.CustomClaims)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return claims, nil
}

