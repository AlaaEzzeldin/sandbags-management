package repo_order

import "gorm.io/gorm"

func ConfirmDelivery(db *gorm.DB, userId, orderId, statusId int) error {
	query := `update public.order set status_id = ?, update_date = now() where id = ?;`
	return db.Exec(query, statusId, orderId).Error
}
