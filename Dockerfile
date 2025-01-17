# Build stage
FROM --platform=$BUILDPLATFORM golang:1.22 AS build
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application for the target platform
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o config-store ./cmd

# Production stage
FROM --platform=$BUILDPLATFORM alpine:latest AS production
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/config-store .

# Ensure execution permission
RUN chmod +x /app/config-store

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["./config-store"]
