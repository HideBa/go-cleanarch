
run-app:
	go run ./app/server.go

gen/migrate:
	go migrate create -ext sql -dir ./app/infrastructure/migrations

.PHONY: run-app gen/migrate