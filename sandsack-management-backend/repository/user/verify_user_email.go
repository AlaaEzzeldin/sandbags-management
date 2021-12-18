package repo_user

import "gorm.io/gorm"

func VerifyUserEmail(db *gorm.DB, email string, isVerified bool) error {
	query := `update public.user set is_email_verified = ? where email = ?;`
	if err := db.Exec(query, isVerified, email).Error; err != nil {
		return err
	}
	return nil
}
