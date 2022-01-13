package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// AddDriver
// @Description This endpoint adds new driver
// @Summary This endpoint adds new driver
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.AddDriverInput true "name, description"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/driver/add [post]
func (a *App) AddDriver(c *gin.Context) {
	var input models.AddDriverInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("AddDriver error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("GetClaims error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if claims.Role != "Einsatzleiter" {
		log.Println("Role is not Einsatzleiter")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode: http.StatusForbidden,
			ErrMessage: "you do not have such permissions",
		})
		return
	}

	if err := service.AddDriver(a.DB, input.Name, input.Description); err != nil {
		log.Println("AddDriver error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
