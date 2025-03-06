# Use the official Go image based on Alpine Linux
FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o money-transfer ./cmd/server

# Expose the application port (ensure it matches your config)
EXPOSE 8080

# Run the application
CMD ["./money-transfer"]
