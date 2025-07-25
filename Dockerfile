# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache build-base
COPY . .
RUN go mod tidy && go build -o bookservice

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bookservice .
EXPOSE 8080
CMD ["./bookservice"] 