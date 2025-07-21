# syntax=docker/dockerfile:1

# 1. Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Для поддержки env, если нужно: устанавливаем git (для go mod)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o imageflow cmd/app/main.go

# 2. Run stage
FROM alpine:latest

WORKDIR /app

# Копируем бинарник из build stage
COPY --from=builder /app/imageflow .

# (опционально) Копируем .env, если нужен
# COPY .env .

EXPOSE 8080

CMD ["./imageflow"]
