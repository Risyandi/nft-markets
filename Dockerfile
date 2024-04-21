# Start from the official Go image
FROM golang:1.17-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o /go/bin/app

# Start a new stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=builder /go/bin/app /app/app

# Expose the port the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["/app/app"]
