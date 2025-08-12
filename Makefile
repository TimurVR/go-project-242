.PHONY: build test lint run

build:
	go build -o bin\hexlet-path-size   

test:
	go test -v ./testdata

lint:
	golangci-lint run

run: build
	bin\hexlet-path-size  