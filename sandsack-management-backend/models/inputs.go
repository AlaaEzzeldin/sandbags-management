package models

type CreateUser struct {
	Name string `json:"name,required"`
	Email string `json:"email,required" validate:"email" structs:"email,omitempty"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	ParentId int `json:"parent_id"`
}
