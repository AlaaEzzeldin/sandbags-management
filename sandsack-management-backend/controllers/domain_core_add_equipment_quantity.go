package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// AddEquipmentQuantity
// @Description This endpoint increase the quantity of chosen equipment
// @Summary This endpoint increase the quantity of chosen equipment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.OrderEquipment true "equipment_id, quantity"
// @Success 200 {array} models.OrderEquipment
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/equipment/return [patch]
func (a *App) AddEquipmentQuantity(c *gin.Context) {
	// TODO: controller name to be clarified with Nastya
	var input models.OrderEquipment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("AddEquipmentQuantity error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	claims, err := GetClaims(c)
	if claims.Role != "Mollnhof" {
		log.Println("Other users cannot add quantity of returned equipment")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "other users cannot add quantity of returned equipment",
		})
		return
	}

	err = service.AddEquipmentQuantity(a.DB, input.EquipmentId, input.Quantity)
	if err != nil {
		log.Println("AddEquipmentQuantity error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	equipment, err := service.GetEquipment(a.DB)
	c.JSON(http.StatusOK, equipment)
	return

}
