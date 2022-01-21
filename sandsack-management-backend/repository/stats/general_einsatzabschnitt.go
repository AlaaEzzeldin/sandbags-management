package repo_stats

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
)

func TotalNumberEinsatzabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `select count(o.id), b.name
			from "user" u, "order" o, "log" l, "action_type" at, branch b
			where u.id = l.updated_by
			  and l.order_id = o.id
			  and l.action_type_id = at.id
			  and (at.name = 'ACCEPTED' or at.name = 'DECLINED')
			  and b.id = u.branch_id
			  and b.name = 'Einsatzabschnitt'
			group by b.name;`
	var totalNumber TotalNumber
	if err := db.Raw(query).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func TotalNumberAcceptedEinsatzabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `select count(o.id), b.name
			from "user" u, "order" o, "log" l, "action_type" at, branch b
			where u.id = l.updated_by
			  and l.order_id = o.id
			  and l.action_type_id = at.id
			  and (at.name = 'ACCEPTED')
			  and b.id = u.branch_id
			  and b.name = 'Einsatzabschnitt'
			group by b.name;`
	var totalNumber TotalNumber

	if err := db.Raw(query).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func GeneralStatisticsEinsatzabschnitt(db *gorm.DB, startDate, endDate string) models.GeneralStatistics {
	totalNumber := strconv.Itoa(TotalNumberEinsatzabschnitt(db, startDate, endDate))
	totalNumberAccepted := strconv.Itoa(TotalNumberAcceptedEinsatzabschnitt(db, startDate, endDate))
	averageProcessingTime := "10 mins"
	return models.GeneralStatistics{
		TotalNumberOfOrders: totalNumber,
		TotalNumberOfAcceptedOrders: totalNumberAccepted,
		AverageProcessingTime: averageProcessingTime,
	}
}

func TotalNumberOrdersEinsatzabschnitt(db *gorm.DB, startDate, endDate string) []TotalNumberPerAbschnitt {
	query := `select u.name, count(o.id), b.name branch_name, u.id user_id
				from "user" u, "order" o, "log" l, "action_type" at, branch b
				where u.id = l.updated_by
				  and l.order_id = o.id
				  and l.action_type_id = at.id
				  and (at.name = 'ACCEPTED' or at.name = 'DECLINED')
				  and b.id = u.branch_id
				  and b.name = 'Einsatzabschnitt'
				group by u.name, b.name, u.id;`
	var totalNumber []TotalNumberPerAbschnitt
	if err := db.Raw(query).Scan(&totalNumber).Error; err != nil {
		return nil
	}
	return totalNumber
}
func StatisticsPerEinsatzabschnitt(db *gorm.DB, startDate, endDate string) []models.StatisticsPerAbschnitt {
	var stats []models.StatisticsPerAbschnitt
	totalNumberPerAbschnitt := TotalNumberOrdersEinsatzabschnitt(db, startDate, endDate)
	log.Println("TotalNumberPerAbschnitt", totalNumberPerAbschnitt)
	for _, row := range totalNumberPerAbschnitt {
		var stat models.StatisticsPerAbschnitt
		stat.Name = row.Name
		stat.TotalNumberOfOrders = strconv.Itoa(row.Count)
		query := `select u.name name, count(o.id), b.name
				from "user" u, "order" o, "log" l, "action_type" at, branch b
				where u.id = l.updated_by
				  and l.order_id = o.id
				  and l.action_type_id = at.id
				  and (at.name = 'ACCEPTED')
				  and b.id = u.branch_id
				  and b.name = 'Einsatzabschnitt'
					and u.id = ?
				group by u.name, b.name;`
		var totalNumber TotalNumber
		err := db.Raw(query, row.UserId).Scan(&totalNumber).Error
		if err != nil {
			totalNumber.Count = 0
		}
		stat.TotalNumberOfAcceptedOrders = strconv.Itoa(totalNumber.Count)
		stats = append(stats, stat)
	}

	return stats

}