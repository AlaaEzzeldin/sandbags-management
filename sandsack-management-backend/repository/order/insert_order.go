package repo_order

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
)

type Id struct {
	Id int `gorm:"column:id"`
}

func CreateOrder(a *gorm.DB, order *models.Order) (int, error) {
	orderId, err := insertOrder(a, order)
	log.Println("OrderID", orderId)
	if err != nil {
		return 0, err
	}

	orderName := order.Name + " #" + strconv.Itoa(orderId)
	log.Println("OrderName", orderName)
	if err := updateOrderName(a, orderId, orderName); err != nil {
		log.Println("CreateOrder error", err.Error())
		return orderId, nil
	}


	return orderId, nil
}

func insertOrder(a *gorm.DB, order *models.Order) (int, error) {
	query := `insert into public.order(name, user_id, priority_id, status_id, address_to, address_from)values(?,?,?,?,?,?) 
				returning id;`
	var id Id
	err := a.Raw(query, order.Name, order.UserId, order.PriorityId, order.StatusId,
		order.AddressTo, order.AddressFrom).Scan(&id).Error
	if err != nil {
		log.Println("InsertOrder error", err.Error())
		return 0, err
	}
	return id.Id, nil
}


func updateOrderName(a *gorm.DB, orderId int, orderName string) error {
	query := `update public.order set name = ?, update_date = now() where id = ?;`
	err := a.Exec(query, orderName, orderId).Error
	if err != nil {
		log.Println("Updating name of order error", err.Error())
		return err
	}
	return nil
}