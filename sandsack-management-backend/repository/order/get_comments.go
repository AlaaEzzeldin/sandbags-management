package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetComments(a *gorm.DB, orderId int) []models.Comment {
	var comments []models.Comment
	query := `select comment_text from public.comment where order_id = ?;`
	err := a.Raw(query, orderId).Scan(&comments).Error
	if err != nil {
		log.Println("GetComments error", err.Error())
		return nil
	}
	return comments
}
