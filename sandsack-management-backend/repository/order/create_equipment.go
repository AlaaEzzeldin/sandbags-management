package repo_order

import "gorm.io/gorm"

func CreateEquipment(a *gorm.DB, name string, quantity int) error {
	query := `insert into equipment(name, quantity) values(?,?);`
	err := a.Exec(query, name, quantity).Error
	return err
}
