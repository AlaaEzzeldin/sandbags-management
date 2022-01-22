package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	repo_order "team2/sandsack-management-backend/repository/order"
	"team2/sandsack-management-backend/service"
)

// EditOrder
// @Description This endpoint edits priority and/or quantity of equipment of the order
// @Summary This endpoint edits priority and/or quantity of equipment of the order
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.EditOrderInput true "EditOrder"
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Order
// @Router /order/edit [patch]
func (a *App) EditOrder(c *gin.Context) {
	var input models.EditOrderInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: EditOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}
	log.Println("Input in editOrder", input)
	claims, _ := GetClaims(c)
	user, _ := service.GetUserByID(a.DB, claims.Id)
	order, _ := service.GetOrder(a.DB, user.Id, input.OrderId)
	var logs []models.Log

	if len(input.Equipments) != 0 {
		for _, row := range input.Equipments {
			err := service.EditOrderEquipment(a.DB, input.OrderId, row.EquipmentId, row.Quantity)
			if err != nil {
				log.Println("Fehler: EditOrderEquipment", err.Error())
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					ErrCode:    http.StatusInternalServerError,
					ErrMessage: "Da ist etwas schief gelaufen",
				})
				return
			}
		}

		log := models.Log{
			OrderId:      input.OrderId,
			ActionTypeId: models.DictActionTypeName["EDITED"],
			Description:  user.Name + " hat die Bestellausrüstungsanzahl " + order.Name + "geändert  #" + strconv.Itoa(order.Id),
			UpdatedBy:    user.Id,
		}

		logs = append(logs, log)
	}

	if input.Priority != 0 {
		err := service.EditOrderPriority(a.DB, input.OrderId, input.Priority)
		if err != nil {
			log.Println("Fehler: EditOrderPriority", err.Error())
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				ErrCode:    http.StatusInternalServerError,
				ErrMessage: "Da ist etwas schief gelaufen",
			})
			return
		}

		log := models.Log{
			OrderId:      input.OrderId,
			ActionTypeId: models.DictActionTypeName["EDITED"],
			Description:  user.Name + " hat die Priorität der Bestellung " + order.Name + "geändert #" + strconv.Itoa(order.Id),
			UpdatedBy:    user.Id,
		}

		logs = append(logs, log)
	}

	_ = repo_order.InsertLogs(a.DB, logs)

	order, _ = service.GetOrder(a.DB, claims.Id, input.OrderId)

	c.JSON(http.StatusOK, order)
	return
}
