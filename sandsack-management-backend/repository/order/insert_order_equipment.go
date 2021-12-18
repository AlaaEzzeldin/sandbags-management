package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func InsertOrderEquipment(a *gorm.DB, orderId int, equipments []models.OrderEquipment) error {
	for _, row := range equipments {
		query := `insert into public.order_equipment(equipment_id, quantity, order_id) values(?,?,?);`
		err := a.Exec(query, row.EquipmentId, row.Quantity, orderId).Error
		if err != nil {
			log.Println("InsertOrderEquipment error", err.Error())
			return err
		}
	}
	return nil
}