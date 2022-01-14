package models

import "time"

type OrderEquipment struct {
	EquipmentId       int `json:"id,omitempty" gorm:"column:equipment_id"`
	Name string `json:"name,omitempty" gorm:"column:name"`
	Quantity int `json:"quantity,omitempty" gorm:"column:quantity"`
}

type Comment struct {
	CommentId   int       `json:"id,omitempty" gorm:"column:comment_id"`
	CommentText string    `json:"comment_text,omitempty" gorm:"column:comment_text"`
	OrderId     int       `json:"order_id,omitempty" gorm:"column:order_id"`
	Name        string    `json:"name,omitempty" gorm:"column:name"`
	Role  string    `json:"role,omitempty" gorm:"column:role"`
	CreateDate  time.Time `json:"create_date,omitempty" gorm:"column:create_date"`
	UpdateDate  time.Time `json:"update_date,omitempty" gorm:"column:update_date"`
}


type Log struct {
	LogId             int       `json:"id,omitempty" gorm:"column:log_id"`
	OrderId        int       `json:"order_id,omitempty" gorm:"column:order_id"`
	Description    string    `json:"description,omitempty" gorm:"column:description"`
	ActionTypeId   int       `json:"action_type_id,omitempty" gorm:"column:action_type_id"`
	ActionTypeName string    `json:"action_type_name,omitempty" gorm:"column:action_type_name"`
	UpdatedBy      int       `json:"updated_by,omitempty" gorm:"column:updated_by"`
	UpdatedByName  string    `json:"updated_by_name,omitempty" gorm:"column:updated_by_name"`
	CreateDate     time.Time `json:"create_date,omitempty" gorm:"column:create_date"`
	UpdateDate     time.Time `json:"update_date,omitempty" gorm:"column:update_date"`
}

type Permission struct {
	PermissionId   int    `json:"id,omitempty" gorm:"column:permission_id"`
	Name string `json:"name,omitempty" gorm:"column:name"`
}

type Order struct {
	Id            int              `json:"id,omitempty" gorm:"column:id"`
	UserId        int              `json:"user_id,omitempty" gorm:"column:user_id"`
	Name          string           `json:"name,omitempty" gorm:"column:name"`
	AddressTo     string           `json:"address_to,omitempty" gorm:"column:address_to"`
	AddressFrom   string           `json:"address_from,omitempty" gorm:"column:address_from"`
	StatusId      int              `json:"status_id,omitempty" gorm:"column:status_id"`
	StatusName    string           `json:"status_name,omitempty" gorm:"column:status_name"`
	PriorityId    int              `json:"priority_id,omitempty" gorm:"column:priority_id"`
	AssignedTo    int              `json:"assigned_to,omitempty" gorm:"column:assigned_to"`
	CreateDate    time.Time        `json:"create_date,omitempty" gorm:"column:create_date"`
	UpdateDate    time.Time        `json:"update_date,omitempty" gorm:"column:update_date"`
	Comments      []Comment        `json:"comments,omitempty"`
	Logs          []Log            `json:"logs,omitempty"`
	Equipments    []OrderEquipment `json:"equipments,omitempty"`
	Permissions   []Permission     `json:"permissions,omitempty"`
	EstimatedTime int              `json:"estimated_time,omitempty"`
}

type SimpleOrder struct {
	Id          int              `json:"id,omitempty" gorm:"column:id"`
	UserId      int              `json:"user_id,omitempty" gorm:"column:user_id"`
	Name        string           `json:"name,omitempty" gorm:"column:name"`
	AddressTo   string           `json:"address_to,omitempty" gorm:"column:address_to"`
	AddressFrom string           `json:"address_from,omitempty" gorm:"column:address_from"`
	StatusId    int              `json:"status_id,omitempty" gorm:"column:status_id"`
	StatusName  string           `json:"status_name,omitempty" gorm:"column:status_name"`
	PriorityId  int              `json:"priority_id,omitempty" gorm:"column:priority_id"`
	AssignedTo  int              `json:"assigned_to,omitempty" gorm:"column:assigned_to"`
	CreateDate  time.Time        `json:"create_date,omitempty" gorm:"column:create_date"`
	UpdateDate  time.Time        `json:"update_date,omitempty" gorm:"column:update_date"`
}

type Priority struct {
	Id int `json:"id,omitempty" gorm:"column:id"`
	Level int `json:"level,omitempty" gorm:"column:level"`
	Name string `json:"name,omitempty" gorm:"column:name"`
}