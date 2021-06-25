package models

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Username   string
	TelegramId int
	Country    string
}
