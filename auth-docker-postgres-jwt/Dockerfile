#GPT
# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the container and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the Go application from the cmd directory
RUN go build -o /go/bin/myapp ./cmd/main.go

# Stage 2: Create a smaller image for running the Go application
FROM alpine:3.18

# Copy the built Go binary from the builder stage
COPY --from=builder /go/bin/myapp /usr/local/bin/myapp

# Set the entrypoint to run the application
ENTRYPOINT ["/usr/local/bin/myapp"]

# Expose the port the application is running on (if necessary, e.g., port 8080)
EXPOSE 8080