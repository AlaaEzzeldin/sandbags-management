package repo_user

import "gorm.io/gorm"

func UpdateUserToken(db *gorm.DB, email string, refreshToken string) error {
	query := `update public.user set token = ? where email = ?;`
	if err := db.Exec(query, refreshToken, email).Error; err != nil {
		return err
	}
	return nil
}
