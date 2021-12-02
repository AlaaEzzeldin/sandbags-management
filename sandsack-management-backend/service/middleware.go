package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"os"
	"team2/sandsack-management-backend/models"
	"time"
)

const ACCESS_TOKEN_TTL = time.Minute * 60
const REFRESH_TOKEN_TTL = time.Hour * 24 * 120
const VERIFY_TOKEN_TTL = time.Hour * 24
const RESET_PASSWORD_TOKEN_TTL = time.Hour * 24

var jwtPrivateKey = []byte(os.Getenv("JWT_SECRET"))



func GenerateTokens(db *gorm.DB, email string) (map[string]string, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		return nil, err
	}


	var atClaims = models.CustomClaims{
		Id:    user.Id,
		Email: email,
		Type: "access",
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


// todo
func RefreshAccessToken(refreshToken string) (accessToken string, err error) {
	return "", nil
}



