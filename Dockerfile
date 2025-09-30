# Step 1: Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install necessary tools
RUN apk add --no-cache git build-base

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Copy the source code
COPY . .

# Debug: List files in /app
RUN ls -la /app

# Build the executable
RUN go build -o main ./cmd/main.go

# Step 2: Create minimal image
FROM alpine:latest

WORKDIR /app

# Copy the built executable
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Command to run
CMD ["./main"]
