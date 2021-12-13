package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func CreateOrder(a *gorm.DB, userName string, order *models.Order) error {
	orderId, err := InsertOrder(a, order)
	if err != nil {
		return err
	}
	order.Id = orderId
	if err := InsertOrderEquipment(a, order.Id, order.Equipments); err != nil {
		return err
	}
	log.Println(order.Comments, len(order.Comments))

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
		1,2,3,4,5,6,
	}
	if err := InsertUserOrderPermissions(a, order.UserId, order.Id, userOrderPermissions); err != nil {
		return err
	}

	parentOrderPermissions := []int{
		1,2,3,4,5,6,
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
	for _, row := range equipments {
		query := `insert into public.order_equipment(equipment_id, quantity, order_id) values(?,?,?);`
		err := a.Exec(query, row.Id, row.Quantity, orderId).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertComments(a *gorm.DB, userId, orderId int, comments []models.Comment) error {
	for _, row := range comments {
		query := `insert into public.comment(order_id, comment_text, user_id) values(?,?,?);`
		err := a.Exec(query, orderId, row.CommentText, userId).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertLogs(a *gorm.DB, logs []models.Log) error {
	for _, row := range logs {
		query := `insert into public.log(action_type_id, description, order_id, updated_by)values(?,?,?,?);`
		err := a.Exec(query, row.ActionTypeId, row.Description, row.OrderId, row.UpdatedBy).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertUserOrderPermissions(a *gorm.DB, userId int, orderId int, permissionList []int) error {
	for _, row := range permissionList {
		query := `insert into public.user_order_permission(user_id, order_id, permission_id) values(?,?,?);`
		err := a.Exec(query, userId, orderId, row).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUserOrderPermission(a *gorm.DB, userId int, orderId int, permissionId int) error {
	query := `delete from public.user_order_permission where user_id = ? and order_id = ? and permission_id = ?;`
	err := a.Exec(query, userId, orderId, permissionId).Error
	return err
}

func GetOrder(a *gorm.DB, userId, orderId int) (order *models.Order, err error) {
	query := `select 	o.id,
						  o.name,
						  o.user_id,
						  o.address_to,
						  o.address_from,
						  o.status_id,
						  s.name as status_name,
						  o.priority_id,
						  o.create_date,
						  o.update_date,
							u.name
				from public.order o, public.status s, public.user u
				where o.status_id = s.id
				  and o.id in (select order_id from user_order_permission where user_id = ? and permission_id= (select id from permission where name = 'CAN VIEW'))
				and o.user_id = u.id
				and o.id=?;
	`

	err = a.Raw(query, userId, orderId).Scan(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil


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
				and o.id in (select order_id from user_order_permission where user_id = ? 
							 and permission_id = (select id from permission where name = 'CAN VIEW'));`

	if err := a.Raw(query, userId).Scan(&orderList).Error; err != nil {
		log.Println("GetOrderList err", err.Error())
		return nil, err
	}

	for _, row := range *orderList {
		log.Println("Row", row)
		query := `select comment_text from public.comment where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&row.Comments).Error
		if err != nil {
			return nil, err
		}
		log.Println("Comments", row.Comments)
		query = `select oe.id, e.name, oe.quantity from public.order_equipment oe, public.equipment e where oe.order_id = ? and e.id=oe.equipment_id;`
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

func ExistUserOrderPermission(db *gorm.DB, userId, orderId int) bool {
	query := `select count(1) from public.user_order_permission where user_id = ? and order_id = ?;`
	var count int
	if err := db.Raw(query, userId, orderId).Scan(&count).Error; err != nil {
		log.Println("ExistUserOrderPermission error", err.Error())
		return false
	}
	if count == 0 {
		return false
	}
	return true
}


func GetUserOrderPermissions(a *gorm.DB, userId, orderId int) ([]int, error) {
	if exist := ExistUserOrderPermission(a, userId, orderId); !exist {
		return nil, errors.New("user does not have permissions")
	}
	var permissions []models.Permission
	query := `select permission_id from public.user_order_permission where user_id = ? and order_id = ?;`
	err := a.Raw(query, userId, orderId).Scan(&permissions).Error

	var perms []int
	for _, row := range permissions {
		perms = append(perms, row.Id)
	}
	return perms, err
}

func AcceptOrder(db *gorm.DB, userId, orderId int) error {
	user, err := GetUserByID(db, userId)
	if err != nil {
		return err
	}

	children, err := GetChildren(db, userId)

	order, err := GetOrder(db, userId, orderId)

	var statusId int
	if user.BranchId == 5 { // Unterabschnitt
		return errors.New("User from Unterabschnitt cannot accept orders")
	}

	if user.BranchId == 4 { // Einsatzabschnitt
		if statusId == 1 {
			//can do
		} else {
			return errors.New("cannot weitergeleit")
		}
	}

	if user.BranchId == 3 { // Hauptabschnitt
		if order.StatusId == 3 {
			statusId = 4
		} else {
			return errors.New("cannot weiterhgeleit")
		}
	}

	if user.BranchId == 4 {
		if order.StatusId == 4 { //weitergeleitet bei Hauptabschnitt
			statusId = 6 // akzeptiert
		}
	}

	query := `update public.order set status_id = ? where id = ?;`
	err = db.Exec(query, statusId, orderId).Error
	if err != nil {
		return nil
	}

	// permissionId = 1 -- CAN VIEW
	for _, child := range *children {
		query = `delete from user_order_permission 
			where user_id = ? and order_id = ? and 
			permission_id in 
				(select id from permission 
				where name = 'CAN EDIT' or name = 'CAN DECLINE' or name = 'CAN ACCEPT'));`
		err = db.Exec(query, child.Id, order.Id).Error
		if err != nil {
			return err
		}
	}

	return nil
}


