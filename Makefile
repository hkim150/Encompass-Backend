db:
	docker compose up db

migrateup:
	migrate -path migration -database "postgres://encompass:password@localhost:5432/encompass?sslmode=disable" up

migratedown:
	migrate -path migration -database "postgres://encompass:password@localhost:5432/encompass?sslmode=disable" down

.PHONY: db migrateup migratedown