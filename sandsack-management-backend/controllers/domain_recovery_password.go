package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// RecoveryPassword
// @Description This endpoint enables to set new password after resetting
// @Summary Set new password after reset
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.RecoveryPasswordInput true "User recovery password model"
// @Success 200 "Success message"
// @Failure 401 "Token is not valid"
// @Failure 400 "Bad request (e.g. parameter in body is not given or incorrect)"
// @Failure 500 "Something went wrong"
// @Tags Authentication
// @Router /users/recovery_password [post]
func (a *App) RecoveryPassword(c *gin.Context) {
	var input models.RecoveryPasswordInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("SendVerifyEmail error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	_, err := service.GetOTP(a.DB, input.OTP, "recovery")
	if err != nil {
		log.Println("GetOTP error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if len(input.Password) < 6 {
		log.Println("Wrong otp")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "password should be longer 6 symbols",
		})
		return
	}

	user , err := service.GetUserByOTP(a.DB, input.OTP, "recovery")
	if err != nil {
		log.Println("GetUserByOTP error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	hashedPassword, err := functions.HashPassword(input.Password)

	if err := service.UpdatePassword(a.DB, user.Email, hashedPassword); err != nil {
		log.Println("UpdatePassword error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if err := service.DeleteOTP(a.DB, user.Id, "recovery"); err != nil {
		log.Println("DeleteOTP error:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, "The new password has been set")
	return
}
