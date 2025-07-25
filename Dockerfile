# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o bookservice

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bookservice .
EXPOSE 8080
ENV GOOGLE_BOOKS_API_KEY=""
CMD ["./bookservice"] 