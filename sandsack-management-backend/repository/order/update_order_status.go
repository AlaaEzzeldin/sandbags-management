package repo_order

import "gorm.io/gorm"

func UpdateOrderStatus(db *gorm.DB, orderId int, statusId int) error {
	query := `update public.order set status_id = ?, update_date = now() where id = ?;`
	if err := db.Exec(query, statusId, orderId).Error; err != nil {
		return err
	}
	return nil
}
