package repo_stats

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"team2/sandsack-management-backend/models"
)

type TotalNumber struct {
	Count int `gorm:"column:count"`
	Name string `gorm:"column:name"`
}

func TotalNumberUnterabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `select count(o.id)
				from "user" u, "order" o, branch b
				where u.id = o.user_id
				and b.id = u.branch_id
				and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber
	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

type TotalNumberPerAbschnitt struct {
	Count int `gorm:"column:count"`
	UserId int `gorm:"column:user_id"`
	Name string `gorm:"column:name"`
}

func TotalNumberAcceptedUnterabschnitt(db *gorm.DB, startDate, endDate string) int {
	query := `select count(o.id)
			 from "user" u, "order" o, branch b
			 where u.id = o.user_id
			   and o.status_id = (select id from status where name = 'GELIEFERT')
			   and b.id = u.branch_id
			   and b.name = 'Unterabschnitt'
				and o.create_date between ?::timestamp and ?::timestamp;`
	var totalNumber TotalNumber

	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return 0
	}
	return totalNumber.Count
}

func GeneralStatisticsUnterabschnitt(db *gorm.DB, startDate, endDate string) models.GeneralStatistics {
	totalNumber := strconv.Itoa(TotalNumberUnterabschnitt(db, startDate, endDate))
	totalNumberAccepted := strconv.Itoa(TotalNumberAcceptedUnterabschnitt(db, startDate, endDate))
	averageProcessingTime := "10 mins"
	return models.GeneralStatistics{
		TotalNumberOfOrders: totalNumber,
		TotalNumberOfAcceptedOrders: totalNumberAccepted,
		AverageProcessingTime: averageProcessingTime,
	}
}


func CreatedOrderNumber(db *gorm.DB, startDate, endDate string) []TotalNumberPerAbschnitt {
	query := `select u.id, u.name, count(o.id)
			from "user" u, "order" o, branch b
			where o.user_id = u.id
			  and b.id = u.branch_id
			  and b.name = 'Unterabschnitt'
			  and o.create_date between ?::timestamp and ?::timestamp
			group by u.id, u.name;`
	var totalNumber []TotalNumberPerAbschnitt
	if err := db.Raw(query, startDate, endDate).Scan(&totalNumber).Error; err != nil {
		return nil
	}
	return totalNumber
}

func StatisticsPerUnterabschnitt(db *gorm.DB, startDate, endDate string) []models.StatisticsPerAbschnitt {
	var stats []models.StatisticsPerAbschnitt
	totalNumberPerAbschnitt := CreatedOrderNumber(db, startDate, endDate)
	log.Println("TotalNumberPerAbschnitt", totalNumberPerAbschnitt)
	for _, row := range totalNumberPerAbschnitt {
		var stat models.StatisticsPerAbschnitt
		stat.Name = row.Name
		stat.TotalNumberOfOrders = strconv.Itoa(row.Count)
		query := `select u.name as name, count(o.id), b.name as branch_name
				from "user" u, "order" o, branch b
				where u.id = o.user_id
				and o.status_id = (select id from status where name = 'GELIEFERT')
				and b.id = u.branch_id
				and b.name = 'Unterabschnitt'
				and u.id = ?
				and o.create_date between ?::timestamp and ?::timestamp
				group by u.name, b.name;`
		var totalNumber TotalNumber
		err := db.Raw(query, row.UserId, startDate, endDate).Scan(&totalNumber).Error
		if err != nil {
			log.Println("Error in total number of accepted orders", err.Error())
			totalNumber.Count = 0
		}
		log.Println("Total number", row.Name, totalNumber.Count)
		stat.TotalNumberOfAcceptedOrders = strconv.Itoa(totalNumber.Count)
		stats = append(stats, stat)
	}

	return stats
}