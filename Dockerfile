# Dockerfile for Go backend
# Start from the official Golang image
FROM golang:1.22-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY api/go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY api/. .

# Build the Go application
RUN go build -o main .

# Start a new minimal image for running the Go app
FROM alpine:latest

# Set environment variables (Optional)
ENV PORT=8080

# Copy the built binary from the builder stage
COPY --from=build /app/main /app/main

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["/app/main"]

