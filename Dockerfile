#################
# Build project #
#################

FROM golang:1.22 as builder

# We don't care about CGO and are in a Linux environment
ENV CGO_ENABLED=0
ENV GOOS=linux

# Enable caching
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

# Setup workspace
WORKDIR /workspace
COPY . .

# Dependencies with cache
RUN --mount=type=cache,target=/gomod-cache \
    make setup

# Build with cache
RUN --mount=type=cache,target=/gomod-cache \
    --mount=type=cache,target=/go-cache \
    make build-all

#############
# Packaging #
#############

FROM scratch as app

USER nobody

# Setup Linux user
COPY --from=builder /etc/passwd /etc/passwd

# Get build artifact
COPY --from=builder /workspace/build/htmx-go-portfolio /htmx-go-portfolio
# Get web templates
COPY --from=builder /workspace/web/public /web/public

# Run GIN in release mode
ENV GIN_MODE=release

EXPOSE 3000
ENTRYPOINT ["/htmx-go-portfolio"]
