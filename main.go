package main

import (
	"go_covid/src/bot"
	config "go_covid/src/config"
	"go_covid/src/db/dbcontrollers"
	"go_covid/src/db/dbmodels"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func InitDB(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&dbmodels.Subscription{})
	return db
}

func main() {
	dbName := goDotEnvVariable("DB_NAME")
	if len(dbName) < 1 {
		dbName = "test.db"
	}
	db := InitDB(dbName)
	config.DB = db
	token := goDotEnvVariable("TELEGRAM_TOKEN")
	b, err := bot.TeleCovidBot(token)
	config.Bot = b

	ticker, _ := dbcontrollers.SetupSubscriptionsFetching(10)
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
	ticker.Stop()
}
