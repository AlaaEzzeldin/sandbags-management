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
		log.Println("DispatchOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	if input.DriverId == 0 {
		input.DriverId = 1
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("AcceptOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	if claims.Role != "Mollnhof" {
		log.Println("ROle is not Mollnhof: ", claims.Role)
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode: http.StatusForbidden,
			ErrMessage: "you do not have a permission to do it",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, input.OrderId)
	if err != nil {
		log.Println("AcceptOrder error: ", err.Error())
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
			ErrMessage: "you cannot view this order",
		})
		return
	}

	err = service.DispatchOrder(a.DB, claims.Id, input.OrderId, input.DriverId)
	if err != nil {
		log.Println("DispatchOrder error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}


	order, err := service.GetOrder(a.DB, claims.Id, input.OrderId)
	c.JSON(http.StatusOK, order)
	return
}
