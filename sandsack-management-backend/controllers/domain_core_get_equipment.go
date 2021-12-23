package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// GetEquipment
// @Description GetEquipment - array of equipment and current quantity of it in Mollnhof
// @Summary GetEquipment - array of equipment and current quantity of it in Mollnhof
// @Accept json
// @Param Authorization header string true " "
// @Success 200 {array} models.OrderEquipment
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Core
// @Router /core/equipment [get]
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
