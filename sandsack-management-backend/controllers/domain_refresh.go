package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
	"time"
)


// RefreshAccessToken
// @Description RefreshAccessToken - refreshes access token
// @Summary RefreshAccessToken - refreshes access token
// @Accept json
// @Param input body models.RefreshAccessTokenRequest true "RefreshAccessToken"
// @Success 200 {object} models.Tokens
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/refresh [post]
func (a *App) RefreshAccessToken(c *gin.Context) {
	var input models.RefreshAccessTokenRequest
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("RefreshAccessToken error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse {
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	user, err := service.GetUserByToken(a.DB, input.RefreshToken)
	if err != nil {
		log.Println("GetUserByToken error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse {
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	token, err := service.VerifyToken(user.Token)
	if err != nil {
		log.Println("VerifyToken error: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse {
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "token is not valid",
		})
		return
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok || claims.ExpiresAt < time.Now().Unix() {
		log.Println("VerifyToken error: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse {
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "token is not valid",
		})
		return
	}

	tokens, err := service.GenerateTokens(a.DB, claims.Email)
	if err != nil{
		log.Println("RefreshToken error in GeneratePairToken: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, models.Tokens{
		AccessToken: tokens["access_token"],
		RefreshToken: input.RefreshToken,
	})

}