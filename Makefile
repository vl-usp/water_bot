run: build
	./bin/water_bot

build:
	go build -o bin/water_bot cmd/main.go