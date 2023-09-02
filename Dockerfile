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

# install postgresql-client
RUN apk update
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# run service
CMD ["./go-notik"]
