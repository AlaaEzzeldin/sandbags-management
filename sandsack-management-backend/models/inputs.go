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
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type CreateOrderInput struct {
	AddressTo  string           `json:"address_to,required"`
	Priority   int              `json:"priority,required"`
	Equipments []OrderEquipment `json:"equipments,required"`
	Comment   string       `json:"comment,omitempty"`
}


type AcceptOrderInput struct {
	OrderId int `json:"order_id"`
}

type CommentInput struct {
	OrderId int `json:"order_id"`
	Comment string `json:"comment"`
}

type ConfirmDeliveryInput struct {
	OrderId int `json:"order_id"`
}


type DispatchOrderInput struct {
	OrderId int `json:"order_id,required"`
	DriverId int `json:"driver_id,omitempty"`
}

type GetStatisticsInput struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type EditOrderInput struct {
	OrderId int `json:"order_id"`
	Equipments []OrderEquipment `json:"equipments,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AddDriverInput struct {
	Name string `json:"name,required"`
	Description string `json:"description,omitempty"`
}

type UpdateEquipmentInput struct {
	Id int `json:"id"`
	Quantity int `json:"quantity"`
}

type GetAllOrders struct {
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}