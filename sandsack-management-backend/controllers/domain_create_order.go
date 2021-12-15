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
// @Param input body models.CreateOrderInput true "CreateOrder"
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /users/order [post]
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
	if err != nil {
		log.Println("GetUserByEmail error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}
	log.Println("User branch", user.BranchId)
	if user.BranchId != 5 {
		log.Println("It's not Unterabschnitt")
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

	updatedOrder, err := service.GetOrder(a.DB, user.Id, order.Id)
	if err != nil {
		log.Println("GetOrder error", err.Error())
	}

	c.JSON(http.StatusOK, updatedOrder)

}
