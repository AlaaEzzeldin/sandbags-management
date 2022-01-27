package repo_stats

import (
	"gorm.io/gorm"
	"strconv"
	"team2/sandsack-management-backend/models"
)

func GetHauptabschnitten(db *gorm.DB) []models.User {
	query := `select u.id, u.name
				from "user" u, branch b
				where u.branch_id = b.id
				and b.name = 'Hauptabschnitt';`
	var users []models.User
	if err := db.Raw(query).Scan(&users).Error; err != nil {
		return nil
	}
	return users
}

func TotalNumberHauptabschnitt(db *gorm.DB, startDate, endDate string) int {
	total := 0
	for _, row := range GetHauptabschnitten(db) {
		num := TotalNumberForHauptabschnitt(db, row.Id, startDate, endDate)
		total += num
	}
	return total
}

func TotalNumberAcceptedHauptabschnitt(db *gorm.DB, startDate, endDate string) int {
	total := 0
	for _, row := range GetHauptabschnitten(db) {
		num := TotalNumberAcceptedForHauptabschnitt(db, row.Id, startDate, endDate)
		total += num
	}
	return total
}

func AverageProcessingTimeHauptabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `select avg(date_part('minute', o.update_date - o.create_date)) as count
				from "order" o
				where o.user_id in (
					select user2_id from hierarchy where user1_id in (
						select user2_id from hierarchy where user1_id in (select u.id from "user" u, branch b where u.branch_id = b.id and b.name = 'Hauptabschnitt')))
				  and o.status_id = (select id from status where name = 'GELIEFERT')
				and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func GeneralStatisticsHauptabschnitt(db *gorm.DB, startDate, endDate string) models.GeneralStatistics {
	totalNumber := strconv.Itoa(TotalNumberHauptabschnitt(db, startDate, endDate))
	totalNumberAccepted := strconv.Itoa(TotalNumberAcceptedHauptabschnitt(db, startDate, endDate))
	averageProcessingTime := strconv.Itoa(AverageProcessingTimeHauptabschnitt(db, startDate, endDate))
	return models.GeneralStatistics{
		TotalNumberOfOrders:         totalNumber,
		TotalNumberOfAcceptedOrders: totalNumberAccepted,
		AverageProcessingTime:       averageProcessingTime,
	}
}

func TotalNumberOrdersHauptabschnitt(db *gorm.DB, startDate, endDate string) []TotalNumberPerAbschnitt {
	query := `with hauptabschnitten as (
					select u.id, u.name, b.name as branch_name
					from "user" u, branch b
					where u.id in (
						select u.id from "user" u, branch b
						where b.id = u.branch_id and b.name = 'Hauptabschnitt')
					  and b.id = u.branch_id
				)
				select count(o.id), hpt.name, hpt.id as user_id
				from "order" o, "user" u, hierarchy, hauptabschnitten hpt
				where u.id = o.user_id
				and u.id = hierarchy.user2_id
				and hierarchy.user1_id in (select user2_id from hierarchy where user1_id = hpt.id)
				and o.create_date between ?::timestamp and ?::timestamp
				group by hpt.name, hpt.id;`
	var totalNumber []TotalNumberPerAbschnitt
	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return nil
	}
	return totalNumber
}
func StatisticsPerHauptabschnitt(db *gorm.DB, startDate, endDate string) []models.StatisticsPerAbschnitt {
	var stats []models.StatisticsPerAbschnitt
	totalNumberPerAbschnitt := TotalNumberOrdersHauptabschnitt(db, startDate, endDate)
	for _, row := range totalNumberPerAbschnitt {
		var stat models.StatisticsPerAbschnitt
		stat.Name = row.Name
		stat.TotalNumberOfOrders = strconv.Itoa(row.Count)
		query := `with hauptabschnitten as (
						select u.id, u.name, b.name as branch_name
						from "user" u, branch b
						where u.id in (
							select u.id from "user" u, branch b
							where b.id = u.branch_id and b.name = 'Hauptabschnitt')
						  and b.id = u.branch_id
					)
					
					select count(o.id), hpt.name, hpt.id as user_id
					from "order" o, "user" u, hierarchy, hauptabschnitten hpt
					where u.id = o.user_id
					and u.id = hierarchy.user2_id
					and o.status_id = (select id from status where name = 'GELIEFERT')
					and hierarchy.user1_id in (select user2_id from hierarchy where user1_id = hpt.id)
					and hpt.id = ?
					and o.create_date between ?::timestamp and ?::timestamp
					group by hpt.name, hpt.id;`
		var totalNumber TotalNumber
		err := db.Raw(query, row.UserId, startDate, endDate).Scan(&totalNumber).Error
		if err != nil {
			totalNumber.Count = 0
		}
		stat.TotalNumberOfAcceptedOrders = strconv.Itoa(totalNumber.Count)
		stats = append(stats, stat)
	}

	return stats

}
