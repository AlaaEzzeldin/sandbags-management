package models

import "time"

type Login struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type User struct {
	Id              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Phone           string    `json:"phone,omitempty"`
	Password        string    `json:"password,omitempty"`
	Email           string    `json:"email,omitempty"`
	Token           string    `json:"token,omitempty"`
	IsActivated     bool      `json:"is_activated,omitempty"`
	IsSuperUser     bool
	IsEmailVerified bool      `json:"is_email_verified,omitempty"`
	CreateDate      time.Time `json:"create_date"`
	UpdateDate      time.Time `json:"update_date"`
}