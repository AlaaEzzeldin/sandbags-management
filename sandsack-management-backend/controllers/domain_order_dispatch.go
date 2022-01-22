package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// DispatchOrder
// @Description DispatchOrder - Mollnhof can dispatch the order and assign it to the driver
// @Summary DispatchOrder - Mollnhof can dispatch the order and assign it to the driver
// @Accept json
// @Param Authorization header string true " "
// @Param input body models.DispatchOrderInput true "DispatchOrder"
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/dispatch [post]
func (a *App) DispatchOrder(c *gin.Context) {
	var input models.DispatchOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: DispatchOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	if input.DriverId == 0 {
		input.DriverId = 1
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("Fehler: AcceptOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	if claims.Role != "Mollnhof" {
		log.Println("Die Rolle ist nicht Mollnhof: ", claims.Role)
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Das ist Ihnen nicht erlaubt",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, input.OrderId)
	if err != nil {
		log.Println("Fehler: AcceptOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	flag := 0
	for _, i := range permissions {
		if i == models.DictPermissionName["CAN ACCEPT"] {
			flag = 1
			break
		}
	}

	if flag != 1 {
		log.Println("User cannot accept this order")
		c.JSON(http.StatusConflict, models.ErrorResponse{
			ErrCode:    http.StatusConflict,
			ErrMessage: "Sie können diese Bestellung nicht annehmen",
		})
		return
	}

	err = service.DispatchOrder(a.DB, claims.Id, input.OrderId, input.DriverId)
	if err != nil {
		log.Println("Fehler: DispatchOrder:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	order, err := service.GetOrder(a.DB, claims.Id, input.OrderId)
	c.JSON(http.StatusOK, order)
	return
}
