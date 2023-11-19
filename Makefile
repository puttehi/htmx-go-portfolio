SHELL:=/bin/bash

# Prevent pushing publicly by accident
DOCKER_REGISTRY:=no-registry
DOCKER_IMAGE:=htmx-go-portfolio
EXPOSED_AT:=8080

.PHONY:$(MAKECMDGOALS)

#########
# Local #
#########

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

setup:
	go mod download -x

dev: setup
	air

##########
# Docker #
##########

docker: docker-build docker-run

docker-build:
	docker build --tag $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest .

docker-run:
	docker run -p $(EXPOSED_AT):3000 $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)

