package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// DispatchOrder
// @Description DispatchOrder - Mollnhof can dispatch the order and assign it to the driver
// @Summary DispatchOrder - Mollnhof can dispatch the order and assign it to the driver
// @Accept json
// @Param Authorization header string true " "
// @Param id path string true "id of the order"
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/dispatch [post]
func (a *App) DispatchOrder(c *gin.Context) {
	id := c.Query("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("DispatchOrder, Fehler beim Parsen:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültiges Eingabeformat",
		})
		return
	}


	var driverId = 1

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("DispatchOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültiges Eingabeformat",
		})
		return
	}

	if claims.Role != "Mollnhof" {
		log.Println("Die Rolle ist nicht Mollnhof: ", claims.Role)
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode: http.StatusForbidden,
			ErrMessage: "Das ist Ihnen nicht erlaubt",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, orderId)
	if err != nil {
		log.Println("Fehler: DispatchOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
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
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrCode:    http.StatusNotFound,
			ErrMessage: "Sie können diese Bestellung nicht annehmen",
		})
		return
	}

	err = service.DispatchOrder(a.DB, claims.Id, orderId, driverId)
	if err != nil {
		log.Println("DispatchOrder error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}


	order, err := service.GetOrder(a.DB, claims.Id, orderId)
	c.JSON(http.StatusOK, order)
	return
}
