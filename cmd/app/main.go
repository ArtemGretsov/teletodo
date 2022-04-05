package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ArtemGretsov/teletodo/config"
	"github.com/ArtemGretsov/teletodo/internal/service/gateways/telegram"
	"github.com/ArtemGretsov/teletodo/internal/service/usecases"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config reading error: %+v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm creating error: %+v", err)
	}

	telegramBot, err := telegram.NewTelegram(cfg, usecases.NewUsecases(db))
	if err != nil {
		log.Fatalf("telegram bot creating error: %+v", err)
	}

	if err = telegramBot.Start(); err != nil {
		log.Fatalf("telegram bot init error: %+v", err)
	}
}
