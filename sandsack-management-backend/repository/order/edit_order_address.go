package repo_order

import "gorm.io/gorm"

func EditOrderAddress(db *gorm.DB, orderId int, addressTo string) error {
	query := `update public.order set address_to = ? where id = ?;`
	return db.Exec(query, orderId, addressTo).Error
}
