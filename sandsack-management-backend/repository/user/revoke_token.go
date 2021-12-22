package repo_user

import "gorm.io/gorm"

func RevokeToken(db *gorm.DB, token string) error {
	query := `update public.user set token = null, update_date = now() where token = ?'`
	if err := db.Exec(query, token).Error; err != nil {
		return err
	}
	return nil
}
