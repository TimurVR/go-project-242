.PHONY: build test lint run

build:
	@if not exist bin mkdir bin
	go build -o bin/hexlet-path-size cmd/hexlet-path-size/main.go

test:
	go test -v testdata/path_size_test.go

lint:
	golangci-lint run

run: build
	bin\hexlet-path-size.exe $(ARGS)