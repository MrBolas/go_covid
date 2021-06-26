package apimodels

import (
	"time"

	"github.com/kyokomi/emoji/v2"
	"go_covid/src/utils"
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

const layoutUS = "1/2/06"

// GetReport returns a report in string format
func (c *CountryHistory) GetReport() string {

	sortedCasesKeys := utils.GetSortedKeys(c.Timeline.Cases)
	sortedDeathsKeys := utils.GetSortedKeys(c.Timeline.Deaths)
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

// GetCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetCasesTimeSeries() ([]time.Time, []float64, error) {

	// Order cases by date
	var orderedHistoryData []string = utils.GetSortedKeys(c.Cases)

	// Define return slices
	var TimeSeries []time.Time
	var ValueSeries []float64

	// Goes for Dates and separates two sets of data.
	// TimeSeries are the dates.
	// ValueSeries are the corresponding values.
	for _, date := range orderedHistoryData {
		timeEvent, err := time.Parse(layoutUS, date)

		// Appends Dates
		TimeSeries = append(TimeSeries, timeEvent)

		// Appends Date Values
		ValueSeries = append(ValueSeries, float64(c.Cases[date]))

		if err != nil {
			return nil, nil, err
		}
	}

	return TimeSeries, ValueSeries, nil
}

// GetRelativeCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetRelativeCasesTimeSeries() ([]time.Time, []float64, error) {

	// Order cases by date
	var orderedHistoryData []string = utils.GetSortedKeys(c.Cases)

	// Define return slices
	var TimeSeries []time.Time
	var ValueSeries []float64

	// Goes for Dates and separates two sets of data.
	// TimeSeries are the dates.
	// ValueSeries are the corresponding values.
	for dateIndex, date := range orderedHistoryData {
		timeEvent, err := time.Parse(layoutUS, date)

		// Defines a relative value to the previous
		newValue := float64(c.Cases[date])
		if dateIndex > 1 {
			newValue -= float64(c.Cases[orderedHistoryData[dateIndex-1]])

			// Appends Dates
			TimeSeries = append(TimeSeries, timeEvent)

			// Appends Date Values
			ValueSeries = append(ValueSeries, newValue)
		}

		if err != nil {
			return nil, nil, err
		}
	}

	return TimeSeries, ValueSeries, nil
}

// GetCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetDeathsTimeSeries() ([]time.Time, []float64, error) {

	// Order cases by date
	var orderedHistoryData []string = utils.GetSortedKeys(c.Deaths)

	// Define return slices
	var TimeSeries []time.Time
	var ValueSeries []float64

	// Goes for Dates and separates two sets of data.
	// TimeSeries are the dates.
	// ValueSeries are the corresponding values.
	for _, date := range orderedHistoryData {
		timeEvent, err := time.Parse(layoutUS, date)

		// Appends Dates
		TimeSeries = append(TimeSeries, timeEvent)

		// Appends Date Values
		ValueSeries = append(ValueSeries, float64(c.Deaths[date]))

		if err != nil {
			return nil, nil, err
		}
	}

	return TimeSeries, ValueSeries, nil
}

// GetCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetRelativeDeathsTimeSeries() ([]time.Time, []float64, error) {

	// Order cases by date
	var orderedHistoryData []string = utils.GetSortedKeys(c.Deaths)

	// Define return slices
	var TimeSeries []time.Time
	var ValueSeries []float64

	// Goes for Dates and separates two sets of data.
	// TimeSeries are the dates.
	// ValueSeries are the corresponding values.
	for dateIndex, date := range orderedHistoryData {
		timeEvent, err := time.Parse(layoutUS, date)

		// Defines a relative value to the previous
		newValue := float64(c.Deaths[date])
		if dateIndex > 1 {
			newValue -= float64(c.Deaths[orderedHistoryData[dateIndex-1]])

			// Appends Dates
			TimeSeries = append(TimeSeries, timeEvent)

			// Appends Date Values
			ValueSeries = append(ValueSeries, newValue)
		}

		if err != nil {
			return nil, nil, err
		}
	}

	return TimeSeries, ValueSeries, nil
}
