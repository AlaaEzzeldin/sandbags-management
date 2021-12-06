package service

import (
	"errors"
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

func GenerateAndSaveOTP(db *gorm.DB, userId int, reason string) (string, error) {
	otp, err := GenerateOTP(6)
	if err != nil {
		return "", err
	}

	if err := SaveOTP(db, userId, otp, reason); err != nil {
		return "", err
	}
	return otp, nil
}

func existOTPByUser(db *gorm.DB, userId int) bool {
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


func SaveOTP(db *gorm.DB, userId int, otp, reason string) error {
	if exist := existOTPByUser(db, userId); exist {
		query := `update otp set code = ? where user_id = ?;`
		if err := db.Exec(query, otp, userId).Error; err != nil {
			return err
		}
	}

	query := `insert into otp(code, user_id, type) values(?, ?);`
	if err := db.Exec(query, otp, userId, reason).Error; err != nil {
		return err
	}
	return nil
}

func existOtpByCode(db *gorm.DB, otp string, reason string) bool {
	query := `select count(*) from otp where code = ? and reason = ?;`
	var count int
	if err := db.Raw(query, otp, reason).Scan(&count).Error; err != nil {
		return false
	}
	if count != 0 {
		return true
	}
	return false
}

func GetOTP(db *gorm.DB, otp string, reason string) (string, error){
	if exists := existOtpByCode(db, otp, reason); !exists {
		return "", errors.New("does not exist")
	}
	query := `select code from otp where code = ? and type = ?`
	var code string
	if err := db.Raw(query, otp, reason).Scan(&code).Error; err != nil {
		return "", err
	}
	return code, nil
}