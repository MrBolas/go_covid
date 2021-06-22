package models

import "github.com/kyokomi/emoji/v2"

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
