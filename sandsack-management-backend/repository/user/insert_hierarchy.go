package repo_user

import "gorm.io/gorm"

func InsertHierarchy(db *gorm.DB, parentId, childId int) error {
	query := `insert into hierarchy(user1_id, user2_id) values(?,?);`
	if err := db.Exec(query, parentId, childId).Error; err != nil {
		return err
	}
	return nil
}
