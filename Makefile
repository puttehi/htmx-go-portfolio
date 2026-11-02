ifeq ($(OS),Windows_NT)
SHELL:=bash.exe
EXE_EXT:=.exe
else
SHELL:=/bin/bash
EXE_EXT:=
endif

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
HOST_ARCH:=$(shell go env GOHOSTARCH)
TAILWINDCSS_PLATFORM:=$(if $(filter Darwin,$(shell uname -s)),macos,linux)
ifeq ($(OS),Windows_NT)
TAILWINDCSS_PLATFORM:=windows
endif
TAILWINDCSS_ARCH:=$(if $(filter amd64,$(HOST_ARCH)),x64,$(if $(filter arm,$(HOST_ARCH)),armv7,$(HOST_ARCH)))
TAILWINDCSS_VERSION:=v3.3.5
TAILWINDCSS:=$(TOOLS_ROOT_DIR)/tailwindcss$(EXE_EXT)
TAILWINDCSS_IN:=$(WEB_ROOT_DIR)/assets/css/styles.css
TAILWINDCSS_OUT:=$(WEB_ROOT_DIR)/assets/css/compiled.css

# Hugo CLI installation
HUGO:=hugo
ifeq ($(OS),Windows_NT)
HUGO:=$(TOOLS_ROOT_DIR)/hugo.exe
endif
HUGO_VERSION:=0.127.0

# Air CLI installation
AIR_VERSION:=v1.61.7
AIR:=$(TOOLS_ROOT_DIR)/air$(EXE_EXT)

.PHONY:$(MAKECMDGOALS) setup run-built air-cli

#########
# Local #
#########

all: test build-all
	@echo "Build complete. Run 'make run-built' to start the compiled build."

run-built:
	./build/htmx-go-portfolio$(EXE_EXT)

run:
	go run cmd/htmx-go-portfolio/main.go

build: test
	go build -o build/htmx-go-portfolio$(EXE_EXT) cmd/htmx-go-portfolio/main.go

build-all: tailwind-build hugo-build
	go build -o build/ ./...

clean:
	rm -r build || echo "Already gone?"
	rm -r tmp || echo "Already gone?"
	rm -r $(WEB_ROOT_DIR)/public || echo "Already gone?"

test:
	go test -v ./...

setup-go:
	go mod download -x

setup: setup-go setup-tools air-cli

dev: setup
ifeq ($(OS),Windows_NT)
	$(AIR) --build.cmd "go build -o ./tmp/main.exe ./cmd/htmx-go-portfolio" --build.bin '$(subst /,\,$(CURDIR)/tmp/main.exe)' --build.exclude_dir 'web\public,build'
else
	$(AIR)
endif

############
# External #
############

setup-tools: tailwind-cli hugo-cli

ensure-tools:
	mkdir -p $(TOOLS_ROOT_DIR)

air-cli: ensure-tools
	@(command -v $(AIR) && echo "Air found!") \
		|| (echo "$(AIR) missing. Installing..." \
		&& GOBIN="$(abspath $(TOOLS_ROOT_DIR))" go install github.com/air-verse/air@$(AIR_VERSION) \
		&& $(AIR) -v)

tailwind-cli: ensure-tools
	@(command -v $(TAILWINDCSS) && echo "Tailwind found!") \
		|| (echo "$(TAILWINDCSS) missing. Installing..." \
		&& curl -fL https://github.com/tailwindlabs/tailwindcss/releases/download/$(TAILWINDCSS_VERSION)/tailwindcss-$(TAILWINDCSS_PLATFORM)-$(TAILWINDCSS_ARCH)$(EXE_EXT) -o $(TAILWINDCSS) \
		&& chmod +x $(TAILWINDCSS) \
		&& echo "Installed to $(TAILWINDCSS), showing help..." \
		&& $(TAILWINDCSS) --help)

tailwind-build:
	@echo "Building CSS..."
	$(TAILWINDCSS) -i $(TAILWINDCSS_IN) -o $(TAILWINDCSS_OUT)

hugo-cli: ensure-tools
ifeq ($(OS),Windows_NT)
	@(command -v $(HUGO) && echo "Hugo found!") \
		|| (echo "$(HUGO) missing. Installing..." \
		&& curl -fL https://github.com/gohugoio/hugo/releases/download/v$(HUGO_VERSION)/hugo_extended_$(HUGO_VERSION)_windows-$(HOST_ARCH).zip -o $(TOOLS_ROOT_DIR)/hugo.zip \
		&& unzip -o $(TOOLS_ROOT_DIR)/hugo.zip hugo.exe -d $(TOOLS_ROOT_DIR) \
		&& rm $(TOOLS_ROOT_DIR)/hugo.zip \
		&& $(HUGO) version)
else
	@(command -v hugo && echo "Hugo found!") \
		|| (echo "hugo missing. Installing to system Go path..." \
		&& CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@v$(HUGO_VERSION) \
		&& echo "Installed to $$(which hugo), showing version..." \
		&& hugo version)
endif

hugo-build:
	$(HUGO) --gc -s $(WEB_ROOT_DIR)

# Enables draft pages and changes webroot to local gin port
hugo-build-dev:
	$(HUGO) --gc -s $(WEB_ROOT_DIR) -D -b http://localhost:$(EXPOSED_AT)

# Enables draft page and changes webroot to local air proxy port (for make dev)
hugo-build-dev-air-proxy:
	$(HUGO) --gc -s $(WEB_ROOT_DIR) -D -b http://localhost:$(EXPOSED_AT_AIR_PROXY)

# Shows drafts, uses localhost as baseUrl
hugo-run-dev:
	$(HUGO) server -D -s $(WEB_ROOT_DIR)

# Production build, uses localhost as baseUrl
hugo-run:
	$(HUGO) server -s $(WEB_ROOT_DIR)

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

