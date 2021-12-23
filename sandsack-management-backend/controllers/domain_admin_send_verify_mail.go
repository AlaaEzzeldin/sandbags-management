package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// SendVerifyEmail
// @Description SendVerifyEmail - admin sends email to user for him to verify
// @Summary SendVerifyEmail - admin sends email to user for him to verify
// @Accept json
// @Param Authorization header string true " "
// @Param input body models.SendVerifyEmail true "SendVerifyEmail"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Admin
// @Router /email_verification [post]
func (a *App) SendVerifyEmail(c *gin.Context) {
	var input models.SendVerifyEmail

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("SendVerifyEmail error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, input.Email)
	if err != nil {
		log.Println("GetUserByEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	otp, err := service.GenerateAndSaveOTP(a.DB, user.Id, "verification")
	if err != nil {
		log.Println("GenerateAndSaveOTP error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if err := functions.SendEmail(a.DB, user.Email, otp, "verification"); err != nil {
		log.Println("SendEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, models.SendVerifyMailOutput{
		OTP: otp,
	})
	return
}
