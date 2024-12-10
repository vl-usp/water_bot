FROM golang:1.23.4-alpine AS builder

WORKDIR /workdir

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN if [ -f .env ]; then echo ".env file exists, using .env"; \
    else cp .env.default .env && echo ".env file not found, using .env.default"; fi

RUN go build -o ./bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /workdir

COPY --from=builder /workdir/bin/bot .
COPY --from=builder /workdir/.env .

CMD ["./bot"]
