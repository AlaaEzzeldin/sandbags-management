package repo_user

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetUserList(db *gorm.DB) (userList *[]models.User, err error) {
	query := `select c.id, c.name, c.email, c.branch_id, b.name as branch_name, 
					c.create_date, c.update_date, p.id as parent_id, p.name as parent_name
				from hierarchy h
				join "user" c on h.user2_id = c.id
				join "user" p on h.user1_id = p.id
				join branch b on c.branch_id=b.id
				order by c.branch_id;`
	if err := db.Raw(query).Scan(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}
