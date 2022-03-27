migrate-up:
	go run migrations/migrate.go up
migrate-down:
	go run migrations/migrate.go down
migrate-status:
	go run migrations/migrate.go status
