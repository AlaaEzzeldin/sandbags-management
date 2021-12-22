package repo_user

import "gorm.io/gorm"

func UpdatePassword(db *gorm.DB, email, password string) error {
	query := `update public.user set password = ? where email = ?;`
	if err := db.Exec(query, password, email).Error; err != nil {
		return err
	}
	return nil
}
