package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// AdminAllOrders
// @Description AdminAllOrders - list of all orders
// @Summary AdminAllOrders - list of all orders
// @Accept json
// @Param Authorization header string true " "
// @Param id path string true "Id of the order"
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /admin/orders/ [get]
func (a *App) AdminAllOrders(c *gin.Context) {
	var input models.GetAllOrders
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: EditOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		return
	}

	if claims.Role != "Einsatzleiter" {
		log.Println("Role is not Einsatzleiter")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Sie sind nicht den Einsatzleiter",
		})
		return
	}

	orders, err := service.GetAllOrders(a.DB, input.StartDate, input.EndDate)
	if err != nil {
		log.Println("GetAllOrders error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusOK, orders)
	return
}
