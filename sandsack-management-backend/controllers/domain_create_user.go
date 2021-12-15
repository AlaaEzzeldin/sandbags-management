package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
// @Router /user/ [post]
func (a *App) CreateUser(c *gin.Context) {
	var input models.CreateUser

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("Registration error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	const bearer = "Bearer "
	header := c.GetHeader("Authorization")
	tokenStr := header[len(bearer):]

	email, err := service.GetEmail(tokenStr)
	if err != nil {
		log.Println("GetEmail error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: err.Error(),
		})
		return
	}

	meUser, err := service.GetUserByEmail(a.DB, email)
	if err != nil {
		log.Println("GetUserByToken error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}
	log.Println("USER", meUser)

	if !meUser.IsSuperUser {
		log.Println("Trying to create user error")
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "you don't have this right",
		})
		return
	}



	parent, err := service.GetUserByID(a.DB, input.ParentId)
	if parent.Name == "Mollnhof" {
		log.Println("Mollnhof cannot have any branches except Einsatzleiter")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "Mollnhof cannot have any branched except Einsatzleiter",
		})
	}

	if err := service.CreateUser(a.DB, &input); err != nil {
		log.Println("CreateUser error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, models.ErrorResponse{})
	return
}
