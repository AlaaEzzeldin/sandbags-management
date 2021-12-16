package repo_user

import "gorm.io/gorm"

func UpdateUserActivity(db *gorm.DB, email string, isActivated bool) error {
	query := `update public.user set is_activated = ? where email = ?;`
	if err := db.Exec(query, isActivated, email).Error; err != nil {
		return err
	}
	return nil
}
