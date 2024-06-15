SHELL:=/bin/bash

# Prevent pushing publicly by accident
DOCKER_REGISTRY:=no-registry.local
DOCKER_IMAGE:=htmx-go-portfolio
DOCKER_TAG:=latest
DOCKER_FULL:=$(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(DOCKER_TAG)
DOCKER_BUILD_ARGS:=
DOCKER_SAVE_ARGS:=
EXPOSED_AT:=3000

# E.g. -i filepath/of/DOCKER_SAVE_ARGS
PODMAN_LOAD_ARGS:=

TOOLS_ROOT_DIR:=./tools

# TailwindCSS CLI installation
TAILWINDCSS_PLATFORM:=linux
TAILWINDCSS_ARCH:=x64
TAILWINDCSS_VERSION:=v3.3.5
TAILWINDCSS:=$(TOOLS_ROOT_DIR)/tailwindcss

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

ensure-tools:
	mkdir -p $(TOOLS_ROOT_DIR)

tailwind-cli: ensure-tools
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
	docker build --tag $(DOCKER_FULL) $(DOCKER_BUILD_ARGS) .

docker-run:
	docker run -v ./configs:/configs -p $(EXPOSED_AT):3000 $(DOCKER_FULL)

docker-save:
	docker save $(DOCKER_SAVE_ARGS) $(DOCKER_FULL)

##########
# Podman #
##########

podman-docker-sync: docker-save
	podman load $(PODMAN_LOAD_ARGS)

podman-clean:
	mv $$HOME/tmp/htmx-go-portfolio/ $$HOME/tmp/htmx-go-portfolio.old.d

podman-tarball:
	mkdir -p $$HOME/tmp/htmx-go-portfolio && podman save -o $$HOME/tmp/htmx-go-portfolio/$(DOCKER_TAG).tar --format oci-archive $(DOCKER_FULL)

