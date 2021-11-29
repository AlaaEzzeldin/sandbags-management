package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetUserByEmail(db *gorm.DB, email string) (user *models.User, err error) {
	query := `select id, name, phone, password, email, token, is_activated, create_date 
				from public.user
				where email = ?;`

	if err = db.Raw(query, email).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}

func GetUserByToken(db *gorm.DB, token string) (user *models.User, err error) {
	if len(token) == 0 {
		return user, errors.New("token is empty")
	}
	query := `select id, name, phone, password, email, token, is_activated, create_date 
				from public.user
				where token = ?;`

	if err = db.Raw(query, token).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}

func CheckUserExists(db *gorm.DB, email string) bool {
	query := `select count(1) from public.user where email = ?;`
	var count int
	if err := db.Raw(query, email).Scan(&count).Error; err != nil {
		log.Println("CheckUserExists error", err.Error())
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func UpdateUserToken(db *gorm.DB, email string, refreshToken string) error {
	query := `update public.user set token = ? where email = ?;`
	if err := db.Exec(query, refreshToken, email).Error; err != nil {
		return err
	}
	return nil
}

