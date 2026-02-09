.PHONY: help test build run wire

default: help

APP_NAME = ddd

help:
	@echo "Usage:"
	@echo "    make run"
	@echo "    make build"
	@echo "    make test"

wire:
	wire ./cmd/server

build:
	go build -o bin/$(APP_NAME) ./cmd/server

run: wire build
	./bin/ddd

test:
	go test -v -count=1 ./...
