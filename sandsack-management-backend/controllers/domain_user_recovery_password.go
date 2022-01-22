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
// @Description RecoveryPassword - when user got OTP per email, he needs to input new password and otp to set password
// @Summary RecoveryPassword - when user got OTP per email, he needs to input new password and otp to set password
// @Accept json
// @Param input body models.RecoveryPasswordInput true "RecoveryPassword"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/recovery_password [post]
func (a *App) RecoveryPassword(c *gin.Context) {
	var input models.RecoveryPasswordInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: SendVerifyEmail: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	_, err := service.GetOTP(a.DB, input.OTP, "recovery")
	if err != nil {
		log.Println("Fehler: GetOTP: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	if len(input.Password) < 6 {
		log.Println("Wrong otp")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Passwort muss mindestens 6 Symbolen lang sein",
		})
		return
	}

	user, err := service.GetUserByOTP(a.DB, input.OTP, "recovery")
	if err != nil {
		log.Println("GetUserByOTP error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	hashedPassword, err := functions.HashPassword(input.Password)

	if err := service.UpdatePassword(a.DB, user.Email, hashedPassword); err != nil {
		log.Println("UpdatePassword error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	if err := service.DeleteOTP(a.DB, user.Id, "recovery"); err != nil {
		log.Println("Fehler: DeleteOTP:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusOK, "Neues Passwort wurde hinterlegt")
	return
}
