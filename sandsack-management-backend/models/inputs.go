package models

type CreateUser struct {
	Name string `json:"name,required"`
	Email string `json:"email,required" validate:"email" structs:"email,omitempty"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	ParentId int `json:"parent_id"`
}

type Logout struct {
	Token 	string `json:"token,required"`
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
	Password string `json:"password"`
}

type SendRecoveryPasswordInput struct {
	Email string `json:"email"`
}

type RecoveryPasswordInput struct {
	OTP      string `json:"otp"`
	Password string `json:"password"`
}

type PatchProfileInput struct {
	Name string `json:"name"`
	Phone string `json:"phone"`

type CreateOrderInput struct {
	AddressTo  string           `json:"address_to,required"`
	Priority   int              `json:"priority,required"`
	Equipments []OrderEquipment `json:"equipments,required"`
	Comments   []Comment        `json:"comments,omitempty"`
}


type AcceptOrderInput struct {
	OrderId int `json:"order_id"`
}