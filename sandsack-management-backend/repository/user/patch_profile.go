package repo_user

import "gorm.io/gorm"

func PatchProfile(db *gorm.DB, userId int, name string, phone string, email string) error {
	if len(name) > 0 {
		query := `update public.user set name = ? where id = ?`
		if err := db.Exec(query, name, userId).Error; err != nil {
			return err
		}
	}
	if len(phone) > 0 {
		query := `update public.user set phone = ? where id = ?`
		if err := db.Exec(query, phone, userId).Error; err != nil {
			return err
		}
	}
	if len(email) > 0 {
		query := `update public.user set email = ? where id = ?;`
		if err := db.Exec(query, email, userId).Error; err != nil {
			return err
		}
	}
	return nil
}