package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
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
			ActionTypeId: models.DictActionTypeName["CREATED"],
			UpdatedBy: order.UserId,
			Description: userName + " created order " + order.Name,
		},
	}
	if err := InsertLogs(a, logs); err != nil {
		return err
	}
	userOrderPermissions := []int{
		models.DictPermissionName["CAN VIEW"],
		models.DictPermissionName["CAN EDIT"],
		models.DictPermissionName["CAN DECLINE"],
		models.DictPermissionName["CAN COMMENT"],
	}
	if err := InsertUserOrderPermissions(a, order.UserId, order.Id, userOrderPermissions); err != nil {
		return err
	}

	parentOrderPermissions := []int{
		models.DictPermissionName["CAN VIEW"],
		models.DictPermissionName["CAN EDIT"],
		models.DictPermissionName["CAN DECLINE"],
		models.DictPermissionName["CAN ACCEPT"],
		models.DictPermissionName["CAN COMMENT"],
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
	query = `update public.order set name = ?, update_date = ? where id = ?;`
	err = a.Exec(query, order.Name + " #" + strconv.Itoa(id.Id), id.Id).Error

	return id.Id, nil
}

func InsertOrderEquipment(a *gorm.DB, orderId int, equipments []models.OrderEquipment) error {
	for _, row := range equipments {
		query := `insert into public.order_equipment(equipment_id, quantity, order_id) values(?,?,?);`
		err := a.Exec(query, row.EquipmentId, row.Quantity, orderId).Error
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

func GetOrder(a *gorm.DB, userId, orderId int) (models.Order, error) {
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
	var simpleOrder models.SimpleOrder

	err := a.Raw(query, userId, orderId).Scan(&simpleOrder).Error
	if err != nil {
		return models.Order{}, err
	}

	var (
		comments []models.Comment
		equipments []models.OrderEquipment
		logs []models.Log
		permissions []models.Permission
	)
	query = `select comment_text from public.comment where order_id = ?;`
	err = a.Raw(query, simpleOrder.Id).Scan(&comments).Error

	query = `select oe.id as equipment_id, e.name, oe.quantity 
				from public.order_equipment oe, public.equipment e 
				where oe.order_id = ? and e.id = oe.equipment_id;`
	err = a.Raw(query, simpleOrder.Id).Scan(&equipments).Error

	query = `select description from public.log where order_id = ?;`
	err = a.Raw(query, simpleOrder.Id).Scan(&logs).Error

	query = `select p.name from public.permission p, public.user_order_permission uop 
				where uop.user_id = ? and uop.order_id = ?
				and uop.permission_id = p.id;`
	err = a.Raw(query, userId, simpleOrder.Id).Scan(&permissions).Error

	order := models.Order{
		Id: simpleOrder.Id,
		UserId: simpleOrder.UserId,
		StatusId: simpleOrder.StatusId,
		Name: simpleOrder.Name,
		AddressFrom: simpleOrder.AddressFrom,
		AddressTo: simpleOrder.AddressTo,
		StatusName: simpleOrder.StatusName,
		PriorityId: simpleOrder.PriorityId,
		Comments: comments,
		Logs: logs,
		Equipments: equipments,
		Permissions: permissions,
	}
	return order, nil
}


func GetOrderList(a *gorm.DB, userId int) (orderList []models.Order, err error) {
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
							 and permission_id = (select id from permission where name = 'CAN VIEW'))
				order by o.update_date DESC;`

	var simpleOrderList []models.SimpleOrder

	if err := a.Raw(query, userId).Scan(&simpleOrderList).Error; err != nil {
		log.Println("GetOrderList err", err.Error())
		log.Println("Here", err.Error() )
		return nil, err
	}

	for _, row := range simpleOrderList {
		var (
			comments []models.Comment
			equipments []models.OrderEquipment
			logs []models.Log
			permissions []models.Permission
		)
		query := `select comment_text from public.comment where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&comments).Error

		log.Println("Comments", comments)
		query = `select oe.id as equipment_id, e.name, oe.quantity 
				from public.order_equipment oe, public.equipment e 
				where oe.order_id = ? and e.id = oe.equipment_id;`
		err = a.Raw(query, row.Id).Scan(&equipments).Error

		query = `select description from public.log where order_id = ?;`
		err = a.Raw(query, row.Id).Scan(&logs).Error

		query = `select p.name from public.permission p, public.user_order_permission uop 
				where uop.user_id = ? and uop.order_id = ?
				and uop.permission_id = p.id;`
		err = a.Raw(query, userId, row.Id).Scan(&permissions).Error

		order := models.Order{
			Id: row.Id,
			UserId: row.UserId,
			StatusId: row.StatusId,
			Name: row.Name,
			AddressFrom: row.AddressFrom,
			AddressTo: row.AddressTo,
			StatusName: row.StatusName,
			PriorityId: row.PriorityId,
			Comments: comments,
			Logs: logs,
			Equipments: equipments,
			Permissions: permissions,
		}

		orderList = append(orderList, order)
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
		perms = append(perms, row.PermissionId)
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

	log.Println("USER branchID:", user.BranchId)
	log.Println("Order", order.StatusId)

	var statusId int
	if user.BranchId == models.DictBranchName["Unterabschnitt"] {
		return errors.New("unterabschnitt kann nicht Bestellung akzeptieren")

	} else if user.BranchId == models.DictBranchName["Einsatzabschnitt"]{
		if order.StatusId == models.DictStatusName["ANSTEHEND"] ||
		   order.StatusId == models.DictStatusName["ABGELEHNT BEI EINSATZABSCHNITT"] {

				statusId = models.DictStatusName["WEITERGELEITET BEI EINSATZABSCHNITT"]

		} else {
			return errors.New("einsatzabschnitt kann nicht weiterleiten")
		}
	} else if user.BranchId == models.DictBranchName["Hauptabschnitt"] {
		if order.StatusId == models.DictStatusName["WEITERGELEITET BEI EINSATZABSCHNITT"] ||
		   order.StatusId == models.DictStatusName["ABGELEHNT BEI HAUPTABSCHNITT"] {

				statusId = models.DictStatusName["WEITERGELEITET BEI HAUPTABSCHNITT"]

		} else {
			return errors.New("hauptabschnitt kann nicht weiterleiten")
		}
	} else if  user.BranchId == models.DictBranchName["Einsatzleiter"] {
		if order.StatusId == models.DictStatusName["WEITERGELEITET BEI HAUPTABSCHNITT"] ||
		   order.StatusId == models.DictStatusName["ABGELEHNT BEI EINSATZLEITER"] {

				statusId = models.DictStatusName["AKZEPTIERT"]

		} else {
			return errors.New("einsatzleiter kann nicht weiterleiten")
		}
	} else if user.BranchId == models.DictBranchName["Mollnhof"] {
		if order.StatusId == models.DictStatusName["AKZEPTIERT"] {
			statusId = models.DictStatusName["AKZEPTIERT"]

		} else {
			return errors.New("mollnhof kann nicht weiterleiten")
		}
	} else {
		return errors.New("something went wrong")
	}

	query := `update public.order set status_id = ?, update_date = now() where id = ?;`
	err = db.Exec(query, statusId, orderId).Error
	if err != nil {
		return nil
	}

	for _, child := range *children {
		query = `delete from user_order_permission 
			where user_id = ? and order_id = ? and 
			permission_id in 
				(select id from permission 
				where name = 'CAN EDIT' or name = 'CAN DECLINE' or name = 'CAN ACCEPT');`
		err = db.Exec(query, child.Id, order.Id).Error
		if err != nil {
			return err
		}
	}

	if user.BranchId > models.DictBranchName["Mollnhof"] {
		parent, err := GetParent(db, user.Id)
		if err != nil {
			return err
		}
		parentPermissions := []int{
			models.DictPermissionName["CAN VIEW"],
			models.DictPermissionName["CAN EDIT"],
			models.DictPermissionName["CAN DECLINE"],
			models.DictPermissionName["CAN ACCEPT"],
			models.DictPermissionName["CAN COMMENT"],
		}
		if parent.BranchId == 1 {
			parentPermissions = []int{
				models.DictPermissionName["CAN VIEW"],
				models.DictPermissionName["CAN CONFIRM DELIVERY"],
				models.DictPermissionName["CAN ASSIGN TO"],
				models.DictPermissionName["CAN ACCEPT"],
			}
		}
		err = InsertUserOrderPermissions(db, parent.Id, orderId, parentPermissions)
		if err != nil {
			return err
		}

	}

	logs := []models.Log{
		{
			OrderId: orderId,
			ActionTypeId: models.DictActionTypeName["ACCEPTED"],
			UpdatedBy: userId,
			Description: user.Name + " accepted order " + order.Name,
		},
	}

	err = InsertLogs(db, logs)
	if err != nil {
		return err
	}

	return nil
}

func DeclineOrder(db *gorm.DB, userId, orderId int) error {
	user, _ := GetUserByID(db, userId)
	order, _ := GetOrder(db, userId, orderId)

	var statusId int

	if order.StatusId == models.DictStatusName["STORNIERT"] ||
		order.StatusId == models.DictStatusName["ABGELEHNT BEI EINSATZABSCHNITT"] ||
		order.StatusId == models.DictStatusName["ABGELEHNT BEI HAUPTABSCHNITT"] ||
		order.StatusId == models.DictStatusName["ABGELEHNT BEI EINSATZLEITER"] {

		return errors.New("die Bestellung wurde bereits storniert")
	}

	if user.BranchId == models.DictBranchName["Unterabschnitt"] {
		if order.StatusId == models.DictStatusName["ANSTEHEND"] {
				statusId = models.DictStatusName["STORNIERT"]
		} else {
			return errors.New("unterabschnitt kann die Bestellung nicht stornieren")
		}
	} else if user.BranchId == models.DictBranchName["Einsatzabschnitt"] {
		if order.StatusId == models.DictStatusName["ANSTEHEND"] ||
		   order.StatusId == models.DictStatusName["WEITERGELEITET BEI EINSATZABSCHNITT"] {

				statusId = models.DictStatusName["ABGELEHNT BEI EINSATZABSCHNITT"]
		} else {
			return errors.New("einsatzabschnitt kann die Bestellung nicht stornieren")
		}
	} else if user.BranchId == models.DictBranchName["Hauptabschnitt"] {
		if order.StatusId == models.DictStatusName["WEITERGELEITET BEI EINSATZABSCHNITT"] ||
		   order.StatusId == models.DictStatusName["WEITERGELEITET BEI HAUPTABSCHNITT"] {

				statusId = models.DictStatusName["ABGELEHNT BEI HAUPTABSCHNITT"]
		} else {
			return errors.New("hauptabschnitt kann die Bestellung nicht stornieren")
		}
	} else if user.BranchId == models.DictBranchName["Einsatzleiter"] {
		if order.StatusId == models.DictStatusName["WEITERGELEITET BEI HAUPTABSCHNITT"] {

			statusId = models.DictStatusName["ABGELEHNT BEI EINSATZLEITER"]
		} else {
			return errors.New("einsatzleiter kann die Bestellung nicht stornieren")
		}
	} else if user.BranchId == models.DictBranchName["Mollnhof"] {
		return errors.New("mollnhof kann die Bestellung nicht stornieren")
	} else {
		return errors.New("something went wrong")
	}

	query := `update public.order set status_id = ?, update_date = now() where id = ?;`
	if err := db.Exec(query, statusId, orderId).Error; err != nil {
		return err
	}

	if user.BranchId < 5 {
		children, err := GetChildren(db, user.Id)
		if err != nil {
			return err
		}
		for _, child := range *children {
			query = `delete from user_order_permission 
				where user_id = ? and order_id = ?
				and permission_id in 
					(select id 
					from permission 
					where name = 'CAN EDIT' or name = 'CAN DECLINE' or name = 'CAN ACCEPT');`

			err = db.Exec(query, child.Id, orderId).Error
			if err != nil {
				return err
			}
		}
	}

	logs := []models.Log{
		{
			OrderId: orderId,
			ActionTypeId: models.DictActionTypeName["DECLINED"],
			UpdatedBy: userId,
			Description: user.Name + " declined order " + order.Name,
		},
	}

	err := InsertLogs(db, logs)
	if err != nil {
		return err
	}

	return nil

}
