include .env

LOCAL_BIN:=$(CURDIR)/bin

.PHONY: run build

run: build
	./bin/bot

build:
	go build -o bin/bot cmd/bot/main.go

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml

test:
	go test -v ./...

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.23.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.3

generate:
	go generate ./...

migration-status:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres $(POSTGRES_DSN) status -v

migration-up:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres $(POSTGRES_DSN) up -v

migration-down:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres $(POSTGRES_DSN) down -v

docker-up:
	mkdir -p -m 777 logs
	docker compose build --progress plain &> logs/build.log
	docker compose up -d