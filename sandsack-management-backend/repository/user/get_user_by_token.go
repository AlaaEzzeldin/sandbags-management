package repo_user

import (
	"errors"
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetUserByRefreshToken(db *gorm.DB, token string) (user models.User, err error) {
	if len(token) == 0 {
		return user, errors.New("token is empty")
	}
	query := `select id, name, phone, password, email, token, is_activated, create_date 
				from public.user
				where token = ?;`

	if err = db.Raw(query, token).Scan(&user).Error; err != nil {
		return user, err
	}

	return
}
