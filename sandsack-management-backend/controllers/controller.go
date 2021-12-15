package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer = "Bearer "
		header := c.GetHeader("Authorization")
		tokenStr := header[len(bearer):]
		token, err := service.VerifyToken(tokenStr)
		if err != nil {
			log.Println("VerifyToken error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				ErrCode: http.StatusUnauthorized,
				ErrMessage: "no access",
			})
			return
		}

		if token.Valid {
			if claims, ok := token.Claims.(models.CustomClaims); ok && token.Valid && claims.Type == "access" {
				c.Next()
			}
		} else {
			log.Println("Token is not valid")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func (a *App) AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer = "Bearer "
		header := c.GetHeader("Authorization")
		tokenStr := header[len(bearer):]
		token, err := service.VerifyToken(tokenStr)
		if err != nil {
			log.Println("VerifyToken error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				ErrCode: http.StatusUnauthorized,
				ErrMessage: "no access",
			})
		}

		if token.Valid {
			if claims, ok := token.Claims.(models.CustomClaims); ok && token.Valid && claims.Type == "access" {
				user, err := service.GetUserByEmail(a.DB, claims.Email)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
						ErrCode: http.StatusInternalServerError,
						ErrMessage: "something went wrong",
					})
				}
				if !user.IsSuperUser{
					c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
						ErrCode: http.StatusUnauthorized,
						ErrMessage: "no access",
					})
				}
				c.Next()
			}
		} else {
			log.Println("Token is not valid")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

func GetClaims(c *gin.Context) (*models.CustomClaims, error){
	const bearer = "Bearer "
	header := c.GetHeader("Authorization")
	tokenStr := header[len(bearer):]

	claims, err := service.GetClaims(tokenStr)
	if err != nil {
		return nil, err
	}

	return claims, nil
}