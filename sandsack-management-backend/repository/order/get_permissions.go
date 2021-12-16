package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetPermissions(a *gorm.DB, userId, orderId int) []models.Permission {
	var permissions []models.Permission
	query := `select p.name from public.permission p, public.user_order_permission uop 
				where uop.user_id = ? and uop.order_id = ?
				and uop.permission_id = p.id;`
	err := a.Raw(query, userId, orderId).Scan(&permissions).Error
	if err != nil {
		log.Println("GetPermissions error", err.Error())
		return nil
	}
	return permissions
}
