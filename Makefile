.PHONY: help test build run wire build-arm64

default: help

APP_NAME = ddd

help:
	@echo "Usage:"
	@echo "    make run"
	@echo "    make build"
	@echo "    make test"

wire:
	wire ./cmd/server/cmd/

build:
	go build -o bin/$(APP_NAME) ./cmd/server

build-arm64:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o bin/arm64/$(APP_NAME) ./cmd/server

run: wire build
	./bin/ddd

test:
	go test -v -count=1 ./...
