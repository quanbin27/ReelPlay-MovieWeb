#!/bin/sh
set -e



echo "Running migrations..."
migrate -path migrate/migrations -database "mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" up

echo "Starting application..."
exec "$@"
