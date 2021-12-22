package repo_user

import (
	"gorm.io/gorm"
	"log"
)

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
