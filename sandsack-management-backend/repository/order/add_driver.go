package repo_order

import "gorm.io/gorm"

func AddDriver(db *gorm.DB, name, description string) error {
	query := `insert into public.driver(name, description)values(?,?);`
	if err := db.Exec(query, name, description).Error; err != nil {
		return err
	}
	return nil
}
