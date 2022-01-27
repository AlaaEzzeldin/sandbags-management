package repo_stats

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
)

func GetAllEinsatzabschnitten(db *gorm.DB) []models.User {
	query := `select u.id, u.name
				from "user" u, branch b
				where u.branch_id = b.id
				and b.name = 'Einsatzabschnitt'`
	var users []models.User
	if err := db.Raw(query).Scan(&users).Error; err != nil {
		return nil
	}
	return users
}

func TotalNumberEinsatzabschnitt(db *gorm.DB, startDate, endDate string) int {
	total := 0
	for _, row := range GetAllEinsatzabschnitten(db) {
		num := TotalNumberUnterabschnittForEinsatzabschnitt(db, row.Id, startDate, endDate)
		total += num
	}
	return total
}

func TotalNumberAcceptedEinsatzabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `with einsatzabschnitten as (
					select u.id, u.name, b.name as branch_name
					from "user" u, branch b
					where u.id in  (select u.id from "user" u, branch b where b.id = u.branch_id and b.name = 'Einsatzabschnitt')
					  and b.id = u.branch_id
				)
				select count(o.id)
				from "user" u, einsatzabschnitten, hierarchy, "order" o
				where u.id = hierarchy.user2_id
				  and o.status_id = (select id from status where name = 'GELIEFERT')
				  and hierarchy.user1_id = einsatzabschnitten.id
				  and o.user_id = u.id;`
	var totalNumber TotalNumber

	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func AverageProcessingTimeEinsatzabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `select avg(date_part('minute', o.update_date - o.create_date)) as count
				from "order" o
				where o.user_id in (
					select user2_id from hierarchy where user1_id in (
						select u.id from "user" u, branch b where b.id = u.branch_id and b.name = 'Einsatzabschnitt'))
				  and o.status_id = (select id from status where name = 'GELIEFERT')
				and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func GeneralStatisticsEinsatzabschnitt(db *gorm.DB, startDate, endDate string) models.GeneralStatistics {
	totalNumber := strconv.Itoa(TotalNumberEinsatzabschnitt(db, startDate, endDate))
	totalNumberAccepted := strconv.Itoa(TotalNumberAcceptedEinsatzabschnitt(db, startDate, endDate))
	averageProcessingTime := strconv.Itoa(AverageProcessingTimeEinsatzabschnitt(db, startDate, endDate)) + " min"
	return models.GeneralStatistics{
		TotalNumberOfOrders:         totalNumber,
		TotalNumberOfAcceptedOrders: totalNumberAccepted,
		AverageProcessingTime:       averageProcessingTime,
	}
}

func TotalNumberOrdersEinsatzabschnitt(db *gorm.DB, startDate, endDate string) []TotalNumberPerAbschnitt {
	query := `with einsatzabschnitten as (
				select u.id, u.name, b.name as branch_name
				from "user" u, branch b
				where u.id in  (select u.id from "user" u, branch b where b.id = u.branch_id and b.name = 'Einsatzabschnitt')
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
	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return nil
	}
	return totalNumber
}

func StatisticsPerEinsatzabschnitt(db *gorm.DB, startDate, endDate string) []models.StatisticsPerAbschnitt {
	var stats []models.StatisticsPerAbschnitt
	totalNumberPerAbschnitt := TotalNumberOrdersEinsatzabschnitt(db, startDate, endDate)
	log.Println("TotalNumberPerAbschnitt Einsatz", totalNumberPerAbschnitt)
	for _, row := range totalNumberPerAbschnitt {
		var stat models.StatisticsPerAbschnitt
		stat.Name = row.Name
		stat.TotalNumberOfOrders = strconv.Itoa(row.Count)
		query := `with einsatzabschnitten as (
						select u.id, u.name, b.name as branch_name
						from "user" u, branch b
						where u.id in  (select u.id from "user" u, branch b where b.id = u.branch_id and b.name = 'Einsatzabschnitt')
						  and b.id = u.branch_id
					)
					select count(o.id), einsatzabschnitten.name, einsatzabschnitten.id as user_id
					from "user" u, einsatzabschnitten, hierarchy, "order" o
					where u.id = hierarchy.user2_id
					  and o.status_id = (select id from status where name = 'GELIEFERT')
					  and hierarchy.user1_id = einsatzabschnitten.id
					  and einsatzabschnitten.id = ?
					  and o.user_id = u.id
					and o.create_date between ?::timestamp and ?::timestamp
					group by einsatzabschnitten.name, einsatzabschnitten.branch_name, einsatzabschnitten.id;`
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
