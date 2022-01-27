package repo_stats

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
)

func GetEinsatzabschnitten(db *gorm.DB, userId int) []models.User {
	query := `select u.id, u.name
				from "user" u
				where u.id in (select user2_id from hierarchy where user1_id = ?)`
	var users []models.User
	if err := db.Raw(query, userId).Scan(&users).Error; err != nil {
		return nil
	}
	return users
}

func TotalNumberForHauptabschnitt(db *gorm.DB, userId int, startDate, endDate string) int {
	total := 0
	for _, row := range GetEinsatzabschnitten(db, userId) {
		num := TotalNumberUnterabschnittForEinsatzabschnitt(db, row.Id, startDate, endDate)
		total += num
	}
	return total
}


func TotalNumberAcceptedForHauptabschnitt(db *gorm.DB, userId int, startDate, endDate string) int {
	total := 0
	for _, row := range GetEinsatzabschnitten(db, userId) {
		num := TotalAcceptedNumberUnterabschnittForEinsatzabschnitt(db, row.Id, startDate, endDate)
		total += num
	}
	return total
}

func AverageProcessingTimeForHauptabschnitt(db *gorm.DB, userId int, startDate, endDate string) int {
	query := `select avg(date_part('minute', o.update_date - o.create_date)) as count
			from "order" o
			where o.user_id in (select user2_id from hierarchy where user1_id in (select user2_id from hierarchy where user1_id = ?))
			  and o.status_id = (select id from status where name = 'GELIEFERT')
			  and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, userId, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func GeneralStatsForHauptabschnitt(db *gorm.DB, userId int, startDate, endDate string) models.GeneralStatistics {
	totalNumber := strconv.Itoa(TotalNumberForHauptabschnitt(db, userId, startDate, endDate))
	totalNumberAccepted := strconv.Itoa(TotalNumberAcceptedForHauptabschnitt(db, userId, startDate, endDate))
	averageProcessingTime := AverageProcessingTimeForHauptabschnitt(db, userId, startDate, endDate)
	return models.GeneralStatistics{
		TotalNumberOfOrders: totalNumber,
		TotalNumberOfAcceptedOrders: totalNumberAccepted,
		AverageProcessingTime: strconv.Itoa(averageProcessingTime) + " min",
	}
}



func CreatedOrderNumberByHauptabschnitten(db *gorm.DB, userId int, startDate, endDate string) []TotalNumberPerAbschnitt {
	log.Println("USER_ID", userId)
	query := `with einsatzabschnitten as (
					select u.id, u.name, b.name as branch_name
					from "user" u, branch b
					where u.id in (select user2_id from hierarchy where user1_id = ?)
					and b.id = u.branch_id
				)
				select count(o.id), einsatzabschnitten.name, einsatzabschnitten.id as user_id
				from "user" u, einsatzabschnitten, hierarchy, "order" o
				where u.id = hierarchy.user2_id
				and hierarchy.user1_id = einsatzabschnitten.id
				and o.user_id = u.id
				and o.create_date between ?::timestamp and ?::timestamp
				group by einsatzabschnitten.name, einsatzabschnitten.branch_name, einsatzabschnitten.id;`

	var totalNumber []TotalNumberPerAbschnitt
	if err := db.Raw(query, userId, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		log.Println("CreateOrderNumber error", err.Error())
		return nil
	}
	return totalNumber
}

func EinsatzabschnittStatsForHauptabschnitt(db *gorm.DB, userId int, startDate, endDate string) []models.StatisticsPerAbschnitt {
	var stats []models.StatisticsPerAbschnitt
	totalNumberPerAbschnitt := CreatedOrderNumberByHauptabschnitten(db, userId, startDate, endDate)
	for _, row := range totalNumberPerAbschnitt {
		log.Println("UserId", row)
		var stat models.StatisticsPerAbschnitt
		stat.Name = row.Name
		stat.TotalNumberOfOrders = strconv.Itoa(row.Count)
		query := `with einsatzabschnitten as (
						select u.id, u.name, b.name as branch_name
						from "user" u, branch b
						where b.id = u.branch_id
						and b.name = 'Einsatzabschnitt'
					)
					select count(o.id), einsatzabschnitten.name
					from "order" o, "user" u, hierarchy h, einsatzabschnitten
					where u.id = o.user_id
					and h.user1_id = einsatzabschnitten.id
					and h.user2_id = u.id
					and einsatzabschnitten.id = ?
					and o.user_id = u.id
					and o.status_id = (select id from status where name = 'GELIEFERT')
					and o.create_date between ?::timestamp and ?::timestamp
					group by einsatzabschnitten.name;`
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

