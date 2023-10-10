db:
	docker compose up db

migrateup:
	migrate -path db/migration -database "postgres://encompass:password@localhost:5432/encompass?sslmode=disable" up

migratedown:
	migrate -path db/migration -database "postgres://encompass:password@localhost:5432/encompass?sslmode=disable" down

test:
	go test -v -cover ./...

.PHONY: db migrateup migratedown test