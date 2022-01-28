package repo_order

import (
	"gorm.io/gorm"
	"log"
	"team2/sandsack-management-backend/models"
)

func GetSimpleOrder(a *gorm.DB, userId, orderId int) (models.SimpleOrder, error){
	query := `select o.id,
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
		  and o.user_id = u.id
		  and o.id = ?;	`
	var simpleOrder models.SimpleOrder

	err := a.Raw(query, orderId).Scan(&simpleOrder).Error
	if err != nil {
		log.Println("GetOrder error", err.Error())
		return models.SimpleOrder{}, err
	}
	return simpleOrder, nil
}


