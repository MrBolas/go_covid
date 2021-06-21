package controllers

import (
	"time"

	s "strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func TeleCovidBot() (*tb.Bot, error) {

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: ,

		Token:  "1284369386:AAFL4px7I-31qqs5GnZaV-7TFiVt98bMAXA",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	b.Handle("/covid", func(m *tb.Message) {

		// String split
		var input []string = s.Split(m.Text, " ")
		var country string = s.Join(input[1:len(input)], " ")

		if country == "" {
			country = "portugal"
		}

		// Get Country Data
		var countryData = GetCountryData(country)

		// b.Send(m.Sender, "Hello World!")
		b.Send(m.Sender, countryData.GetReport())
	})

	b.Handle("/history", func(m *tb.Message) {

		// String split
		var input []string = s.Split(m.Text, " ")
		var country string = s.Join(input[1:len(input)], " ")

		if country == "" {
			country = "portugal"
		}

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country)

		// b.Send(m.Sender, "Hello World!")
		b.Send(m.Sender, countryHistoryData.GetReport())
	})

	b.Handle("/subscribe", func(m *tb.Message) {
		//RSTARSTARSTARST
		b.Send(m.Sender, "Subscription Compleeede")
	})

	return b, nil
}
