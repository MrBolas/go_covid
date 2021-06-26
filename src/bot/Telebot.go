package bot

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type ExtendedCommand struct {
	Command tb.Command
	Handler func(*tb.Message)
}

func TeleCovidBot(token string) (*tb.Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	commands := []ExtendedCommand{
		{Command: tb.Command{
			Text:        "/covid",
			Description: "Gives Covid Numbers for <Country>. Portugal is Default."},
			Handler: onCovid,
		},
		{Command: tb.Command{
			Text:        "/history",
			Description: "Gives Covid History for <Contry>. Portugal is Default."},
			Handler: onHistory,
		},
		{Command: tb.Command{
			Text:        "/cases",
			Description: "Gives a chart with daily cases for the last <Days> for the given <Country>."},
			Handler: onCases,
		},
		{Command: tb.Command{
			Text:        "/new_cases",
			Description: "Gives a chart with daily new cases for the last <Days> for the given <Country>."},
			Handler: onNewCases,
		},
		{Command: tb.Command{
			Text:        "/deaths",
			Description: "Gives a chart with daily deaths for the last <Days> for the given <Country>."},
			Handler: onDeaths,
		},
		{Command: tb.Command{
			Text:        "/new_deaths",
			Description: "Gives a chart with daily new deaths for the last <Days> for the given <Country>."},
			Handler: onNewDeaths,
		},
		{Command: tb.Command{
			Text:        "/subscribe",
			Description: "Subscribes for <Country> updates. Portugal is Default."},
			Handler: onSubscribe,
		},
		{Command: tb.Command{
			Text:        "/unsubscribe",
			Description: "Unsubscribes <Country> for updates. Portugal is Default."},
			Handler: onUnsubscribe,
		},
		{Command: tb.Command{
			Text:        "/subscriptions",
			Description: "Shows current subscriptions."},
			Handler: onSubscriptions,
		},
	}

	var unextendedCommands []tb.Command
	for _, command := range commands {
		b.Handle(command.Command.Text, command.Handler)
		unextendedCommands = append(unextendedCommands, command.Command)
	}
	b.SetCommands(unextendedCommands)

	return b, nil
}
