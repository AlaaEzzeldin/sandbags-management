package repo_user

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

type Id struct {
	Id int `gorm:"column:id"`
}

func InsertUser(db *gorm.DB, user *models.CreateUser, hashedPassword string) (int, error) {
	var model Id
	query := `insert into public.user(name, email, phone, password) values(?,?,?,?) returning id;`
	if err := db.Raw(query, user.Name, user.Email, user.Phone, hashedPassword).Scan(&model).Error; err != nil {
		return 0, err
	}
	return model.Id, nil
}
