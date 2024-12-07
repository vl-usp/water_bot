FROM golang:1.23.3-alpine AS builder

WORKDIR /workdir

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN if [ -f ".env.docker" ]; then cp .env.docker .env; fi

RUN go build -o ./bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /workdir

COPY --from=builder /workdir/bin/bot .
COPY --from=builder /workdir/.env .

CMD ["./bot"]
