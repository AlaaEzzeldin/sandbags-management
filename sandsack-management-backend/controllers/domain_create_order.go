package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func (a *App) CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("CreateOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("GetClaims error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, claims.Email)
	if user.BranchName != "Unterabschnitt" {
		log.Println("GetClaims error: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "you are not allowed to create order",
		})
		return
	}

	order := &models.Order{
		Name: user.Name,
		UserId: user.Id,
		AddressTo: input.AddressTo,
		AddressFrom: "Mollnhof",
		StatusId: models.DictStatusName["PENDING"],
		PriorityId: models.DictPriorityName["HIGH"],
		Comments: input.Comments,
		Equipments: input.Equipments,
	}

	if err := service.CreateOrder(a.DB, user.Name, order); err != nil {
		log.Println("CreateOrder error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, order)

}
