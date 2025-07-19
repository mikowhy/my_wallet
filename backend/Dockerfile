# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Update packages in the builder image to mitigate vulnerabilities
RUN apk update && apk upgrade

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies first
# This leverages Docker's layer caching
COPY go.* ./
RUN go mod download

# Copy the rest of the source code
COPY src/ ./src/

# Build the application for a Linux environment.
# CGO_ENABLED=0 creates a static binary. The -a flag forces rebuilding of packages that are already up-to-date.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./src/main.go

# Stage 2: Create the final, minimal image
FROM gcr.io/distroless/static-debian12

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the application runs on.
EXPOSE 8080

# The command to run the application
CMD ["/app/server"]