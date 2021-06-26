package config

import (
	"gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Bot *telebot.Bot

