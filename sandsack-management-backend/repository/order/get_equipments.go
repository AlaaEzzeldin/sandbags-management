package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetEquipments(a *gorm.DB, orderId int) []models.OrderEquipment {
	var equipments []models.OrderEquipment
	query := `select e.id as equipment_id, e.name, oe.quantity 
				from public.order_equipment oe, public.equipment e 
				where oe.order_id = ? and e.id = oe.equipment_id;`
	err := a.Raw(query, orderId).Scan(&equipments).Error
	if err != nil {
		log.Println("GetEquipment error", err.Error())
		return nil
	}
	return equipments
}
