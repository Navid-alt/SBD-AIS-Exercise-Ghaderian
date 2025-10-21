#!/bin/sh

DB_CONTAINER="db"
NETWORK_NAME="order-network"
ENV_FILE="debug.env"
POSTGRES_IMAGE="postgres:18"
VOLUME_NAME="order-data"
HOST_PORT_DB=5432

docker network create "$NETWORK_NAME"

docker run -d \
  --name "$DB_CONTAINER" \
  --network "$NETWORK_NAME" \
  --env-file "$ENV_FILE" \
  -v "$VOLUME_NAME:/var/lib/postgresql/18/docker" \
  -p "$HOST_PORT_DB:5432" \
  "$POSTGRES_IMAGE"