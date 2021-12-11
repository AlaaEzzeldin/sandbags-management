package service

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func CreateOrder(a *gorm.DB, userName string, order *models.Order) error {
	if err := InsertOrder(a, order); err != nil {
		return err
	}
	if err := InsertOrderEquipment(a, order.Id, order.Equipments); err != nil {
		return err
	}
	if len(order.Comments) > 0 {
		if err := InsertComments(a, order.UserId, order.Id, order.Comments); err != nil {
			return err
		}
	}

	logs := []models.Log{
		{
			OrderId: order.Id,
			ActionTypeId: models.DictActionTypeName["CREATE_ORDER"],
			UpdatedBy: order.UserId,
			Description: userName + " created order " + order.Name,
		},
	}
	if err := InsertLogs(a, logs); err != nil {
		return err
	}
	userOrderPermissions := []int{
		0,
		1,
		2,
	}
	if err := InsertUserOrderPermissions(a, order.UserId, order.Id, userOrderPermissions); err != nil {
		return err
	}

	parentOrderPermissions := []int{
		0,1,2,
	}
	parent, err := GetParent(a, order.UserId)
	if err != nil {
		return err
	}
	if err := InsertUserOrderPermissions(a, parent.Id, order.Id, parentOrderPermissions); err != nil {
		return err
	}
	return nil
}

func InsertOrder(a *gorm.DB, order *models.Order) error {
	query := `insert into order(name, user_id, priority_id, status_id, address_to, address_from)values(?,?,?,?,?,?) 
				returning id;`
	err := a.Exec(query, order.Name, order.UserId, order.PriorityId, order.StatusId,
		order.AddressTo, order.AddressFrom).Scan(&order).Error
	return err
}

func InsertOrderEquipment(a *gorm.DB, orderId int, equipments []models.OrderEquipment) error {
	query := `insert into order_equipment(equipment_id, quantity, order_id)values `
	vals := []interface{}{}

	for _, row := range equipments {
		query += "(?, ?, ?),"
		vals = append(vals, row.Id, row.Quantity, orderId)
	}

	//trim the last ,
	query = query[0:len(query)-1]

	//format all vals at once
	err := a.Exec(query, vals...).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertComments(a *gorm.DB, userId, orderId int, comments []models.Comment) error {
	query := `insert into comment(order_id, comment_text, user_id) values `
	vals := []interface{}{}

	for _, row := range comments {
		query += "(?, ?, ?),"
		vals = append(vals, orderId, row.CommentText, userId)
	}

	//trim the last ,
	query = query[0:len(query)-1]

	//format all vals at once
	err := a.Exec(query, vals...).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertLogs(a *gorm.DB, logs []models.Log) error {
	query := `insert into logs(action_type_id, description, order_id, updated_by)values `
	vals := []interface{}{}

	for _, row := range logs {
		query += "(?, ?, ?, ?),"
		vals = append(vals, row.ActionTypeId, row.Description, row.OrderId, row.UpdatedBy)
	}

	//trim the last ,
	query = query[0:len(query)-1]

	//format all vals at once
	err := a.Exec(query, vals...).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertUserOrderPermissions(a *gorm.DB, userId int, orderId int, permissionList []int) error {
	query := `insert into user_order_permission(user_id, order_id, permission_id) values `
	vals := []interface{}{}

	for _, row := range permissionList {
		query += "(?, ?, ?),"
		vals = append(vals, userId, orderId, row)
	}

	//trim the last ,
	query = query[0:len(query)-1]

	//format all vals at once
	err := a.Exec(query, vals...).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserOrderPermission(a *gorm.DB, userId int, orderId int, permissionId int) error {
	query := `delete from user_order_permission where user_id = ? and order_id = ? and permission_id = ?;`
	err := a.Exec(query, userId, orderId, permissionId).Error
	return err
}


func GetOrderList(a *gorm.DB, userId int) (orderList *[]models.Order, err error) {
	query := `select 	o.id, 
						o.name, 
						o.user_id, 
						o.address_to, 
						o.address_from, 
						o.status_id,
						s.name as status_name, 
						o.priority_id, 
						o.create_date,
						o.update_date 
				from order o, status s 
				where o.status_id = s.id
				and o.id in (select order_id from user_order_permission where user_id = ?);`

	if err := a.Raw(query, userId).Scan(&orderList).Error; err != nil {
		return nil, err
	}

	for _, row := range *orderList {
		query := `select comment_text from comment where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Comments).Error
		if err != nil {
			return nil, err
		}
		query = `select name, quantity from order_equipment where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Equipments).Error
		if err != nil {
			return nil, err
		}
		query = `select description from log where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Logs).Error
		if err != nil {
			return nil, err
		}
		query = `select p.name from permission p, user_order_permission uop 
				where uop.user_id = ?
				and uop.permission_id = p.id;`
		err = a.Raw(query, row.Id).Scan(&row.Permissions).Error
		if err != nil {
			return nil, err
		}
	}

	return orderList, nil
}