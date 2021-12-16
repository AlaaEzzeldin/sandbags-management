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
				and u.branch_id = b.id limit 1;`
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
	parent, err := GetUserByID(db, user.ParentId)
	if err != nil {
		return err
	}
	log.Println("Parent ", parent.BranchId)
	if parent.BranchId == models.DictBranchName["Unterabschnitt"] {
		return errors.New("unterabschnitt darf keine Abschnitten haben")
	}

	var model Id
	hashedPassword, err := functions.HashPassword(user.Password)
	if err != nil {
		return err
	}

	query := `insert into public.user(name, email, phone, password) values(?,?,?,?) returning id;`
	if err := db.Raw(query, user.Name, user.Email, user.Phone, hashedPassword).Scan(&model).Error; err != nil {
		return err
	}

	if err := AddHierarchy(db, user.ParentId, model.Id, parent.BranchId); err != nil {
		return err
	}
	return nil
}

func AddHierarchy(db *gorm.DB, parentId int, childId int, branchId int) error {
	query := `insert into hierarchy(user1_id, user2_id) values(?,?);`
	if err := db.Exec(query, parentId, childId).Error; err != nil {
		return err
	}

	branchId += 1
	query = `update public.user set branch_id = ?, update_date = now() where id = ?;`
	if err := db.Exec(query, branchId, childId).Error; err != nil {
		return err
	}
	return nil
}


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


func RevokeToken(db *gorm.DB, token string) error {
	query := `update public.user set token = null, update_date = now() where token = ?'`
	if err := db.Exec(query, token).Error; err != nil {
		return err
	}
	return nil
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
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date, branch_id 
				from public.user
				where id = (select user_id from otp where code = ? and type = ?);`
	if err = db.Raw(query, otp, reason).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}

func PatchProfile(db *gorm.DB, userId int, name string, phone string) error {
	query := `update public.user set name = ?, phone = ? where id = ?`
	if err := db.Exec(query, name, phone, userId).Error; err != nil {
		return err
	}
	return nil
}


func GetUserByID(db *gorm.DB, userId int) (user *models.User, err error) {
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date, branch_id 
				from public.user where id = ?`
	if err = db.Raw(query, userId).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}

func GetParent(db *gorm.DB, userId int) (user *models.User, err error) {
	query := `select id, name, phone, password, email, token, is_activated, is_email_verified, is_super_user, create_date, branch_id 
				from public.user
			where id = (select user1_id from hierarchy where user2_id = ?);`

	if err = db.Raw(query, userId).Scan(&user).Error; err != nil {
		return nil, err
	}

	return
}


func GetChildren(db *gorm.DB, parentId int) (users *[]models.User, err error) {
	query := `select u.id, u.name, u.email, u.branch_id, b.name as branch_name 
				from public.user u, branch b, hierarchy h 
				where b.id = u.branch_id
				and u.id = ?
				and u.id = h.user2_id;`
	if err := db.Raw(query, parentId).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
