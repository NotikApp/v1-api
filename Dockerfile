FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o go-notik cmd/main.go

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

# install postgresql-client
RUN apk update
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

CMD [". /go-notik"]