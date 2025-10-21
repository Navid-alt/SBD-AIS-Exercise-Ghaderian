#!/bin/sh

SERVICE_CONTAINER="orderservice"
IMAGE_NAME="orderservice-image"
HOST_PORT_SERVICE=3000
NETWORK_NAME="order-network"
ENV_FILE="debug.env"

docker build -t "$IMAGE_NAME" .

docker run -d \
  --name "$SERVICE_CONTAINER" \
  --network "$NETWORK_NAME" \
  --env-file "$ENV_FILE" \
  -p "$HOST_PORT_SERVICE:$CONTAINER_PORT_SERVICE" \
  "$IMAGE_NAME"
