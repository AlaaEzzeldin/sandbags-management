package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/middleware"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// Login
// @Description Login - user inputs auth credentials and gets tokens
// @Summary Login - user inputs auth credentials and gets tokens
// @Accept json
// @Param input body models.Login true "Login"
// @Success 200 {object} models.Tokens
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Authentication
// @Router /users/login [post]
func (a *App) Login(c *gin.Context) {
	var input models.Login

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Registration error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	email := strings.ToLower(input.Email)

	exists := service.CheckUserExists(a.DB, email)
	if !exists {
		log.Println("Fehler: CheckUserExists: dieser Benutzer gibt es nicht")
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrCode:    http.StatusNotFound,
			ErrMessage: "Benutzer kann nicht gefunden werden",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	// if email is not verified, user cannot be logged in
	if !user.IsEmailVerified {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Email muss bestätigt werden",
		})
		return
	}

	// if user is not activated, not possible to log in
	if !user.IsActivated {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Dieser Benutzer wurde deaktiviert",
		})
		return
	}

	// check password is correct
	ok := functions.CheckPasswordHash(input.Password, user.Password)
	if !ok {
		log.Println("Fehler: CheckPasswordHash")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültiges Passwort",
		})
		return
	}

	tokens, err := middleware.GenerateTokens(a.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	var refreshToken string
	// if we do not have token in database, then put it in
	if len(user.Token) == 0 {
		if err := service.UpdateUserToken(a.DB, email, tokens["refresh_token"]); err != nil {
			log.Println("Fehler: UpdateUserToken", err.Error())
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				ErrCode:    http.StatusInternalServerError,
				ErrMessage: "Da ist etwas schief gelaufen",
			})
			return
		}
		refreshToken = tokens["refresh_token"]
	} else {
		refreshToken = user.Token
	}

	tokens["refresh_token"] = refreshToken

	// return access and refresh tokens
	c.JSON(http.StatusOK, models.Tokens{
		RefreshToken: tokens["refresh_token"],
		AccessToken:  tokens["access_token"],
		Role:         user.BranchName,
		Name:         user.Name,
	})
	return
}
