.PHONY: build test clean lint build-linux build-darwin
.SILENT: build test clean lint

BIN=$(CURDIR)/bin
GO=$(shell which go)
APP=dockercolorize

build: build-linux build-darwin

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-linux-amd64 ./cmd/cli

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-darwin-amd64 ./cmd/cli

test:
	go test ./...

clean:
	rm -f $(BIN)/$(APP)*

lint:
	docker run --rm -w /opt -v $(shell pwd):/opt golangci/golangci-lint golangci-lint run
