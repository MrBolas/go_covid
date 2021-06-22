package models

import (
	"fmt"
	s "strings"

	"github.com/kyokomi/emoji/v2"
)

type CountryInfo struct {
	Id   int     `json:"_id"`
	Iso2 string  `json:"iso2"`
	Iso3 string  `json:"iso3"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
	Flag string  `json:"flag"`
}

type Country struct {
	Updated                int         `json:"updated"`
	Country                string      `json:"country"`
	CountryInfo            CountryInfo `json:"countryInfo"`
	Cases                  int         `json:"cases"`
	TodayCases             int         `json:"todayCases"`
	Deaths                 int         `json:"deaths"`
	TodayDeaths            int         `json:"todayDeaths"`
	Recovered              int         `json:"recovered"`
	TodayRecovered         int         `json:"todayRecovered"`
	Active                 int         `json:"active"`
	Critical               int         `json:"critical"`
	CasesPerOneMillion     float64     `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float64     `json:"deathsPerOneMillion"`
	Tests                  int         `json:"tests"`
	TestsPerOneMillion     float64     `json:"testsPerOneMillion"`
	Population             int         `json:"population"`
	Continent              string      `json:"continent"`
	OneCasePerPeople       float64     `json:"oneCasePerPeople"`
	OneDeathPerPeople      float64     `json:"oneDeathPerPeople"`
	OneTestPerPeople       float64     `json:"oneTestPerPeople"`
	ActivePerOneMillion    float64     `json:"activePerOneMillion"`
	RecoveredPerOneMillion float64     `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float64     `json:"criticalPerOneMillion"`
}

func (c *Country) GetReport() string {
	// Sets fleg
	flag := fmt.Sprintf(":flag-" + s.ToLower(c.CountryInfo.Iso2) + ":")

	// Creates report
	report := emoji.Sprintf("Covid for %s%s\nNew cases 😷: %d\nCases 😷: %d\nNew deaths 💀: %d\nDeaths 💀: %d\nActive 🧟‍♂️: %d\nRecovered 🏃‍♂️: %d\nTests💉: %d",
		c.Country, flag, c.TodayCases, c.Cases, c.TodayDeaths, c.Deaths, c.Active, c.Recovered, c.Tests)

	return report
}
