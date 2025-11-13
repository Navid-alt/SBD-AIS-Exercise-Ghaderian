#!/bin/sh
# Exit if any command fails
set -e
cd /app
go mod tidy
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem