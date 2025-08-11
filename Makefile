.PHONY: build test lint

build:
	go build -o bin/hexlet-path-size .
	chmod +x bin/hexlet-path-size

test:
	go test -v ./...

lint:
	golangci-lint run

run:
	./bin/hexlet-path-size