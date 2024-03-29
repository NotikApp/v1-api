#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

# sleep until db is initialized
until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
    >&2 echo "Postgres is unavailable - sleeping"
    sleep 1
done

>&2 echo "Postgres is up - executing command"
docker run -v /build/schema:/migrations --network host migrate/migrate -path=/build/schema -database $psql_url up 2

chmod +x go-notik
# run go service
exec $cmd
