package service

import (
	"errors"
	"gorm.io/gorm"
	"math/rand"
	repo_otp "team2/sandsack-management-backend/repository/otp"
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



func SaveOTP(db *gorm.DB, userId int, otp, reason string) error {
	if exist := repo_otp.ExistOTPByUser(db, userId); exist {
		if err := repo_otp.UpdateOTP(db, userId, otp, reason); err != nil {
			return err
		}
	} else {
		if err := repo_otp.InsertOTP(db, userId, otp, reason); err != nil {
			return err
		}
	}
	return nil
}


func GetOTP(db *gorm.DB, otp string, reason string) (string, error){
	if exists := repo_otp.ExistOtpByCode(db, otp, reason); !exists {
		return "", errors.New("does not exist")
	}
	code, err := repo_otp.GetOTP(db, otp, reason)
	if err != nil {
		return "", err
	}
	return code, nil
}


func DeleteOTP(db *gorm.DB, userId int, reason string) error {
	return repo_otp.DeleteOTP(db, userId, reason)
}