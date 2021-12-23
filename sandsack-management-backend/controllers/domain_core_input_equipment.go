package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// CreateEquipment Registration
// @Description Admin can add new type of equipment
// @Summary Admin can add new type of equipment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.OrderEquipment true "Needed only name and quantity of new equipment"
// @Success 200 {array} models.OrderEquipment
// @Failure 505 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/equipment [post]
func (a *App) CreateEquipment(c *gin.Context) {
	var input models.OrderEquipment
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("CreateEquipment error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}


	equipments, err := service.CreateEquipment(a.DB, input.Name, input.Quantity)
	if err != nil {
		log.Println("CreateEquipment error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, equipments)
	return
}
