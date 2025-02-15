package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// PatchProfile
// @Description This endpoint enables to change some user profile information
// @Summary Change profile information
// @Accept json
// @Produce json
// @Param Authorization header string true " "
// @Param input body models.PatchProfileInput true "User profile change model"
// @Success 200 "Success message"
// @Failure 401 "Token is not valid"
// @Failure 400 "Bad request (e.g. validation error) OR wrong password given"
// @Failure 500 "Something unexpected went wrong"
// @Tags Authentication
// @Router /users/me [patch]
func (a *App) PatchProfile(c *gin.Context) {
	var input models.PatchProfileInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("PatchProfile error: ", err.Error())
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
			ErrMessage: "Something went wrong",
		})
		return
	}
	/*
	if len(input.Name) == 0 || len(input.Phone) == 0 {
		log.Println("Length of name is ", len(input.Name), ", length of phone is ", len(input.Phone))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "please fill the form",
		})
		return
	}
	*/

	if err := service.PatchProfile(a.DB, claims.Id, input.Name, input.Phone, input.Email); err != nil {
		log.Println("Profile error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	user, err := service.GetUserByID(a.DB, claims.Id)
	user.Password = ""

	c.JSON(http.StatusOK, user)
	return
}

// GetProfile
// @Description GetProfile - get info of the user
// @Summary GetProfile - get info of the user
// @Accept json
// @Param Authorization header string true " "
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/me [get]
func (a *App) GetProfile(c *gin.Context) {
	claims, err := GetClaims(c)
	if err != nil {
		log.Println("Fehler: GetClaims:", err.Error())
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Access Token ist ungültig",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, claims.Email)
	if err != nil {
		log.Println("Fehler: GetUserByID:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	user.Password = ""
	user.Token = ""

	c.JSON(http.StatusOK, user)
	return

}
