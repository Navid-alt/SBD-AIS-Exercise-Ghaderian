# todo
# Stage 1: The Builder 

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN /bin/sh scripts/build-application.sh

# Stage 2: The Final Runner
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ordersystem .

EXPOSE 8080

ENTRYPOINT [ "./ordersystem" ]