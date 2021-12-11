package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
)

func GetUserByEmail(db *gorm.DB, email string) (user *models.User, err error) {
	if exist := CheckUserExists(db, email); !exist {
		return nil, errors.New("user not found")
	}
	query := `select u.id, u.name, u.phone, u.password, u.email, u.token, u.is_activated, u.is_email_verified, u.is_super_user, u.create_date, b.name as branch_name, b.id as branch_id
				from public.user u, branch b
				where u.email = ?
				and u.branch_id = b.id;`
	if err = db.Raw(query, email).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}

func GetUserByToken(db *gorm.DB, token string) (user models.User, err error) {
	if len(token) == 0 {
		return user, errors.New("token is empty")
	}
	query := `select id, name, phone, password, email, token, is_activated, create_date 
				from public.user
				where token = ?;`

	if err = db.Raw(query, token).Scan(&user).Error; err != nil {
		return user, err
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

type Id struct {
	Id int `gorm:"column:id"`
}

func CreateUser(db *gorm.DB, user *models.CreateUser) error {
	var model Id
	hashedPassword, err := functions.HashPassword(user.Password)
	if err != nil {
		return err
	}
	query := `insert into public.user(name, email, phone, password) values(?,?,?,?) returning id;`
	if err := db.Raw(query, user.Name, user.Email, user.Phone, hashedPassword).Scan(&model).Error; err != nil {
		return err
	}

	if err := AddHierarchy(db, user.ParentId, model.Id); err != nil {
		return err
	}
	return nil
}

func AddHierarchy(db *gorm.DB, parentId int, childId int) error {
	query := `insert into hierarchy(user1_id, user2_id) values(?,?);`
	if err := db.Exec(query, parentId, childId).Error; err != nil {
		return err
	}
	return nil
}


func GetUserList(db *gorm.DB) (userList *[]models.User, err error) {
	query := `select id, name, email from public.user where is_super_user=false;`
	if err := db.Raw(query).Scan(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func UpdatePassword(db *gorm.DB, email, password string) error {
	query := `update public.user set password = ? where email = ?;`
	if err := db.Exec(query, password, email).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserActivity(db *gorm.DB, email string, isActivated bool) error {
	query := `update public.user set is_activated = ? where email = ?;`
	if err := db.Exec(query, isActivated, email).Error; err != nil {
		return err
	}
	return nil
}

func VerifyUserEmail(db *gorm.DB, email string, isVerified bool) error {
	query := `update public.user set is_email_verified = ? where email = ?;`
	if err := db.Exec(query, isVerified, email).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByOTP(db *gorm.DB, otp, reason string)  (user *models.User, err error) {
	if exist := existOtpByCode(db, otp, reason); !exist {
		return nil, errors.New("user not found")
	}
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date 
				from public.user
				where id = (select user_id from otp where code = ? and type = ?);`
	if err = db.Raw(query, otp, reason).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}

func GetParent(db *gorm.DB, userId int) (user *models.User, err error) {
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date 
				from public.user
			where id = (select user1_id from hierarchy where user2_id = ?);`

	if err = db.Raw(query, userId).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}
