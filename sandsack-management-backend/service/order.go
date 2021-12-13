package service

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func CreateOrder(a *gorm.DB, userName string, order *models.Order) error {
	orderId, err := InsertOrder(a, order)
	if err != nil {
		return err
	}
	if err := InsertOrderEquipment(a, orderId, order.Equipments); err != nil {
		return err
	}
	log.Println(order.Comments, len(order.Comments))

	if len(order.Comments) > 0 {
		if err := InsertComments(a, order.UserId, orderId, order.Comments); err != nil {
			return err
		}
	}

	logs := []models.Log{
		{
			OrderId: orderId,
			ActionTypeId: models.DictActionTypeName["CREATE_ORDER"],
			UpdatedBy: order.UserId,
			Description: userName + " created order " + order.Name,
		},
	}
	if err := InsertLogs(a, logs); err != nil {
		return err
	}
	userOrderPermissions := []int{
		1,2,3,4,5,6,
	}
	if err := InsertUserOrderPermissions(a, order.UserId, orderId, userOrderPermissions); err != nil {
		return err
	}

	parentOrderPermissions := []int{
		1,2,3,4,5,6,
	}
	parent, err := GetParent(a, order.UserId)
	if err != nil {
		return err
	}
	if err := InsertUserOrderPermissions(a, parent.Id, orderId, parentOrderPermissions); err != nil {
		return err
	}
	return nil
}

func InsertOrder(a *gorm.DB, order *models.Order) (int, error) {
	query := `insert into public.order(name, user_id, priority_id, status_id, address_to, address_from)values(?,?,?,?,?,?) 
				returning id;`
	var id Id
	err := a.Raw(query, order.Name, order.UserId, order.PriorityId, order.StatusId,
		order.AddressTo, order.AddressFrom).Scan(&id).Error
	if err != nil {
		log.Println("InsertOrder error", err.Error())
		//a.Rollback()
		return 0, err
	}
	return id.Id, nil
}

func InsertOrderEquipment(a *gorm.DB, orderId int, equipments []models.OrderEquipment) error {
	query := `insert into public.order_equipment(equipment_id, quantity, order_id)values `
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
		log.Println("InsertOrderEquipment error", err.Error())
		//a.Rollback()
		return err
	}
	return nil
}

func InsertComments(a *gorm.DB, userId, orderId int, comments []models.Comment) error {
	query := `insert into public.comment(order_id, comment_text, user_id) values `
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
		log.Println("InsertComments error", err.Error())
		//a.Rollback()
		return err
	}
	return nil
}

func InsertLogs(a *gorm.DB, logs []models.Log) error {
	query := `insert into public.log(action_type_id, description, order_id, updated_by)values `
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
		log.Println("InsertLogs error", err.Error())
		return err
	}
	return nil
}

func InsertUserOrderPermissions(a *gorm.DB, userId int, orderId int, permissionList []int) error {
	query := `insert into public.user_order_permission(user_id, order_id, permission_id) values `
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
		log.Println("InsertUserOrderPermissions error", err.Error())
		//a.Rollback()
		return err
	}
	return nil
}

func DeleteUserOrderPermission(a *gorm.DB, userId int, orderId int, permissionId int) error {
	query := `delete from public.user_order_permission where user_id = ? and order_id = ? and permission_id = ?;`
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
				from public.order o, public.status s 
				where o.status_id = s.id
				and o.id in (select order_id from user_order_permission where user_id = ?);`

	if err := a.Raw(query, userId).Scan(&orderList).Error; err != nil {
		return nil, err
	}

	for _, row := range *orderList {
		query := `select comment_text from public.comment where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Comments).Error
		if err != nil {
			return nil, err
		}
		query = `select name, quantity from public.order_equipment where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Equipments).Error
		if err != nil {
			return nil, err
		}
		query = `select description from public.log where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Logs).Error
		if err != nil {
			return nil, err
		}
		query = `select p.name from public.permission p, public.user_order_permission uop 
				where uop.user_id = ?
				and uop.permission_id = p.id;`
		err = a.Raw(query, row.Id).Scan(&row.Permissions).Error
		if err != nil {
			return nil, err
		}
	}

	return orderList, nil
}