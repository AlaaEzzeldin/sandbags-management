package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func InsertComments(a *gorm.DB, userId, orderId int, comments []models.Comment) error {
	for _, row := range comments {
		query := `insert into public.comment(order_id, comment_text, user_id) values(?,?,?);`
		err := a.Exec(query, orderId, row.CommentText, userId).Error
		if err != nil {
			log.Println("InsertComments error", err.Error())
			return err
		}
	}
	return nil
}
