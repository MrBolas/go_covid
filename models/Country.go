package models

import (
	"fmt"

	"sort"
	s "strings"

	"github.com/kyokomi/emoji/v2"
)

type CountryTimeline struct {
	Cases     map[string]int `json:"cases"`
	Deaths    map[string]int `json:"deaths"`
	Recovered map[string]int `json:"recovered"`
}

type CountryHistory struct {
	Country  string   `json:"country"`
	Province []string `json:"province"`
	Timeline CountryTimeline
}

func (c *CountryHistory) GetReport() string {

	sortedCasesKeys := getSortedKeys(c.Timeline.Cases)
	sortedDeathsKeys := getSortedKeys(c.Timeline.Deaths)
	//sortedRecoveredKeys := getSortedKeys(c.Timeline.Recovered)

	todayCases := c.Timeline.Cases[sortedCasesKeys[2]] - c.Timeline.Cases[sortedCasesKeys[1]]
	yesterdayCases := c.Timeline.Cases[sortedCasesKeys[1]] - c.Timeline.Cases[sortedCasesKeys[0]]

	todayDeaths := c.Timeline.Deaths[sortedDeathsKeys[2]] - c.Timeline.Deaths[sortedDeathsKeys[1]]
	yesterdayDeaths := c.Timeline.Deaths[sortedDeathsKeys[1]] - c.Timeline.Deaths[sortedDeathsKeys[0]]

	// Creates report
	report := emoji.Sprintf("Covid for %s\nNew cases 😷: %d(%d)\nNew deaths 💀: %d(%d)\n",
		c.Country, todayCases, yesterdayCases, todayDeaths, yesterdayDeaths)

	return report
}

func getSortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	// values := make([]int, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
		// values = append(values, v)
	}

	sort.Strings(keys)

	return keys
}

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
