package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/middleware"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer = "Bearer "
		header := c.GetHeader("Authorization")
		if len(header) == 0 {
			log.Println("Fehler: Header ist leer ")
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				ErrCode:    http.StatusBadRequest,
				ErrMessage: "Header ist leer",
			})
			return
		}
		tokenStr := header[len(bearer):]
		token, err := middleware.VerifyToken(tokenStr)
		if err != nil {
			log.Println("Fehler: VerifyToken", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				ErrCode:    http.StatusUnauthorized,
				ErrMessage: "Kein Zugang",
			})
			return
		}

		if token.Valid {
			if claims, ok := token.Claims.(models.CustomClaims); ok && token.Valid && claims.Type == "access" {
				c.Next()
			}
		} else {
			log.Println("Token ist ungültig")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func (a *App) AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer = "Bearer "
		header := c.GetHeader("Authorization")
		tokenStr := header[len(bearer):]
		token, err := middleware.VerifyToken(tokenStr)
		if err != nil {
			log.Println("Fehler VerifyToken", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				ErrCode:    http.StatusUnauthorized,
				ErrMessage: "Kein Zugang",
			})
		}

		if token.Valid {
			if claims, ok := token.Claims.(models.CustomClaims); ok && token.Valid && claims.Type == "access" {
				user, err := service.GetUserByEmail(a.DB, claims.Email)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
						ErrCode:    http.StatusInternalServerError,
						ErrMessage: "Da ist etwas schief gelaufen",
					})
				}
				if !user.IsSuperUser {
					c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
						ErrCode:    http.StatusUnauthorized,
						ErrMessage: "Kein Zugang",
					})
				}
				c.Next()
			}
		} else {
			log.Println("Token ist ungültig")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

func GetClaims(c *gin.Context) (*models.CustomClaims, error) {
	const bearer = "Bearer "
	header := c.GetHeader("Authorization")
	tokenStr := header[len(bearer):]

	claims, err := middleware.GetClaims(tokenStr)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
