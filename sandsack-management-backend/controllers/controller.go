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
		const bearer = "Bearer"
		header := c.GetHeader("Authorization")
		tokenStr := header[len(bearer):]
		token, err := service.VerifyToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				ErrCode: http.StatusUnauthorized,
				ErrMessage: "something went wrong",
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