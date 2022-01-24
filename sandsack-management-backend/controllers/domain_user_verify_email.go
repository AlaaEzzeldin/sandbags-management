package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// VerifyEmail
// @Description VerifyEmail - when user got email with otp to verify email, it has to input this otp to verify email and set new password
// @Summary VerifyEmail - when user got email with otp to verify email, it has to input this otp to verify email and set new password
// @Accept json
// @Param Authorization header string true " "
// @Param input body models.VerifyEmailInput true "VerifyEmail"
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/activation [post]
func (a *App) VerifyEmail(c *gin.Context) {
	var input models.VerifyEmailInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: SendVerifyEmail: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	user, err := service.GetUserByOTP(a.DB, input.Otp, "verification")
	if err != nil {
		log.Println("Fehler: GetUserByOTP: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	otp, err := service.GetOTP(a.DB, input.Otp, "verification")
	if err != nil {
		log.Println("Fehler: GetOTP: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	if otp != input.Otp {
		log.Println("Wrong otp")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "OTP ist ungültig",
		})
		return
	}
	if len(input.Password) < 6 {
		log.Println("Wrong password")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Passwort muss mindestens 6 Symbolen lang sein",
		})
		return
	}

	hashedPassword, err := functions.HashPassword(input.Password)

	if err := service.UpdatePassword(a.DB, user.Email, hashedPassword); err != nil {
		log.Println("UpdatePassword error:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "something went wrong",
		})
		return
	}

	if err := service.UpdateUserActivity(a.DB, user.Email, true); err != nil {
		log.Println("Fehler: UpdateUserActivity:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	if err := service.VerifyUserEmail(a.DB, user.Email, true); err != nil {
		log.Println("Fehler: VerifyUserEmail:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	if err := service.DeleteOTP(a.DB, user.Id, "verification"); err != nil {
		log.Println("Fehler: DeleteOTP:", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
