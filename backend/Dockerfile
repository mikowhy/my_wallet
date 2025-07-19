# Use the golang-alpine image which is suitable for development
FROM golang:1.24-alpine

# Install git (required by 'go install') and air for hot-reloading
RUN apk add --no-cache git

# Install air, a tool for live-reloading Go applications
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies.
# This step is cached by Docker if the files don't change.
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the application source code.
# The volume mount in docker-compose.yml will keep this in sync.
COPY . .

# Expose the port the application runs on
EXPOSE 8080

# The command to run the application using air for hot-reloading
CMD ["air"]