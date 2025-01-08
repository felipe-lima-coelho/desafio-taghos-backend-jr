#!/usr/bin/env sh
set -e

echo "=> Running migrations..."
migrate -path ./migrations \
  -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
  up

echo "=> Starting Air..."
air -c .air.toml
