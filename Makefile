.PHONY: build test clean lint
.SILENT: build test clean lint

BIN=$(CURDIR)/bin
GO=$(shell which go)
APP=dockercolorize

build:
	$(GO) build -o $(BIN)/$(APP) ./cmd/cli

test:
	go test ./...

clean:
	rm -f $(BIN)/$(APP)*

lint:
	docker run --rm -w /opt -v $(shell pwd):/opt golangci/golangci-lint golangci-lint run
