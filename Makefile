# Novapi Makefile

.PHONY: all build dev test lint clean

# Default target
all: test build

# Build the application
build:
	wails build

# Run in development mode
dev:
	wails dev

# Run backend tests
test:
	go test -v ./...

# Lint the code (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf build/bin
