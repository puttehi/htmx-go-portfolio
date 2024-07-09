SHELL:=/bin/bash

# Prevent pushing publicly by accident
DOCKER_REGISTRY:=no-registry.local
DOCKER_IMAGE:=htmx-go-portfolio
DOCKER_TAG:=latest
DOCKER_FULL:=$(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(DOCKER_TAG)
DOCKER_BUILD_ARGS:=
DOCKER_SAVE_ARGS:=
EXPOSED_AT:=3000
# Live reload port
EXPOSED_AT_AIR_PROXY:=3001

# E.g. -i filepath/of/DOCKER_SAVE_ARGS
PODMAN_LOAD_ARGS:=

TOOLS_ROOT_DIR:=./tools

WEB_ROOT_DIR:=./web

# TailwindCSS CLI installation
TAILWINDCSS_PLATFORM:=linux
TAILWINDCSS_ARCH:=x64
TAILWINDCSS_VERSION:=v3.3.5
TAILWINDCSS:=$(TOOLS_ROOT_DIR)/tailwindcss
TAILWINDCSS_IN:=$(WEB_ROOT_DIR)/assets/css/styles.css
TAILWINDCSS_OUT:=$(WEB_ROOT_DIR)/assets/css/compiled.css

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

build-all: tailwind-build hugo-build
	go build -o build/ ./...

clean:
	rm -r build || echo "Already gone?"
	rm -r tmp || echo "Already gone?"
	rm -r $(WEB_ROOT_DIR)/public || echo "Already gone?"

test:
	go test -v ./...

setup: tailwind-cli hugo-cli
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
		&& echo "Installed to $(TAILWINDCSS), showing help..." \
		&& $(TAILWINDCSS) --help

tailwind-build:
	@echo "Building CSS..."
	$(TAILWINDCSS) -i $(TAILWINDCSS_IN) -o $(TAILWINDCSS_OUT)

hugo-cli:
	@command -v hugo \
		|| echo "hugo missing. Installing to system Go path..." \
		&& CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@v0.127.0 \
		&& echo "Installed to $$(which hugo), showing version..." \
		&& hugo version

hugo-build:
	hugo --gc -s $(WEB_ROOT_DIR)

# Enables draft pages and changes webroot to local gin port
hugo-build-dev:
	hugo --gc -s $(WEB_ROOT_DIR) -D -b http://localhost:$(EXPOSED_AT)

# Enables draft page and changes webroot to local air proxy port (for make dev)
hugo-build-dev-air-proxy:
	hugo --gc -s $(WEB_ROOT_DIR) -D -b http://localhost:$(EXPOSED_AT_AIR_PROXY)

# Shows drafts, uses localhost as baseUrl
hugo-run-dev:
	hugo server -D -s $(WEB_ROOT_DIR)

# Production build, uses localhost as baseUrl
hugo-run:
	hugo server -s $(WEB_ROOT_DIR)

##########
# Docker #
##########

docker: docker-build docker-run

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

