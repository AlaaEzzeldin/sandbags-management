package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// AddPriority
// @Description This endpoint adds new priority
// @Summary This endpoint adds new priority
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.Priority true "level, name"
// @Success 200 {array} models.Priority
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/priority/add [post]
func (a *App) AddPriority(c *gin.Context) {
	var input models.Priority
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("AddPriority error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	err := service.AddPriority(a.DB, input.Name, input.Level)
	if err != nil {
		log.Println("AddPriority error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	priorities, err := service.GetPriority(a.DB)
	c.JSON(http.StatusOK, priorities)
	return
}

// AddEquipment
// @Description This endpoint adds new equipment
// @Summary This endpoint adds new equipment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.OrderEquipment true "name, quantity"
// @Success 200 {array} models.OrderEquipment
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/equipment/add [post]
func (a *App) AddEquipment(c *gin.Context) {
	var input models.OrderEquipment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: AddEquipment: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}

	err := service.AddEquipment(a.DB, input.Name, input.Quantity)
	if err != nil {
		log.Println("Fehler: AddEquipment: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	equipments, err := service.GetEquipment(a.DB)
	c.JSON(http.StatusOK, equipments)
	return
}
