package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// Logout
// @Description Logout an authenticated user
// @Summary Logout an authenticated user
// @Accept json
// @Param Authorization header string true " "
// @Param input body models.Logout true "Logout"
// @Success 204 "Logged out successfully"
// @Failure 401 "Access token is missing"
// @Failure 400 "Bad request (e.g. refresh in body is not given)"
// @Tags Authentication
// @Router /users/logout [post]
func (a *App) Logout(c *gin.Context) {
	var refresh models.Logout

	if err := c.ShouldBindJSON(&refresh); err != nil {
		log.Println("Logout error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Refresh Token ist ungültig oder nicht angegeben",
		})
		return
	}

	if err := service.RevokeToken(a.DB, refresh.Token); err != nil {
		log.Println("Fehler: Logout: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusNoContent, "Erfolgreich ausgeloggt")
	return
}
