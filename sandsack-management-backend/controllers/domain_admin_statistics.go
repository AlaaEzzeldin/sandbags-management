package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// GetStatistics
// @Description Gets list of stats of orders
// @Summary
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer "
// @Param input body models.GetStatisticsInput true "GetStats Input"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Order
// @Router /order/stats [get]
func (a *App) GetStatistics(c *gin.Context) {
	var input models.GetStatisticsInput
	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: GetStatistics: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage oder Eingabeformat",
		})
		return
	}

	unterabschnittStats, err := service.GetUnterabschnittStatistics(a.DB, input.StartDate, input.EndDate)
	if err != nil {
		log.Println("Fehler: GetStatistics: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}
	einsatzabschnittStats, err := service.GetEinsatzabschnittStatistics(a.DB, input.StartDate, input.EndDate)
	if err != nil {
		log.Println("Fehler: GetStatistics: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	hauptabschnittStats, err := service.GetHauptabschnittStatistics(a.DB, input.StartDate, input.EndDate)
	if err != nil {
		log.Println("Fehler: GetStatistics: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}

	statArray := []interface{}{unterabschnittStats, einsatzabschnittStats, hauptabschnittStats}
	c.JSON(http.StatusOK, gin.H{
		"statistics": statArray,
	})
	return

}
