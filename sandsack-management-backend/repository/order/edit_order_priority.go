package repo_order

import "gorm.io/gorm"

func EditOrderPriority(db *gorm.DB, orderId, priority int) error {
	query := `update public.order set priority_id = ? where id = ?;`
	return db.Exec(query, priority, orderId).Error
}
