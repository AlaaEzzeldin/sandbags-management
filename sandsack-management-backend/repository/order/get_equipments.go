package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetEquipments(a *gorm.DB, orderId int) []models.OrderEquipment {
	var equipments []models.OrderEquipment
	query := `select e.id as equipment_id, e.name, oe.quantity, oe.create_date, oe.update_date
				from public.order_equipment oe, public.equipment e 
				where oe.order_id = ? and e.id = oe.equipment_id;`
	err := a.Raw(query, orderId).Scan(&equipments).Error
	if err != nil {
		log.Println("GetEquipment error", err.Error())
		return nil
	}
	return equipments
}

func GetEquipment(a *gorm.DB) ([]models.OrderEquipment, error) {
	var equipments []models.OrderEquipment
	query := `select id as equipment_id, name, quantity, measure
				from public.equipment;`
	err := a.Raw(query).Scan(&equipments).Error
	if err != nil {
		log.Println("GetEquipment error", err.Error())
		return nil, err
	}
	return equipments, nil
}


func GetEquipmentById(a *gorm.DB, equipmentId int) (models.OrderEquipment, error) {
	var equipments models.OrderEquipment
	query := `select id as equipment_id, name, quantity, measure
				from public.equipment where id = ?;`
	err := a.Raw(query, equipmentId).Scan(&equipments).Error
	if err != nil {
		log.Println("GetEquipment error", err.Error())
		return models.OrderEquipment{}, err
	}
	return equipments, nil
}