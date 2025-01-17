# Build stage
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

# Build the application from the cmd directory
RUN go build -o config-store ./cmd

# Production stage
FROM alpine:latest AS production
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/config-store .

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["./config-store"]