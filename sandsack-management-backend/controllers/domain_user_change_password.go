package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// ChangePassword
// @Description This endpoint enables to set new password for the user
// @Summary Change password of an authenticated user
// @Accept json
// @Produce json
// @Param Authorization header string true " "
// @Param input body models.ChangePasswordInput true "User change password model"
// @Success 200 "Success message"
// @Failure 401 "Token is not valid"
// @Failure 400 "Bad request (e.g. validation error) OR wrong password given"
// @Failure 500 "Something unexpected went wrong"
// @Tags Authentication
// @Router /users/change_password [patch]
func (a *App) ChangePassword(c *gin.Context) {
	var input models.ChangePasswordInput

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
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, claims.Email)
	if err != nil {
		log.Println("GetUserByEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if ok := functions.CheckPasswordHash(input.OldPassword, user.Password); !ok {
		log.Println("CheckPasswordHash error")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "wrong password",
		})
		return
	}

	if len(input.NewPassword) < 6 {
		log.Println("NewPassword length: less than 6 symbols ")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "new password should be longer than 6 symbols",
		})
		return
	}

	hashedPassword, err := functions.HashPassword(input.NewPassword)
	if err != nil {
		log.Println("HashPassword error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if err := service.UpdatePassword(a.DB, claims.Email, hashedPassword); err != nil {
		log.Println("UpdatePassword error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, "The password has been changed successfully")
	return

}
