package controllers

import (
	"time"

	s "strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func TeleCovidBot() (*tb.Bot, error) {

	// Create new bot
	b, err := tb.NewBot(tb.Settings{
		Token:  "1284369386:AAFL4px7I-31qqs5GnZaV-7TFiVt98bMAXA",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	// /covid handler. Accepts /covid <country>. Default country is Portugal
	b.Handle("/covid", func(m *tb.Message) {

		// String split
		var input []string = s.Split(m.Text, " ")
		var country string = s.Join(input[1:], " ")

		if country == "" {
			country = "portugal"
		}

		// Get Country Data
		var countryData = GetCountryData(country)

		// b.Send(m.Sender, "Hello World!")
		b.Send(m.Sender, countryData.GetReport())
	})

	// /history handler. Accepts /history <country>. Default country is Portugal
	b.Handle("/history", func(m *tb.Message) {

		// String split
		var input []string = s.Split(m.Text, " ")
		var country string = s.Join(input[1:], " ")

		if country == "" {
			country = "portugal"
		}

		// Get Country Data
		var countryHistoryData = GetHistoricalCountryData(country)

		// b.Send(m.Sender, "Hello World!")
		b.Send(m.Sender, countryHistoryData.GetReport())
	})

	// /subscribe handler.
	b.Handle("/subscribe", func(m *tb.Message) {

		b.Send(m.Sender, "Subscription Compleeede")
	})

	return b, nil
}
