package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// CreateOrder
// @Description CreateOrder - Unterabschnitt creates the order
// @Summary CreateOrder - Unterabschnitt creates the order
// @Accept json
// @Param Authorization header string true "Bearer "
// @Param input body models.CreateOrderInput true "CreateOrder"
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/ [post]
func (a *App) CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("CreateOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("GetClaims error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, claims.Email)
	if err != nil {
		log.Println("GetUserByEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}
	log.Println("User branch", user.BranchId)
	if user.BranchId != models.DictBranchName["Unterabschnitt"] {
		log.Println("It's not Unterabschnitt")
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "you are not allowed to create order",
		})
		return
	}

	var comments []models.Comment
	comments = append(comments, models.Comment{CommentText: input.Comment})

	order := &models.Order{
		Name:        user.Name,
		UserId:      user.Id,
		AddressTo:   input.AddressTo,
		AddressFrom: "Mollnhof",
		StatusId:    models.DictStatusName["ANSTEHEND"],
		PriorityId:  models.DictPriorityName["HIGH"],
		Comments:    comments,
		Equipments:  input.Equipments,
	}

	if err := service.CreateOrder(a.DB, user.Name, order); err != nil {
		log.Println("CreateOrder error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	updatedOrder, err := service.GetOrder(a.DB, user.Id, order.Id)
	if err != nil {
		log.Println("GetOrder error", err.Error())
	}

	c.JSON(http.StatusOK, updatedOrder)

}
