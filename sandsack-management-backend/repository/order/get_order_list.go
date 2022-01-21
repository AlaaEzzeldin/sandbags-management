package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetSimpleOrderList(a *gorm.DB, userId int) ([]models.SimpleOrder, error){
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
		return nil, err
	}
	return simpleOrderList, nil
}


func GetAllSimpleOrderList(a *gorm.DB, startDate, endDate string) ([]models.SimpleOrder, error){
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
				order by o.update_date DESC;`

	var simpleOrderList []models.SimpleOrder

	// todo: add startDate, endDate
	if err := a.Raw(query).Scan(&simpleOrderList).Error; err != nil {
		log.Println("GetOrderList err", err.Error())
		return nil, err
	}
	return simpleOrderList, nil
}
