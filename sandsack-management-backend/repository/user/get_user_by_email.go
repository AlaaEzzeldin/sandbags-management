package repo_user

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetUserByEmail(db *gorm.DB, email string) (user *models.User, err error) {
	query := `select u.id, u.name, u.phone, u.password, u.email, u.token, u.is_activated, u.is_email_verified, u.is_super_user, u.create_date, b.name as branch_name, b.id as branch_id
				from public.user u, branch b
				where u.email = ?
				and u.branch_id = b.id limit 1;`
	if err = db.Raw(query, email).Scan(&user).Error; err != nil {
		return nil, err
	}
	return
}
