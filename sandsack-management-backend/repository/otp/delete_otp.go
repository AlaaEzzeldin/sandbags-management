package repo_otp

import "gorm.io/gorm"

func DeleteOTP(db *gorm.DB, userId int, reason string) error {
	query := `delete from otp where user_id = ? and type = ?;`
	if err := db.Exec(query, userId, reason).Error; err != nil {
		return err
	}
	return nil
}
