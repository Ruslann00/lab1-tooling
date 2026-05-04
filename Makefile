APP_NAME=lab1-tooling

.PHONY: fmt lint test build all clean

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./...

build:
	go build -o bin/$(APP_NAME) ./cmd/app

all: fmt lint test build

clean:
	rm -rf bin