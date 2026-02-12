.PHONY: help test build run wire build-arm64 proto

default: help

APP_NAME = ddd

VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
GIT_HASH := $(shell git rev-parse HEAD)
# short hash
# GIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -X 'ddd/internal/version.Version=$(VERSION)' \
           -X 'ddd/internal/version.GitHash=$(GIT_HASH)' \
           -X 'ddd/internal/version.BuildTime=$(BUILD_TIME)'

help:
	@echo "Usage:"
	@echo "    make run"
	@echo "    make build"
	@echo "    make test"

wire:
	wire ./cmd/server/cmd/

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) ./cmd/server

build-arm64:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -ldflags "$(LDFLAGS)" -o bin/arm64/$(APP_NAME) ./cmd/server

run: wire build
	./bin/$(APP_NAME)

test:
	go test -v -count=1 ./...

proto:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		api/proto/user/v1/auth.proto
