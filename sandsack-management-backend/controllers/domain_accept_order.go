package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func (a *App) AcceptOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect input",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("AcceptOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, orderId)
	if err != nil {
		log.Println("AcceptOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
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
			ErrCode: http.StatusNotFound,
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

	c.JSON(http.StatusOK, nil)

}
