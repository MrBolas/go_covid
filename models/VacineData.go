package models

type VaccineTimeline struct {
	Total           int     `json:"total"`
	Daily           int     `json:"daily"`
	TotalPerHundred float64 `json:"totalPerHundred"`
	DailyPerMillion float64 `json:"dailyPerMillion"`
	Date            string  `json:"date"`
}

type VaccineCountryData struct {
	Country  string            `json:"country"`
	Timeline []VaccineTimeline `json:"timeline"`
}
