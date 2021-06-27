package config

import (
	"gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var Bot *telebot.Bot
var Loc, _ = time.LoadLocation("Europe/Lisbon")

