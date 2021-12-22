package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func InsertLogs(a *gorm.DB, logs []models.Log) error {
	for _, row := range logs {
		query := `insert into public.log(action_type_id, description, order_id, updated_by)values(?,?,?,?);`
		err := a.Exec(query, row.ActionTypeId, row.Description, row.OrderId, row.UpdatedBy).Error
		if err != nil {
			log.Println("InsertLogs error", err.Error())
			return err
		}
	}
	return nil
}
