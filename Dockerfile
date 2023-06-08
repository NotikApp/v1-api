# BUILD STAGE
FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

# build 
RUN go build -o go-notik cmd/main.go

# install migare cli so we can use it in prod stage
RUN GOBIN=/usr/local/bin/ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


# PROD STAGE
# use alpine to reduce image`s size
FROM alpine

WORKDIR /build

# copy built exe file
COPY --from=builder /build/go-notik /build/go-notik

# copy wait-for-postgres.sh
COPY --from=builder /build/wait-for-postgres.sh /build/wait-for-postgres.sh

# copy static files of website
COPY --from=builder /build/dist /build/dist

# copy .env file, psql_url will be overrided in compose
COPY --from=builder /build/.env /build/.env

# copy email html files
COPY --from=builder /build/static /build/static

# copy migrations dir from build
COPY --from=builder /build/schema /build/schema

# copy golang-migrate
COPY --from=builder /usr/local/bin/ /usr/local/bin/

# install postgresql-client
RUN apk update
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# run service
CMD [". /go-notik"]