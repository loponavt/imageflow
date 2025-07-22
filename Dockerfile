# syntax=docker/dockerfile:1

# 1. Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o imageflow cmd/app/main.go

# 2. Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/imageflow .

COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./imageflow"]
