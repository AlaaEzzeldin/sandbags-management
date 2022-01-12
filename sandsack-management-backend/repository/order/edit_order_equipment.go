package repo_order

import "gorm.io/gorm"

func EditOrderEquipment(db *gorm.DB, orderId, equipmentId, quantity int) error {
	query := `update order_equipment set quantity = ? where order_id = ? and equipment_id = ?;`

	return db.Exec(query, quantity, orderId, equipmentId).Error
}
