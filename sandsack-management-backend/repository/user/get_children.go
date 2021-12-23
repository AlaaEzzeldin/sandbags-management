package repo_user

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetChildren(db *gorm.DB, parentId int) (users *[]models.User, err error) {
	query := `select u.id, u.name, u.email, u.branch_id, b.name as branch_name
				from public.user u, branch b, hierarchy h
				where b.id = u.branch_id
				  and h.user1_id = ?
				  and u.id = h.user2_id;`
	if err := db.Raw(query, parentId).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
