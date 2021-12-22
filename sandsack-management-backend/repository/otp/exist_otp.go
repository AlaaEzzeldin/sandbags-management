package repo_otp

import (
	"gorm.io/gorm"
	"log"
)

func ExistOtpByCode(db *gorm.DB, otp string, reason string) bool {
	query := `select count(*) from otp where code = ? and type = ?;`
	var count int
	if err := db.Raw(query, otp, reason).Scan(&count).Error; err != nil {
		return false
	}
	if count != 0 {
		return true
	}
	return false
}

func ExistOTPByUser(db *gorm.DB, userId int) bool {
	query := `select count(1) from otp where user_id = ?;`
	var count int
	if err := db.Raw(query, userId).Scan(&count).Error; err != nil {
		log.Println("existOTP error", err.Error())
		return false
	}
	if count == 0 {
		return false
	}
	return true
}
