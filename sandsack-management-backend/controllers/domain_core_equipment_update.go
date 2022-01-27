package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// UpdateEquipment
// @Description This endpoint change quantity of equipment. It is used only by Einsatzleiter
// @Summary This endpoint change quantity of equipment. It is used only by Einsatzleiter
// @Accept json
// @Produce json
// @Param Authorization header string true " "
// @Param input body models.UpdateEquipmentInput true "id, quantity"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/equipment/update [patch]
func (a *App) UpdateEquipment(c *gin.Context) {
	var input models.UpdateEquipmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: UpdateEquipment: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("Fehler: GetClaims:", err.Error())
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	if claims.Role != "Einsatzleiter" {
		log.Println("Die Rolle ist kein Einsatzleiter")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Das ist Ihnen nicht erlaubt. Kein Zugang",
		})
		return
	}

	err = service.UpdateEquipment(a.DB, input.Id, input.Quantity)
	if err != nil {
		log.Println("Fehler: UpdateEquipment", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
