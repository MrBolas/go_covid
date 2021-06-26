package dbmodels

import (
	"go_covid/src/config"
	apicontrollers "go_covid/src/api/apicontrollers"
	"log"
	"time"

	"gopkg.in/tucnak/telebot.v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Username       string
	TelegramId     int
	Country        string
	LastReportDate datatypes.Date
}

func (subscription *Subscription) Notify() bool {
	db := config.DB
	bot := config.Bot
	today := datatypes.Date(time.Now())
	data := apicontrollers.GetCountryData(subscription.Country)
	if data.IsToday() {
		log.Printf("Will notify %s of %s\n", subscription.Username, subscription.Country)
		chat := telebot.Chat{ID: int64(subscription.TelegramId)}
		bot.Send(&chat, data.GetReport())
		subscription.LastReportDate = today
		db.Save(&subscription)
		return true
	}
	return false
}

func (subscription Subscription) GetToUptade(subscriptions *[]Subscription) {
	today := datatypes.Date(time.Now())
	db := config.DB
	db.Where("last_report_date < ? OR last_report_date is ?", today, nil).Find(&subscriptions)
}
