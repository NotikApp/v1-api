include .env

POSTGRES_PASSWORD = ${psql_password}
POSTGRES_URL = ${psql_url}

# drop migrations
down:
	migrate -path ./schema -database $(POSTGRES_URL) down

# aplly migrations to db
up:
	migrate -path ./schema -database $(POSTGRES_URL) up

# build golang service
build:
	set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 &&go build -o go-notik cmd/main.go
	./go-notik

# run golang service
run:
	go run cmd/main.go

# use for development. inits postgres
psql-init:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres

build-docker:
	docker build -t go-notik .

# init docker container with compose
compose:
	docker-compose up --build go-notik