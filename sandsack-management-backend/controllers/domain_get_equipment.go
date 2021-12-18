package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func (a *App) GetEquipment(c *gin.Context) {
	equipment, err := service.GetEquipment(a.DB)
	if err != nil {
		log.Println("GetEquipment error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, equipment)
}
