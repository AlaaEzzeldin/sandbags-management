package repo_user

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetParent(db *gorm.DB, userId int) (user *models.User, err error) {
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date, branch_id 
				from public.user
			where id = (select user1_id from hierarchy where user2_id = ?);`

	if err = db.Raw(query, userId).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}