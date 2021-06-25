package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"
	"time"

	chart "github.com/wcharczuk/go-chart/v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

func TeleCovidBot(token string) (*tb.Bot, error) {

	// Create new bot
	b, err := tb.NewBot(tb.Settings{
		Token: token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	// /covid handler. Accepts /covid <country>. Default country is Portugal
	b.Handle("/covid", func(m *tb.Message) {

		// String split
		var input []string = s.Fields(m.Text)
		var country string = "Portugal"

		if len(input) > 1 {
			country = s.Join(input[1:], " ")
		}

		// Get Country Data
		var countryData = GetCountryData(country)

		// b.Send(m.Sender, "Hello World!")
		b.Send(m.Chat, countryData.GetReport())
	})

	// /history handler. Accepts /history <country>. Default country is Portugal
	b.Handle("/history", func(m *tb.Message) {

		// String split
		var input []string = s.Fields(m.Text)
		var country string = "Portugal"

		if len(input) > 1 {
			country = s.Join(input[1:], " ")
		}

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country, 3)

		// b.Send(m.Sender, "Hello World!")
		b.Send(m.Chat, countryHistoryData.GetReport())
	})

	// /cases handler. Accepts /cases <country> <days>. Default country is Portugal. Default number of days is 7.
	b.Handle("/cases", func(m *tb.Message) {

		// String split
		var input []string = s.Fields(m.Text)

		// Parse Input string
		country, days := parseInputs(input)

		fmt.Printf("Cases Chart -> country: %s days: %d\n", country, days)

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country, days)

		// Transform Data into TimeSeries
		//timeseries, valueseries, err := countryHistoryData.Timeline.GetCasesTimeSeries()
		timeseries, valueseries, err := countryHistoryData.Timeline.GetCasesTimeSeries()

		if err != nil {
			log.Fatal(err)
		}

		// Build graph
		covidGraph := chart.Chart{
			Series: []chart.Series{
				chart.TimeSeries{
					XValues: timeseries,
					YValues: valueseries,
				},
			},
		}

		// Create image file and render image
		f, _ := os.Create("assets/covid-cases-graph.png")
		defer f.Close()
		covidGraph.Render(chart.PNG, f)

		// Upload graph image
		g := &tb.Photo{File: tb.FromDisk("assets/covid-cases-graph.png")}

		// Send image
		b.Send(m.Chat, g)
	})

	// /casesR handler. Accepts /casesR <country> <days>. Default country is Portugal. Default number of days is 7.
	b.Handle("/casesR", func(m *tb.Message) {

		// String split
		var input []string = s.Fields(m.Text)

		// Parse Input string
		country, days := parseInputs(input)

		fmt.Printf("Cases Chart -> country: %s days: %d\n", country, days)

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country, days)

		// Transform Data into TimeSeries
		//timeseries, valueseries, err := countryHistoryData.Timeline.GetCasesTimeSeries()
		timeseries, valueseries, err := countryHistoryData.Timeline.GetRelativeCasesTimeSeries()

		if err != nil {
			log.Fatal(err)
		}

		// Build graph
		covidGraph := chart.Chart{
			Series: []chart.Series{
				chart.TimeSeries{
					XValues: timeseries,
					YValues: valueseries,
				},
			},
		}

		// Create image file and render image
		f, _ := os.Create("assets/covid-cases-graph.png")
		defer f.Close()
		covidGraph.Render(chart.PNG, f)

		// Upload graph image
		g := &tb.Photo{File: tb.FromDisk("assets/covid-cases-graph.png")}

		// Send image
		b.Send(m.Chat, g)
	})

	// /deaths handler. Accepts /deaths <country> <days>. Default country is Portugal. Default number of days is 7.
	b.Handle("/deaths", func(m *tb.Message) {

		// String split
		var input []string = s.Fields(m.Text)

		// Parse Input string
		country, days := parseInputs(input)

		fmt.Printf("Deaths Chart -> country: %s days: %d\n", country, days)

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country, days)

		// Transform Data into TimeSeries
		timeseries, valueseries, err := countryHistoryData.Timeline.GetDeathsTimeSeries()

		if err != nil {
			log.Fatal(err)
		}

		// Build graph
		covidGraph := chart.Chart{
			Series: []chart.Series{
				chart.TimeSeries{
					XValues: timeseries,
					YValues: valueseries,
				},
			},
		}

		// Create image file and render image
		f, _ := os.Create("assets/covid-death-graph.png")
		defer f.Close()
		covidGraph.Render(chart.PNG, f)

		// Upload graph image
		g := &tb.Photo{File: tb.FromDisk("assets/covid-death-graph.png")}

		// Send image
		b.Send(m.Chat, g)
	})

	// /deathsR handler. Accepts /deathsR <country> <days>. Default country is Portugal. Default number of days is 7.
	b.Handle("/deathsR", func(m *tb.Message) {

		// String split
		var input []string = s.Fields(m.Text)

		// Parse Input string
		country, days := parseInputs(input)

		fmt.Printf("Deaths Chart -> country: %s days: %d\n", country, days)

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country, days)

		// Transform Data into TimeSeries
		timeseries, valueseries, err := countryHistoryData.Timeline.GetRelativeDeathsTimeSeries()

		if err != nil {
			log.Fatal(err)
		}

		// Build graph
		covidGraph := chart.Chart{
			Series: []chart.Series{
				chart.TimeSeries{
					XValues: timeseries,
					YValues: valueseries,
				},
			},
		}

		// Create image file and render image
		f, _ := os.Create("assets/covid-death-graph.png")
		defer f.Close()
		covidGraph.Render(chart.PNG, f)

		// Upload graph image
		g := &tb.Photo{File: tb.FromDisk("assets/covid-death-graph.png")}

		// Send image
		b.Send(m.Chat, g)
	})

	// /subscribe handler.
	b.Handle("/subscribe", func(m *tb.Message) {

		b.Send(m.Sender, "Subscription Compleeede")
	})

	return b, nil
}

func parseInputs(input []string) (string, int) {

	// Default Values
	var country string = "Portugal"
	var days int = 7
	var err error

	// Verifies last slice element has decimal
	if len(input) > 1 {
		days, err = strconv.Atoi(input[len(input)-1])

		if err == nil {
			input = input[:len(input)-1]
		} else {
			// Reset to default
			days = 7
		}
	}

	// Reads country from inputs
	if len(input) > 1 {
		country = s.Join(input[1:], " ")
	}

	return country, days
}
