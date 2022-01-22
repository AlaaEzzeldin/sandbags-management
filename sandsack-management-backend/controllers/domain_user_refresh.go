package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/middleware"
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
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: RefreshAccessToken: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	user, err := service.GetUserByToken(a.DB, input.RefreshToken)
	if err != nil {
		log.Println("Fehler: GetUserByToken: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	token, err := middleware.VerifyToken(user.Token)
	if err != nil {
		log.Println("Fehler: VerifyToken: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Token ist ungültig",
		})
		return
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok || claims.ExpiresAt < time.Now().Unix() {
		log.Println("Fehler: VerifyToken: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Token ist ungültig",
		})
		return
	}

	tokens, err := middleware.GenerateTokens(a.DB, claims.Email)
	if err != nil {
		log.Println("Fehler: RefreshToken in GeneratePairToken: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusOK, models.Tokens{
		AccessToken:  tokens["access_token"],
		RefreshToken: input.RefreshToken,
		Role:         user.BranchName,
	})

}
