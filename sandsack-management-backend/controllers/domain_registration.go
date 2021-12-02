package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)


func (a *App) Registration(c *gin.Context) {
	var input models.RegistrationInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("Registration error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	const bearer = "Bearer"
	header := c.GetHeader("Authorization")
	tokenStr := header[len(bearer):]

	meUser, err := service.GetUserByToken(a.DB, tokenStr)
	if err != nil {
		log.Println("GetUserByToken error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	if !meUser.IsSuperUser {
		log.Println("Trying to create user error: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "you don't have this right",
		})
		return
	}

	if err := service.CreateUser(a.DB, &input); err != nil {
		log.Println("CreateUser error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, models.ErrorResponse{})
	return
}
