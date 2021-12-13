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
// @Description ChangePassword of the user
// @Summary ChangePassword of the user
// @Accept json
// @Param input body models.ChangePasswordInput true "ChangePassword"
// @Success 200 "Password was changed successfully"
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/password [put]
func (a *App) ChangePassword(c *gin.Context) {
	var input models.ChangePasswordInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("ChangePassword error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("GetClaims error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, claims.Email)
	if err != nil {
		log.Println("GetUserByEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if ok := functions.CheckPasswordHash(input.OldPassword, user.Password); !ok {
		log.Println("CheckPasswordHash error")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "wrong password",
		})
		return
	}

	if len(input.NewPassword) < 6 {
		log.Println("NewPassword length: less than 6 symbols ")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "new password should be longer than 6 symbols",
		})
		return
	}

	hashedPassword, err := functions.HashPassword(input.NewPassword)
	if err != nil {
		log.Println("HashPassword error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	if err := service.UpdatePassword(a.DB, claims.Email, hashedPassword); err != nil {
		log.Println("UpdatePassword error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, "Password was changed successfully")
	return

}
