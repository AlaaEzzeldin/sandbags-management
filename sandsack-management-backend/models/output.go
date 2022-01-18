package models


type SendVerifyMailOutput struct {
	OTP string `json:"otp,omitempty"`
}

type Tokens struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	Role         string `json:"role,omitempty"`
	Name         string `json:"name,omitempty"`
}

type GeneralStatistics struct {
	TotalNumberOfOrders string `json:"total_number_of_orders"`
	TotalNumberOfAcceptedOrders string `json:"total_number_of_accepted_orders"`
	AverageProcessingTime string `json:"average_processing_time"`
}

type StatisticsPerAbschnitt struct {
	Name string `json:"name"`
	TotalNumberOfOrders string `json:"total_number_of_orders"`
	TotalNumberOfAcceptedOrders string `json:"total_number_of_accepted_orders"`
}

type GetStatisticsOutput struct {
	Type string `json:"type"`
	GeneralStatistics GeneralStatistics `json:"general_statistics"`
	StatisticsPerAbschnitt []StatisticsPerAbschnitt
}