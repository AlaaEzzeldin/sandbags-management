package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// AcceptOrder
// @Description AcceptOrder - user can accept order
// @Summary AcceptOrder - user can accept order
// @Accept json
// @Param Authorization header string true " "
// @Param id path string true "Id of the order"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/accept [post]
func (a *App) AcceptOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println("Error in parsing", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect input",
		})
		return
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

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, orderId)
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
		// CAN ACCEPT
		if i == 5 {
			flag = 1
			break
		}
	}

	if flag != 1 {
		log.Println("User cannot view this order")
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrCode:    http.StatusNotFound,
			ErrMessage: "you cannot view this order",
		})
		return
	}

	err = service.AcceptOrder(a.DB, claims.Id, orderId)
	if err != nil {
		log.Println("AcceptOrder error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	order, err := service.GetOrder(a.DB, claims.Id, orderId)
	c.JSON(http.StatusOK, order)
	return

}
