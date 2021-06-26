package apimodels

import (
	"errors"
	"log"
	"sort"
	"time"

	"go_covid/src/utils"

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

type countryHistoryDatapoint struct {
	key         string
	value       float64
	iso8602Date string
	time        time.Time
}

const layoutUS = "1/2/06"
const layoutISO8601 = "2006-01-02"

func asSortedDataPoint(m map[string]int) ([]countryHistoryDatapoint, error) {
	var dataSet []countryHistoryDatapoint
	for key, val := range m {
		time, err := time.Parse(layoutUS, key)
		if err != nil {
			log.Println("error paring date")
			return nil, err
		}
		iso8602Date := time.Format(layoutISO8601)
		dataSet = append(dataSet, countryHistoryDatapoint{
			key:         key,
			value:       float64(val),
			iso8602Date: iso8602Date,
			time:        time,
		})
	}
	sort.Slice(dataSet, func(i, j int) bool { return dataSet[i].iso8602Date < dataSet[j].iso8602Date })
	return dataSet, nil
}

func asTimeAndValueSeries(dataSet *[]countryHistoryDatapoint) ([]time.Time, []float64) {
	var TimeSeries []time.Time
	var ValueSeries []float64
	for _, data := range *dataSet {
		TimeSeries = append(TimeSeries, data.time)
		ValueSeries = append(ValueSeries, data.value)
	}
	return TimeSeries, ValueSeries
}

func makeTimeAndValueRelative(timeSeries []time.Time, valueSeries []float64) ([]time.Time, []float64, error) {
	if len(valueSeries) < 2 {
		return nil, nil, errors.New("Value series too short, can't make relative")
	}
	currentValue := valueSeries[0]
	var newValueSeries []float64
	for _, val := range valueSeries[1:] {
		newValueSeries = append(newValueSeries, val-currentValue)
		currentValue = val
	}
	return timeSeries[1:], newValueSeries, nil
}

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
	dataSet, err := asSortedDataPoint(c.Cases)
	if err != nil {
		log.Println("error converting data series")
		return nil, nil, err
	}
	timeSeries, valueSeries := asTimeAndValueSeries(&dataSet)
	return timeSeries, valueSeries, err
}

// GetRelativeCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetRelativeCasesTimeSeries() ([]time.Time, []float64, error) {
	timeSeries, valueSeries, err := c.GetCasesTimeSeries()
	if err != nil {
		return nil, nil, err
	}
	return makeTimeAndValueRelative(timeSeries, valueSeries)
}

// GetCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetDeathsTimeSeries() ([]time.Time, []float64, error) {
	dataSet, err := asSortedDataPoint(c.Deaths)
	if err != nil {
		log.Println("error converting data series")
		return nil, nil, err
	}
	timeSeries, valueSeries := asTimeAndValueSeries(&dataSet)
	return timeSeries, valueSeries, err
}

// GetCasesTimeSeries Seperates a Data set into two different slices for each of the variables in the data set.
func (c CountryTimeline) GetRelativeDeathsTimeSeries() ([]time.Time, []float64, error) {
	timeSeries, valueSeries, err := c.GetDeathsTimeSeries()
	if err != nil {
		return nil, nil, err
	}
	return makeTimeAndValueRelative(timeSeries, valueSeries)
}
