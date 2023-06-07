include .env

POSTGRES_PASSWORD = ${psql_password}
POSTGRES_URL = ${psql_url}

down:
	migrate -path ./schema -database '$(POSTGRES_URL)' down

up:
	migrate -path ./schema -database '$(POSTGRES_URL)' up

build:
	go build cmd/main.go
	./main

run:
	go run cmd/main.go

psql-init:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres

build-docker:
	docker build -t go-notik .

compose:
	docker-compose up --build go-notik
	
migrate-init:
	docker run -v /schema:/schema --network host migrate/migrate -path=/schema/ -database "$(POSTGRES_URL)" up