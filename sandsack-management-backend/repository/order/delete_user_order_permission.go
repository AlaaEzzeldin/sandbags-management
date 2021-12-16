package repo_order

import "gorm.io/gorm"

func DeleteUserOrderPermission(a *gorm.DB, userId, orderId int) error {
	query := `delete from user_order_permission 
			where user_id = ? and order_id = ? and 
			permission_id in 
				(select id from permission 
				where name = 'CAN EDIT' or name = 'CAN DECLINE' or name = 'CAN ACCEPT');`
	err := a.Exec(query, userId, orderId).Error
	if err != nil {
		return err
	}
	return nil
}
