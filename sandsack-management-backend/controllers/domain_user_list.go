package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// GetUserList
// @Description GetUserList - get list of all users
// @Summary GetUserList - get list of all users
// @Accept json
// @Param Authorization header string true " "
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/ [get]
func (a *App) GetUserList(c *gin.Context) {
	userList, err := service.GetUserList(a.DB)
	if err != nil {
		log.Println("GetUserList error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, userList)
	return
}
