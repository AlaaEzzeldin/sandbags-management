package repo_order

import "gorm.io/gorm"

func SetDriverToOrder(db *gorm.DB, orderId, driverId int) error {
	query := `update public.order set driver_id = ? where id = ?;`
	return db.Exec(query, driverId, orderId).Error
}
