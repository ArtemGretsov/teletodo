package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/pressly/goose/v3"

	"github.com/ArtemGretsov/teletodo/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "up":
		if err = goose.Up(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err = goose.Down(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	case "status":
		if err = goose.Status(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown command")
	}
}
