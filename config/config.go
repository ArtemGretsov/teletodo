package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Database struct {
	DSN string `yaml:"dsn" env:"DATABASE_DSN"`
}

type TelegramBot struct {
	Token     string           `yaml:"token" env:"TELEGRAM_BOT_TOKEN"`
	WhiteList StringEnumerable `yaml:"white_list" env:"TELEGRAM_WHITE_LIST"`
}

type Config struct {
	Database    Database    `yaml:"database"`
	TelegramBot TelegramBot `yaml:"telegram_bot"`
}

func NewConfig() (*Config, error) {
	localPath := "config/local.yml"
	cfg := &Config{}

	if _, err := os.Stat(localPath); err == nil {
		err = cleanenv.ReadConfig(localPath, cfg)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	return cfg, nil
}
