package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// ConfirmDelivery
// @Description Unterabschnitt confirms that equipment is successfully delivered
// @Summary Unterabschnitt confirms that equipment is successfully delivered
// @Accept json
// @Produce json
// @Param Authorization header string true " "
// @Param id path string true "id of the order"
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Order
// @Router /order/delivery/confirm/ [post]
func (a *App) ConfirmDelivery(c *gin.Context) {
	id := c.Query("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Fehler beim Parsen", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}

	claims, _ := GetClaims(c)
	if claims.Role != "Unterabschnitt" {
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Nur der Unterabschnitt darf die Zustellung bestätigen",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, orderId)

	flag := 0
	for _, i := range permissions {
		if i == models.DictPermissionName["CAN CONFIRM DELIVERY"] {
			flag = 1
			break
		}
	}

	if flag == 0 {
		log.Println("Fehler: Kein Zugang")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Kein Zugang",
		})
		return
	}

	if err := service.ConfirmDelivery(a.DB, claims.Id, orderId); err != nil {
		log.Println("ConfirmDelivery Fehler: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	order, err := service.GetOrder(a.DB, claims.Id, orderId)
	if err != nil {
		log.Println("Fehler: GetOrder: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusOK, order)
	return
}
