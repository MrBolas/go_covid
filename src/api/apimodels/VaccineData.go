package apimodels

import (
	"fmt"
	"log"
	"sort"
	s "strings"

	"go_covid/src/utils"

	"github.com/kyokomi/emoji/v2"
)

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

func (v *VaccineCountryData) GetReport() string {
	// Sets fleg
	country, _ := utils.SearchCountry(v.Country)
	flag := fmt.Sprintf(":flag-" + s.ToLower(country.Codes.Alpha2) + ":")
	timeline := v.Timeline

	sort.Slice(timeline, func(i, j int) bool { return timeline[i].Date < timeline[j].Date })
	last := timeline[1]
	prev := timeline[0]
	log.Println(timeline)
	total := last.TotalPerHundred
	if total == 0 {
		total = prev.TotalPerHundred
	}

	// Creates report
	report := emoji.Sprintf("Vaccines for %s%s\nNew 💉: %d(%d)\nTotal 💉💉: %d(%d)\nTotal(%%) 🏃‍♂️: %2.0f\n",
		country.Name.BaseLang.Common, flag, last.Daily, prev.Daily, last.Total, last.Total-prev.Total, total)

	return report
}
