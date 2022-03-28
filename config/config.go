package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Database struct {
	DSN string `env:"DATABASE_DSN" env-default:"host=localhost user=teletodo password=teletodo dbname=teletodo port=5432 sslmode=disable TimeZone=Europe/Moscow"`
}

type TelegramBot struct {
	Token     string           `env:"TELEGRAM_BOT_TOKEN" env-required:"true"`
	WhiteList StringEnumerable `env:"TELEGRAM_WHITE_LIST"`
}

type Config struct {
	Database    Database
	TelegramBot TelegramBot
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	return cfg, nil
}
