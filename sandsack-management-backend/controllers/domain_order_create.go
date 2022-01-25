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
// @Param Authorization header string true " "
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
		log.Println("Fehler: CreateOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("Fehler: GetClaims: ", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Access Token ist ungültig",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, claims.Email)
	if err != nil {
		log.Println("Fehler: GetUserByEmail: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}
	log.Println("User branch", user.BranchId)

	//if !(user.BranchId == models.DictBranchName["Unterabschnitt"] || user.BranchId == models.DictBranchName["Einsatzabschnitt"])  {
	if user.BranchId != models.DictBranchName["Unterabschnitt"] {
		log.Println("It's not Unterabschnitt")
		c.JSON(http.StatusConflict, models.ErrorResponse{
			ErrCode:    http.StatusConflict,
			ErrMessage: "Sie können keine Bestellung anlegen",
		})
		return
	}

	var comments []models.Comment
	if len(input.Comment) > 0 {
		comments = append(comments, models.Comment{CommentText: input.Comment})
	}

	order := &models.Order{
		Name:        user.Name,
		UserId:      user.Id,
		AddressTo:   input.AddressTo,
		AddressFrom: "Mollnhof",
		StatusId:    models.DictStatusName["ANSTEHEND"],
		PriorityId:  models.DictPriorityName["HOCH"],
		Comments:    comments,
		Equipments:  input.Equipments,
	}

	if err := service.CreateOrder(a.DB, user.Name, order); err != nil {
		log.Println("Fehler: CreateOrder: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	updatedOrder, err := service.GetOrder(a.DB, user.Id, order.Id)
	if err != nil {
		log.Println("Fehler: GetOrder", err.Error())
	}

	c.JSON(http.StatusOK, updatedOrder)

}
