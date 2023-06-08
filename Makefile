include .env

POSTGRES_PASSWORD = ${psql_password}
POSTGRES_URL = ${psql_url}

# drop migrations
down:
	migrate -path ./schema -database '$(POSTGRES_URL)' down

# aplly migrations to db
up:
	migrate -path ./schema -database '$(POSTGRES_URL)' up

# build golang service
build:
	go build cmd/main.go
	./main

# run golang service
run:
	go run cmd/main.go

# use for development. inits postgres
psql-init:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres
	make up

build-docker:
	docker build -t go-notik .

# init docker container with compose
compose:
	docker-compose up --build go-notik