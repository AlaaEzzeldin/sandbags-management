package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/middleware"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// CreateUser Registration
// @Description This endpoint is implemented to register new user by Einsatzleiter and get a new token pair
// @Summary Create a new user (branch) in the system
// @Accept json
// @Produce json
// @Param input body models.CreateUser true "User registration model"
// @Success 201 "User has been created"
// @Failure 401 "Permission to create the user is not given"
// @Failure 400 "Bad request (e.g. parameter in body is not given or incorrect)"
// @Tags Admin
// @Router /create_user/ [post]
func (a *App) CreateUser(c *gin.Context) {
	var input models.CreateUser

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: Registration: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}

	const bearer = "Bearer "
	header := c.GetHeader("Authorization")
	tokenStr := header[len(bearer):]

	email, err := middleware.GetEmail(tokenStr)
	if err != nil {
		log.Println("Fehler: GetEmail: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: err.Error(),
		})
		return
	}

	meUser, err := service.GetUserByEmail(a.DB, email)
	if err != nil {
		log.Println("Fehler: GetUserByToken: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}
	log.Println("USER", meUser)

	if !meUser.IsSuperUser {
		log.Println("Trying to create user error")
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Das ist Ihnen nicht erlaubt",
		})
		return
	}

	parent, err := service.GetUserByID(a.DB, input.ParentId)
	if parent.Name == "Mollnhof" {
		log.Println("Mollnhof cannot have any branches except Einsatzleiter")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Mollnhof cannot have any branches except Einsatzleiter",
		})
		return
	}

	if err := service.CreateUser(a.DB, &input); err != nil {
		log.Println("Fehler: CreateUser: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusCreated, models.ErrorResponse{})
	return
}
