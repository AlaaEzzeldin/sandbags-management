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
		log.Println("ChangePassword error: ", err.Error())
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

	if len(input.Name) == 0 || len(input.Phone) == 0 {
		log.Println("Length of name is ", len(input.Name), ", length of phone is ", len(input.Phone))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "please fill the form",
		})
		return
	}

	if err := service.PatchProfile(a.DB, claims.Id, input.Name, input.Phone); err != nil {
		log.Println("Profile error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, "The profile data has been changed successfully")
	return
}


func (a *App) GetProfile(c *gin.Context) {

}