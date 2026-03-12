#!/bin/sh
set -e

echo "Running migrations..."
goose -dir /app/migrations postgres "$DB_URL" up

echo "Starting server..."
exec ./server
