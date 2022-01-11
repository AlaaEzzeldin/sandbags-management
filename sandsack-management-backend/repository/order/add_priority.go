package repo_order

import "gorm.io/gorm"

func AddPriority(db *gorm.DB, name string, level int) error {
	query := `insert into priority(name, level) values(?,?);`
	return db.Exec(query, name, level).Error

}


func AddEquipment(db *gorm.DB, name string, quantity int) error {
	query := `insert into equipment(name, quantity) values(?,?);`
	return db.Exec(query, name, quantity).Error
}