package service

import (
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

var letters = "0123456789qwertzuiopasdfghjklyxcvbnm"
var numbers = "0123456789"
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateOTP(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		b[i] = numbers[seededRand.Intn(len(numbers))]
	}
	return string(b), nil
}

func GenerateAndSaveOTP(db *gorm.DB, userId int) (string, error) {
	otp, err := GenerateOTP(6)
	if err != nil {
		return "", err
	}

	if err := SaveOTP(db, userId, otp); err != nil {
		return "", err
	}
	return otp, nil
}

func existOTP(db *gorm.DB, userId int) bool {
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


func SaveOTP(db *gorm.DB, userId int, otp string) error {
	if exist := existOTP(db, userId); exist {
		query := `update otp set code = ? where user_id = ?;`
		if err := db.Exec(query, otp, userId).Error; err != nil {
			return err
		}
	}

	query := `insert into otp(code, user_id) values(?, ?);`
	if err := db.Exec(query, otp, userId).Error; err != nil {
		return err
	}
	return nil
}

func GetOTP(db *gorm.DB, email string) (string, error){
	query := `select code from otp where user_id = (select id from user where email = ?)`
	var code string
	if err := db.Raw(query, email).Scan(&code).Error; err != nil {
		return "", err
	}
	return code, nil
}