package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/repository/order"
	repo_user "team2/sandsack-management-backend/repository/user"
)

func CreateOrder(a *gorm.DB, userName string, order *models.Order) error {
	orderId, err := repo_order.CreateOrder(a, order)
	if err != nil {
		return err
	}
	order.Id = orderId
	if err := repo_order.InsertOrderEquipment(a, order.Id, order.Equipments); err != nil {
		log.Println("InsertOrderEquipment error", err.Error())
		return err
	}

	if len(order.Comments) > 0 {
		if err := repo_order.InsertComments(a, order.UserId, order.Id, order.Comments); err != nil {
			log.Println("InsertComments error", err.Error())
			return err
		}
	}


	userOrderPermissions := []int{
		models.DictPermissionName["CAN VIEW"],
		models.DictPermissionName["CAN EDIT"],
		models.DictPermissionName["CAN DECLINE"],
		models.DictPermissionName["CAN COMMENT"],
	}
	if err := repo_order.InsertUserOrderPermissions(a, order.UserId, order.Id, userOrderPermissions); err != nil {
		log.Println("InsertUserOrderPermissions error", err.Error())
		return err
	}

	parentOrderPermissions := []int{
		models.DictPermissionName["CAN VIEW"],
		models.DictPermissionName["CAN EDIT"],
		models.DictPermissionName["CAN DECLINE"],
		models.DictPermissionName["CAN ACCEPT"],
		models.DictPermissionName["CAN COMMENT"],
	}
	parent, err := repo_user.GetParent(a, order.UserId)
	if err != nil {
		log.Println("GetParent error", err.Error())
		return err
	}
	if err := repo_order.InsertUserOrderPermissions(a, parent.Id, order.Id, parentOrderPermissions); err != nil {
		log.Println("InsertUserOrderPermissions error", err.Error())
		return err
	}

	newOrder, err := GetOrder(a, order.UserId, orderId)
	log.Println("NEW ORDER NAME", newOrder.Name)

	logs := []models.Log{
		{
			OrderId:      order.Id,
			ActionTypeId: models.DictActionTypeName["CREATED"],
			UpdatedBy:    order.UserId,
			Description:  userName + " hat die Bestellung '" +  newOrder.Name + " #" + strconv.Itoa(newOrder.Id) + "' erstellt",
		},
	}

	if err := repo_order.InsertLogs(a, logs); err != nil {
		log.Println("InsertLogs error", err.Error())
		return err
	}

	return nil
}

func GetOrder(a *gorm.DB, userId, orderId int) (models.Order, error) {
	simpleOrder, err := repo_order.GetSimpleOrder(a, userId, orderId)
	if err != nil {
		return models.Order{}, err
	}
	comments := repo_order.GetComments(a, orderId)
	equipments := repo_order.GetEquipments(a, orderId)
	logs := repo_order.GetLogs(a, orderId)
	//permissions := repo_order.GetPermissions(a, userId, orderId)

	order := models.Order{
		Id:          simpleOrder.Id,
		UserId:      simpleOrder.UserId,
		StatusId:    simpleOrder.StatusId,
		Name:        simpleOrder.Name,
		AddressFrom: simpleOrder.AddressFrom,
		AddressTo:   simpleOrder.AddressTo,
		StatusName:  simpleOrder.StatusName,
		PriorityId:  simpleOrder.PriorityId,
		Comments:    comments,
		Logs:        logs,
		Equipments:  equipments,
		CreateDate: simpleOrder.CreateDate,
		UpdateDate: simpleOrder.UpdateDate,
		//Permissions: permissions,
	}
	return order, nil
}

func GetOrderList(a *gorm.DB, userId int) (orderList []models.Order, err error) {
	simpleOrderList, err := repo_order.GetSimpleOrderList(a, userId)
	if err != nil {
		return nil, err
	}

	for _, row := range simpleOrderList {
		comments := repo_order.GetComments(a, row.Id)
		equipments := repo_order.GetEquipments(a, row.Id)
		logs := repo_order.GetLogs(a, row.Id)
		//permissions := repo_order.GetPermissions(a, userId, row.Id)

		order := models.Order{
			Id:          row.Id,
			UserId:      row.UserId,
			StatusId:    row.StatusId,
			Name:        row.Name,
			AddressFrom: row.AddressFrom,
			AddressTo:   row.AddressTo,
			StatusName:  row.StatusName,
			PriorityId:  row.PriorityId,
			Comments:    comments,
			Logs:        logs,
			Equipments:  equipments,
			CreateDate: row.CreateDate,
			UpdateDate: row.UpdateDate,
			//Permissions: permissions,
		}

		orderList = append(orderList, order)
	}

	return orderList, nil
}

func GetUserOrderPermissions(a *gorm.DB, userId, orderId int) ([]int, error) {
	if exist := repo_order.ExistUserOrderPermission(a, userId, orderId); !exist {
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

	var statusId int
	if user.BranchId == models.DictBranchName["Unterabschnitt"] {
		return errors.New("unterabschnitt kann nicht Bestellung akzeptieren")

	} else if user.BranchId == models.DictBranchName["Einsatzabschnitt"] {
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
	} else if user.BranchId == models.DictBranchName["Einsatzleiter"] {
		if order.StatusId == models.DictStatusName["WEITERGELEITET BEI HAUPTABSCHNITT"] ||
			order.StatusId == models.DictStatusName["ABGELEHNT BEI EINSATZLEITER"] {

			statusId = models.DictStatusName["AKZEPTIERT"]

		} else {
			return errors.New("einsatzleiter kann nicht weiterleiten")
		}
	} else if user.BranchId == models.DictBranchName["Mollnhof"] {
		/*if order.StatusId == models.DictStatusName["AKZEPTIERT"] {
			statusId = models.DictStatusName["AUF DEM WEG"]
		 */
		return errors.New("mollnhof kann nicht weiterleiten")

	} else {
		return errors.New("something went wrong")
	}

	query := `update public.order set status_id = ?, update_date = now() where id = ?;`
	err = db.Exec(query, statusId, orderId).Error
	if err != nil {
		return nil
	}

	log.Println("CHILDREN", children)
	for _, child := range *children {
		err = repo_order.DeleteUserOrderPermission(db, child.Id, order.Id)
		if err != nil {
			return err
		}
	}

	if user.BranchId > models.DictBranchName["Mollnhof"] {
		parent, err := repo_user.GetParent(db, user.Id)
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
		err = repo_order.InsertUserOrderPermissions(db, parent.Id, orderId, parentPermissions)
		if err != nil {
			return err
		}

	}

	logs := []models.Log{
		{
			OrderId:      orderId,
			ActionTypeId: models.DictActionTypeName["ACCEPTED"],
			UpdatedBy:    userId,
			Description:  user.Name + " accepted order " + order.Name,
		},
	}

	err = repo_order.InsertLogs(db, logs)
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

	if err := repo_order.UpdateOrderStatus(db, orderId, statusId); err != nil {
		return err
	}

	if user.BranchId < 5 {
		children, err := GetChildren(db, user.Id)
		if err != nil {
			return err
		}
		log.Println("Children", children)
		for _, child := range *children {
			err = repo_order.DeleteUserOrderPermission(db, child.Id, orderId)
			if err != nil {
				return err
			}
		}
	}

	logs := []models.Log{
		{
			OrderId:      orderId,
			ActionTypeId: models.DictActionTypeName["DECLINED"],
			UpdatedBy:    userId,
			Description:  user.Name + " declined order " + order.Name,
		},
	}

	err := repo_order.InsertLogs(db, logs)
	if err != nil {
		return err
	}

	return nil

}

func InsertComments(a *gorm.DB, userId, orderId int, comments []models.Comment) error {
	return repo_order.InsertComments(a, userId, orderId, comments)
}

func GetEquipment(a *gorm.DB) ([]models.OrderEquipment, error) {
	return repo_order.GetEquipment(a)
}

func GetPriority(a *gorm.DB) ([]models.Priority, error) {
	return repo_order.GetPriority(a)
}

func CreateEquipment(a *gorm.DB, name string, quantity int) ([]models.OrderEquipment, error) {
	err := repo_order.CreateEquipment(a, name, quantity)
	if err != nil {
		return nil, err
	}

	equipment, err := repo_order.GetEquipment(a)
	if err != nil {
		return nil, err
	}
	return equipment, nil
}

func AddEquipmentQuantity(a *gorm.DB, equipmentId int, quantity int) error {
	return repo_order.AddEquipmentQuantity(a, equipmentId, quantity)
}

func AddPriority(db *gorm.DB, name string, level int) error {
	return repo_order.AddPriority(db, name, level)
}

func AddEquipment(db *gorm.DB, name string, quantity int) error {
	return repo_order.AddEquipment(db, name, quantity)
}

func ConfirmDelivery(db *gorm.DB, userId int, orderId int) error {
	statusId := models.DictStatusName["GELIEFERT"]
	err := repo_order.ConfirmDelivery(db, userId, orderId, statusId)
	if err != nil {
		return err
	}

	logs := []models.Log{{
		OrderId:   orderId,
		UpdatedBy: userId,
		ActionTypeId: models.DictActionTypeName["CONFIRMED DELIVERY"],
	}}
	err = repo_order.InsertLogs(db, logs)

	return nil
}

func EditOrderEquipment(db *gorm.DB, orderId, equipmentId, quantity int) error {
	return repo_order.EditOrderEquipment(db, orderId, equipmentId, quantity)
}

func EditOrderPriority(db *gorm.DB, orderId, priority int) error {
	return repo_order.EditOrderPriority(db, orderId, priority)
}

func AddDriver(db *gorm.DB, name, description string) error {
	return repo_order.AddDriver(db, name, description)
}

func UpdateEquipment(db *gorm.DB, equipmentId int, quantity int) error {
	return repo_order.UpdateEquipment(db, equipmentId, quantity)
}

func DispatchOrder(db *gorm.DB, userId, orderId, driverId int) error {
	user, _ := GetUserByID(db, userId)
	order, _ := GetOrder(db, userId, orderId)

	var statusId int

	if user.BranchId == models.DictBranchName["Mollnhof"] {
		if order.StatusId == models.DictStatusName["AKZEPTIERT"] {
			statusId = models.DictStatusName["AUF DEM WEG"]
		}
	} else {
		return errors.New("something went wrong")
	}

	if err := repo_order.UpdateOrderStatus(db, orderId, statusId); err != nil {
		return err
	}


	children, err := GetChildren(db, userId)

	for _, child := range *children {
		err = repo_order.DeleteUserOrderPermission(db, child.Id, order.Id)
		if err != nil {
			return err
		}
	}

	if err := repo_order.SetDriverToOrder(db, orderId, driverId); err != nil {
		return err
	}

	logs := []models.Log{
		{
			OrderId:      orderId,
			ActionTypeId: models.DictActionTypeName["ASSIGNED"],
			UpdatedBy:    userId,
			Description:  user.Name + " dispatched order " + order.Name + " and assigned it to the driver",
		},
	}

	err = repo_order.InsertLogs(db, logs)
	if err != nil {
		return err
	}

	return nil
}