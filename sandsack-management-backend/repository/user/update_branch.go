package repo_user

import "gorm.io/gorm"

func UpdateBranch(db *gorm.DB, userId, branchId int) error  {
	query := `update public.user set branch_id = ?, update_date = now() where id = ?;`
	if err := db.Exec(query, branchId, userId).Error; err != nil {
		return err
	}
	return nil
}