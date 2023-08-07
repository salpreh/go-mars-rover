.PHONY=build test

build:
	@go build -o build

run:
	@go run ./cmd

test:
	@go test ./tests/...