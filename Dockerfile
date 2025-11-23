FROM golang:1.25-alpine AS builder

WORKDIR /service

COPY . .

RUN go build -o ./bin/server ./cmd/server

FROM alpine:3.22.2

COPY --from=builder /service/bin/server /app/server

EXPOSE 8080