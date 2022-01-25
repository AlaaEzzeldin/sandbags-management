package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	repo_user "team2/sandsack-management-backend/repository/user"
)

type Id struct {
	Id int `gorm:"column:id"`
}

func GetUserByEmail(db *gorm.DB, email string) (user *models.User, err error) {
	if exist := CheckUserExists(db, email); !exist {
		return nil, errors.New("user not found")
	}
	return repo_user.GetUserByEmail(db, email)
}

func GetUserByToken(db *gorm.DB, token string) (user models.User, err error) {
	return repo_user.GetUserByRefreshToken(db, token)
}

func CheckUserExists(db *gorm.DB, email string) bool {
	return repo_user.CheckUserExists(db, email)
}

func UpdateUserToken(db *gorm.DB, email string, refreshToken string) error {
	return repo_user.UpdateUserToken(db, email, refreshToken)
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

	hashedPassword, err := functions.HashPassword(user.Password)
	if err != nil {
		return err
	}

	userId, err := repo_user.InsertUser(db, user, hashedPassword)
	if err != nil {
		return err
	}

	if err := AddHierarchy(db, user.ParentId, userId, parent.BranchId); err != nil {
		return err
	}
	return nil
}

func AddHierarchy(db *gorm.DB, parentId int, childId int, branchId int) error {
	if err := repo_user.InsertHierarchy(db, parentId, childId); err != nil {
		return err
	}

	branchId += 1
	if branchId > 5 {
		return errors.New("cannot be child of unterabschnitt")
	}
	if err := repo_user.UpdateBranch(db, childId, branchId); err != nil {
		return err
	}
	return nil
}

func GetUserList(db *gorm.DB) (userList *[]models.User, err error) {
	userList, err = repo_user.GetUserList(db)
	if err != nil {
		return nil, err
	}

	mollnhof, err := repo_user.GetUserByID(db, 2)
	mollnhofUser := &models.User{
		Id: mollnhof.Id,
		Name: mollnhof.Name,
		Email: mollnhof.Email,
		BranchName: mollnhof.BranchName,
		BranchId: mollnhof.BranchId,
		CreateDate: mollnhof.CreateDate,
		UpdateDate: mollnhof.UpdateDate,
		ParentId: 0,
		ParentName: "",
	}
	*userList = append(*userList, *mollnhofUser)
	return userList, nil
}

func RevokeToken(db *gorm.DB, token string) error {
	return repo_user.RevokeToken(db, token)
}

func UpdatePassword(db *gorm.DB, email, password string) error {
	return repo_user.UpdatePassword(db, email, password)
}

func UpdateUserActivity(db *gorm.DB, email string, isActivated bool) error {
	return repo_user.UpdateUserActivity(db, email, isActivated)
}

func VerifyUserEmail(db *gorm.DB, email string, isVerified bool) error {
	return repo_user.VerifyUserEmail(db, email, isVerified)
}

func GetUserByOTP(db *gorm.DB, otp, reason string)  (user *models.User, err error) {
	return repo_user.GetUserByOTP(db, otp, reason)
}

func PatchProfile(db *gorm.DB, userId int, name string, phone string, email string) error {
	return repo_user.PatchProfile(db, userId, name, phone, email)
}

func GetUserByID(db *gorm.DB, userId int) (user *models.User, err error) {
	return repo_user.GetUserByID(db, userId)
}

func GetChildren(db *gorm.DB, parentId int) (users *[]models.User, err error) {
	return repo_user.GetChildren(db, parentId)
}
