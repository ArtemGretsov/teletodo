package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pressly/goose/v3"

	"github.com/ArtemGretsov/teletodo/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Database{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	case "status":
		if err := goose.Status(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown command")
	}
}
