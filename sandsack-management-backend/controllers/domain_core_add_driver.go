package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
)

func (a *App) AddDriver(c *gin.Context) {
	var input models.AddDriverInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("AddDriver error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	

}
