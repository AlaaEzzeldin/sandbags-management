package repo_order

import (
	"gorm.io/gorm"
	"log"
)

func InsertUserOrderPermissions(a *gorm.DB, userId int, orderId int, permissionList []int) error {
	for _, row := range permissionList {
		query := `insert into public.user_order_permission(user_id, order_id, permission_id) values(?,?,?);`
		err := a.Exec(query, userId, orderId, row).Error
		if err != nil {
			log.Println("InsertUserOrderPermissions error", err.Error())
			return err
		}
	}
	return nil
}
