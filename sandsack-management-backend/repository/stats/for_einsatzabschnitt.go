package repo_stats

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
)

func TotalNumberUnterabschnittForEinsatzabschnitt(db *gorm.DB, userId int, startDate, endDate string) int {
	query := `select count(o.id)
				from "order" o
				where o.user_id in (select user2_id from hierarchy where user1_id = ?)
				and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, userId, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func TotalAcceptedNumberUnterabschnittForEinsatzabschnitt(db *gorm.DB, userId int, startDate, endDate string) int {
	query := `select count(o.id)
			from "order" o
			where o.user_id in (select user2_id from hierarchy where user1_id = ?)
			and o.status_id = (select id from status where name = 'GELIEFERT')
			and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, userId, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func AverageProcessingTimeUnterabschnittForEinsatzabschnitt(db *gorm.DB, userId int, startDate, endDate string) int {
	query := `select avg(date_part('minute', o.update_date - o.create_date)) as count
			from "order" o
			where o.user_id in (select user2_id from hierarchy where user1_id = ?)
			  and o.status_id = (select id from status where name = 'GELIEFERT')
			  and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, userId, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	log.Println("AVG proc time unter", totalNumber.Count)
	return totalNumber.Count
}

// todo: for einsatzleiter
func GeneralStatsForEinsatzabschnitt(db *gorm.DB, userId int, startDate, endDate string) models.GeneralStatistics {
	totalNumber := strconv.Itoa(TotalNumberUnterabschnittForEinsatzabschnitt(db, userId, startDate, endDate))
	totalNumberAccepted := strconv.Itoa(TotalAcceptedNumberUnterabschnittForEinsatzabschnitt(db, userId, startDate, endDate))
	averageProcessingTime := AverageProcessingTimeUnterabschnittForEinsatzabschnitt(db, userId, startDate, endDate)
	return models.GeneralStatistics{
		TotalNumberOfOrders: totalNumber,
		TotalNumberOfAcceptedOrders: totalNumberAccepted,
		AverageProcessingTime: strconv.Itoa(averageProcessingTime) + " min",
	}
}

func CreatedOrderNumberByUnterabschnitten(db *gorm.DB, userId int, startDate, endDate string) []TotalNumberPerAbschnitt {
	query := `select u.id as user_id, u.name, count(o.id)
				from "user" u, "order" o, branch b
				where o.user_id = u.id
				  and b.id = u.branch_id
				  and b.name = 'Unterabschnitt'
				  and u.id in (select user2_id from hierarchy where user1_id = ?)
				  and o.create_date between ?::timestamp and ?::timestamp
				group by u.id, u.name;`
	var totalNumber []TotalNumberPerAbschnitt
	if err := db.Raw(query, userId, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		log.Println("CreateOrderNumber error", err.Error())
		return nil
	}
	return totalNumber
}

func UnterabschnittStatsForEinsatzabschnitt(db *gorm.DB, userId int, startDate, endDate string) []models.StatisticsPerAbschnitt {
	var stats []models.StatisticsPerAbschnitt
	totalNumberPerAbschnitt := CreatedOrderNumberByUnterabschnitten(db, userId, startDate, endDate)
	for _, row := range totalNumberPerAbschnitt {
		log.Println("UserId", row)
		var stat models.StatisticsPerAbschnitt
		stat.Name = row.Name
		stat.TotalNumberOfOrders = strconv.Itoa(row.Count)
		query := `select u.name, count(o.id)
				from "user" u, "order" o
				where u.id = o.user_id
				and o.status_id = (select id from status where name = 'GELIEFERT')
				and u.id = ?
				and o.create_date between ?::timestamp and ?::timestamp
				group by u.name;`
		var totalNumber TotalNumber
		err := db.Raw(query, row.UserId, startDate, endDate).Scan(&totalNumber).Error
		if err != nil {
			log.Println("Error in total number of accepted orders", err.Error())
			totalNumber.Count = 0
		}
		stat.TotalNumberOfAcceptedOrders = strconv.Itoa(totalNumber.Count)
		stats = append(stats, stat)
	}
	return stats
}

