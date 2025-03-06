.PHONY: build test run clean docker

# Build the application
build:
	go build -o money-transfer ./cmd/server

# Run tests (unit and integration)
test:
	go test -v ./...

# Run the application
run: build
	./money-transfer

# Clean build artifacts
clean:
	rm -f money-transfer

# Build Docker image
docker:
	docker build -t money-transfer:latest .
