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
// @Param start_date path string true "start date"
// @Param end_date path string true "end date"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Order
// @Router /order/stats [get]
func (a *App) GetStatistics(c *gin.Context) {
	var input models.GetStatisticsInput
	input.StartDate = c.Query("start_date")
	input.EndDate = c.Query("end_date")
	claims, _ := GetClaims(c)

	if claims.Role == "Einsatzabschnitt" {
		unterabschnittStats, err := service.UnterabschnittStatsForEinsatzabschnitt(a.DB, claims.Id, input.StartDate, input.EndDate)
		if err != nil {
			log.Println("Fehler: GetStatistics: ", err.Error())
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				ErrCode:    http.StatusInternalServerError,
				ErrMessage: "Da ist etwas schief gelaufen",
			})
			return
		}
		statArray := []interface{}{unterabschnittStats}
		c.JSON(http.StatusOK, gin.H{
			"statistics": statArray,
		})
		return
	} else if claims.Role == "Hauptabschnitt" {
		unterabschnittStats, err := service.EinsatzabschnittStatsForHauptabschnitt(a.DB, claims.Id, input.StartDate, input.EndDate)
		if err != nil {
			log.Println("Fehler: GetStatistics: ", err.Error())
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				ErrCode:    http.StatusInternalServerError,
				ErrMessage: "Da ist etwas schief gelaufen",
			})
			return
		}
		statArray := []interface{}{unterabschnittStats}
		c.JSON(http.StatusOK, gin.H{
			"statistics": statArray,
		})
		return
	} else if claims.Role == "Einsatzleiter" {
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


	/*unterabschnittStats, err := service.GetUnterabschnittStatistics(a.DB, input.StartDate, input.EndDate)
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
	}*/

	//statArray := []interface{}{unterabschnittStats, einsatzabschnittStats, hauptabschnittStats}


}
