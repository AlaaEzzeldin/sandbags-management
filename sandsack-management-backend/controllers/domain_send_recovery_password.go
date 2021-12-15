package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// SendRecoveryPassword
// @Description SendRecoveryPassword - user requests to reset password, when he forgets his password in order to login
// @Summary SendRecoveryPassword - user requests to reset password, when he forgets his password in order to login
// @Accept json
// @Param input body models.SendRecoveryPasswordInput true "SendRecoveryPassword"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/forgot_password [post]
func (a *App) SendRecoveryPassword(c *gin.Context) {
	var input models.SendRecoveryPasswordInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("SendRecoveryPassword error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}


	user, err := service.GetUserByEmail(a.DB, input.Email)
	if err != nil {
		log.Println("GetUserByEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	otp, err := service.GenerateAndSaveOTP(a.DB, user.Id, "recovery")
	if err != nil {
		log.Println("GenerateAndSaveOTP error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if err := functions.SendEmail(a.DB, user.Email, otp, "recovery"); err != nil {
		log.Println("SendEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, models.SendVerifyMailOutput{
		OTP: otp,
	})
	return
}
