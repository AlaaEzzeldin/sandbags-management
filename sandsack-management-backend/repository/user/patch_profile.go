package repo_user

import "gorm.io/gorm"

func PatchProfile(db *gorm.DB, userId int, name string, phone string) error {
	query := `update public.user set name = ?, phone = ? where id = ?`
	if err := db.Exec(query, name, phone, userId).Error; err != nil {
		return err
	}
	return nil
}