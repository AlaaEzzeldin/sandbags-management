package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetLogs(a *gorm.DB, orderId int) []models.Log{
	var logs []models.Log
	query := `select description from public.log where order_id = ?;`
	err := a.Raw(query, orderId).Scan(&logs).Error
	if err != nil {
		log.Println("GetLogs error", err.Error())
		return nil
	}
	return logs
}
