#!/bin/sh
set -e

echo "DB_SOURCE is set to: $DB_SOURCE"
echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "start the app"
exec "$@"