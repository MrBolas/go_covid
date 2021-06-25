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
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, will continue")
	}

	result := os.Getenv(key)
	if len(result) < 1 {
		log.Printf("key %s not found in env\n", key)
	}
	return result
}

func main() {
	token := goDotEnvVariable("TELEGRAM_TOKEN")
	db := bot.InitDB()

	b, err := bot.TeleCovidBot(token, db)

	if err != nil {
		log.Fatal(err)
	}

	b.Start()

}
