package models

import "time"

type Login struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type User struct {
	Id              int       `json:"id,omitempty" gorm:"column:id"`
	Name            string    `json:"name,omitempty" gorm:"column:name"`
	Phone           string    `json:"phone,omitempty" gorm:"column:phone"`
	Password        string    `json:"password,omitempty" gorm:"column:password"`
	Email           string    `json:"email,omitempty" gorm:"column:email"`
	Token           string    `json:"token,omitempty" gorm:"column:token"`
	IsActivated     bool      `json:"is_activated,omitempty" gorm:"column:is_activated"`
	IsSuperUser     bool      `gorm:"column:is_super_user"`
	IsEmailVerified bool      `json:"is_email_verified,omitempty" gorm:"column:is_email_verified"`
	CreateDate      time.Time `json:"create_date" gorm:"column:create_date"`
	UpdateDate      time.Time `json:"update_date" gorm:"column:update_date"`
}

