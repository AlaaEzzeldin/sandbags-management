package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"team2/sandsack-management-backend/service"
)

func (a *App) ListOrder(c *gin.Context) {
	claims, err := GetClaims(c)
	if err != nil {
		return
	}

	orders, err := service.GetOrderList(a.DB, claims.Id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, orders)

}
