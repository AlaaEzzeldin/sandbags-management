package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func (a *App) DeclineOrder(c *gin.Context) {
	id := c.Query("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in parsing", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect input",
		})
		return
	}

	claims, _ := GetClaims(c)

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, orderId)

	flag := 0
	for _, i := range permissions {
		if i == models.DictPermissionName["CAN DECLINE"]{
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

	err = service.DeclineOrder(a.DB, claims.Id, orderId)
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
