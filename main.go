package main

import (
	bot "go_covid/controllers"
	"log"
)

func main() {

	b, err := bot.TeleCovidBot()

	if err != nil {
		log.Fatal(err)
	}

	b.Start()

}
