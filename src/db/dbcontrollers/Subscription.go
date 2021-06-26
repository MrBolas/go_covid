package dbcontrollers

import (
	"go_covid/src/config"
	models "go_covid/src/db/dbmodels"
	"log"
	"time"

	"gorm.io/datatypes"
)

func handleSubscriptions() {
	log.Println("handling subscriptions")
	var subscriptions []models.Subscription
	today := datatypes.Date(time.Now())
	db := config.DB
	db.Where("last_report_date < ? OR last_report_date is ?", today, nil).Find(&subscriptions)
	for _, subscription := range subscriptions {
		subscription.Notify()
	}
}

func SetupSubscriptionsFetching(minutes int) (*time.Ticker, error) {
	tickerDuration := time.Duration(minutes) * time.Minute
	ticker := time.NewTicker(tickerDuration)

	handleSubscriptions()

	go func() {
		for range ticker.C {
			handleSubscriptions()
		}
	}()
	return ticker, nil
}
