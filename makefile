.PHONY: setup run traffic

start: setup run

run:
	go run cmd/webapi/main.go

setup:
	docker compose down
	docker compose up -d
	sleep 2
	go run cmd/migrator/main.go
	go run cmd/seeder/main.go

spawn:
	go run cmd/spawner/main.go
