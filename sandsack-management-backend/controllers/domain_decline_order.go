package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func (a *App) DeclineOrder(c *gin.Context) {
	var input models.AcceptOrderInput
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("DeclineOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}
	claims, _ := GetClaims(c)

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, input.OrderId)

	flag := 0
	for _, i := range permissions {
	// CAN DECLINE
		if i == 4 {
			flag = 1
			break
		}
	}

	if flag == 0 {
		log.Println("No access for this action error")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "no access",
		})
		return
	}

	err = service.DeclineOrder(a.DB, claims.Id, input.OrderId)
	if err != nil {
		log.Println("DeclineOrder error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, nil)

}
