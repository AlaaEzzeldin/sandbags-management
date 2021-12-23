package repo_order

import "gorm.io/gorm"

func AddEquipmentQuantity(a *gorm.DB, equipmentId int, quantity int) error {
	query := `update public.equipment set quantity = quantity + ? where id = ?;`
	return a.Exec(query, quantity, equipmentId).Error
}
