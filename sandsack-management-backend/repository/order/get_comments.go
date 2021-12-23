package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetComments(a *gorm.DB, orderId int) []models.Comment {
	var comments []models.Comment
	query := `select com.comment_text, u.name, b.name as branch_name 
				from public.comment com, public.user u, branch b 
				where com.order_id = ? 
				and u.id = com.user_id
				and u.branch_id = b.id;`
	err := a.Raw(query, orderId).Scan(&comments).Error
	if err != nil {
		log.Println("GetComments error", err.Error())
		return nil
	}
	return comments
}
