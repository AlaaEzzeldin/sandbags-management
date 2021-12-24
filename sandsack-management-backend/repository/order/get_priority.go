package repo_order

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
)

func GetPriority(a *gorm.DB) ([]models.Priority, error) {
	query := `select id, level, name from priority order by level;`
	var priorities []models.Priority
	if err := a.Raw(query).Scan(&priorities).Error; err != nil {
		return nil, err
	}
	return priorities, nil
}
