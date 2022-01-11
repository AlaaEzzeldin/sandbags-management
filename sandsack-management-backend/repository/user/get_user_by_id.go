package repo_user

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetUserByID(db *gorm.DB, userId int) (user *models.User, err error) {
	query := `select
       u.id, u.name, u.phone, u.password, u.email,
       u.token, u.is_activated, u.is_email_verified,
       u.is_super_user, u.create_date, u.branch_id, b.name as branch_name
		from public.user u, branch b where u.id = ? and b.id = u.branch_id`
	if err = db.Raw(query, userId).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}
