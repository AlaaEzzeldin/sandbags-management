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

type SendVerifyEmail struct {
	Email string `json:"email,required"`
}

type VerifyEmailInput struct {
	Otp string `json:"otp"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type SendRecoveryPasswordInput struct {
	Email string `json:"email"`
}

type RecoveryPasswordInput struct {
	OTP      string `json:"otp"`
	Password string `json:"password"`
}