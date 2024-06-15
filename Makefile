SHELL:=/bin/bash

# Prevent pushing publicly by accident
DOCKER_REGISTRY:=no-registry.local
DOCKER_IMAGE:=htmx-go-portfolio
EXPOSED_AT:=8080

# TailwindCSS CLI installation
TAILWINDCSS_PLATFORM:=linux
TAILWINDCSS_ARCH:=x64
TAILWINDCSS_VERSION:=v3.3.5
TAILWINDCSS:=./tools/tailwindcss

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

############
# External #
############

tailwind-cli:
	@command -v $(TAILWINDCSS) \
		|| echo "$(TAILWINDCSS) missing. Installing..." \
		&& curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/$(TAILWINDCSS_VERSION)/tailwindcss-$(TAILWINDCSS_PLATFORM)-$(TAILWINDCSS_ARCH) \
		&& chmod +x tailwindcss-$(TAILWINDCSS_PLATFORM)-$(TAILWINDCSS_ARCH) \
		&& mv tailwindcss-$(TAILWINDCSS_PLATFORM)-$(TAILWINDCSS_ARCH) $(TAILWINDCSS) \
		&& echo "Installed, showing help..." \
		&& $(TAILWINDCSS) --help

tailwind-build:
	@echo "Building CSS..."
	$(TAILWINDCSS) -o ./web/css/styles.css

##########
# Docker #
##########

docker: docker-setup docker-build docker-run

docker-setup: tailwind-cli

docker-build:
	docker build --tag $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest .

docker-run:
	docker run -v ./configs:/configs -p $(EXPOSED_AT):3000 $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)

##########
# Podman #
##########

podman-docker-sync:
	docker save $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest | podman load

podman-clean:
	mv $$HOME/tmp/htmx-go-portfolio/ $$HOME/tmp/htmx-go-portfolio.old.d

podman-tarball:
	mkdir -p $$HOME/tmp/htmx-go-portfolio && podman save -o $$HOME/tmp/htmx-go-portfolio/latest.tar --format oci-archive $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest

