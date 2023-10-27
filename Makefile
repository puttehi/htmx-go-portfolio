SHELL:=/bin/bash

.PHONY:$(MAKECMDGOALS)

all: test build-all
	./build/htmx-go-portfolio

run:
	go run cmd/htmx-go-portfolio/main.go

build: test
	go build -o build/htmx-go-portfolio cmd/htmx-go-portfolio/main.go

build-all: test
	go build -o build/ ./...

clean:
	rm -r build/*

test:
	go test -v ./...
