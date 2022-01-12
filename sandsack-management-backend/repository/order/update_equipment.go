package repo_order

import "gorm.io/gorm"

func UpdateEquipment(db *gorm.DB, equipmentId int, quantity int) error {
	query := `update public.equipment set quantity = ? where id = ?;`
	err := db.Exec(query, quantity, equipmentId).Error
	return err
}
