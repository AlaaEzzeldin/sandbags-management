package service

import (
	"gorm.io/gorm"
	"team2/sandsack-management-backend/models"
	repo_stats "team2/sandsack-management-backend/repository/stats"
)

func GetUnterabschnittStatistics(db *gorm.DB, startDate, endDate string) (models.UnterabschnittStatistics, error) {
	generalStatistics := repo_stats.GeneralStatisticsUnterabschnitt(db, startDate, endDate)
	statisticsPerUnterabschnitt := repo_stats.StatisticsPerUnterabschnitt(db, startDate, endDate)
	return models.UnterabschnittStatistics{
		Type: "Unterabschnitten",
		GeneralStatistics: generalStatistics,
		StatisticsPerUnterabschnitt: statisticsPerUnterabschnitt,
	}, nil

}


func GetEinsatzabschnittStatistics(db *gorm.DB, startDate, endDate string) (models.EinsatzabschnittStatistics, error) {
	generalStatistics := repo_stats.GeneralStatisticsEinsatzabschnitt(db, startDate, endDate)
	statisticsPerEinsatzabschnitt := repo_stats.StatisticsPerEinsatzabschnitt(db, startDate, endDate)
	return models.EinsatzabschnittStatistics{
		Type: "Einsatzabschnitten",
		GeneralStatistics: generalStatistics,
		StatisticsPerEinsatzabschnitt: statisticsPerEinsatzabschnitt,
	}, nil
}

func GetHauptabschnittStatistics(db *gorm.DB, startDate, endDate string) (models.HauptabschnittStatistics, error) {
	generalStatistics := repo_stats.GeneralStatisticsHauptabschnitt(db, startDate, endDate)
	statisticsPerEinsatzabschnitt := repo_stats.StatisticsPerHauptabschnitt(db, startDate, endDate)
	return models.HauptabschnittStatistics{
		Type: "Hauptabschnitten",
		GeneralStatistics: generalStatistics,
		StatisticsPerHauptabschnitt: statisticsPerEinsatzabschnitt,
	}, nil
}
