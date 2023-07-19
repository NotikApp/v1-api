# BUILD STAGE
FROM golang:alpine AS builder

# install migrate cli so we can use it in prod stage
RUN GOBIN=/usr/local/bin/ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# PROD STAGE
# use alpine to reduce image`s size
FROM alpine

WORKDIR /build

# copy built exe file
COPY /go-notik /build/go-notik

# copy wait-for-postgres.sh
COPY wait-for-postgres.sh /build/wait-for-postgres.sh

# copy static files of website
COPY dist /build/dist

# copy .env file, psql_url will be overrided in compose
COPY .env /build/.env

# copy email html files
COPY static /build/static

# copy migrations dir from build
COPY schema /build/schema

# copy golang-migrate
COPY --from=builder /usr/local/bin/ /usr/local/bin/

# install postgresql-client
RUN apk update
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# run service
CMD ["./go-notik"]