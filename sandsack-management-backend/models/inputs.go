package models

type CreateUser struct {
	Name string `json:"name,required"`
	Email string `json:"email,required" validate:"email" structs:"email,omitempty"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	ParentId int `json:"parent_id"`
}

type ChangePasswordInput struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}