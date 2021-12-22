package repo_otp

import "gorm.io/gorm"

func UpdateOTP(db *gorm.DB, userId int, otp, reason string) error {
	query := `update otp set code = ? where user_id = ? and type = ?;`
	if err := db.Exec(query, otp, userId, reason).Error; err != nil {
		return err
	}
	return nil
}

func InsertOTP(db *gorm.DB, userId int, otp, reason string) error {
	query := `insert into otp(code, user_id, type) values(?, ?, ?);`
	if err := db.Exec(query, otp, userId, reason).Error; err != nil {
		return err
	}

	return nil
}