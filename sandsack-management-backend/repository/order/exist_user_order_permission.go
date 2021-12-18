package repo_order

import (
	"gorm.io/gorm"
	"log"
)

func ExistUserOrderPermission(db *gorm.DB, userId, orderId int) bool {
	query := `select count(1) from public.user_order_permission where user_id = ? and order_id = ?;`
	var count int
	if err := db.Raw(query, userId, orderId).Scan(&count).Error; err != nil {
		log.Println("ExistUserOrderPermission error", err.Error())
		return false
	}
	if count == 0 {
		return false
	}
	return true
}
