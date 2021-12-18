package repo_user

import (
	"errors"
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
	repo_otp "team2/sandsack-management-backend/repository/otp"
)

func GetUserByOTP(db *gorm.DB, otp, reason string)  (user *models.User, err error) {
	if exist := repo_otp.ExistOtpByCode(db, otp, reason); !exist {
		return nil, errors.New("user not found")
	}
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date, branch_id 
				from public.user
				where id = (select user_id from otp where code = ? and type = ?);`
	if err = db.Raw(query, otp, reason).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}
