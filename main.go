package main

import (
	bot "go_covid/controllers"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
	token := goDotEnvVariable("TELEGRAM_TOKEN")

	b, err := bot.TeleCovidBot(token)

	if err != nil {
		log.Fatal(err)
	}

	b.Start()

}
