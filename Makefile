# Simple Makefile for a Go project

# Define the name of the output binary
BINARY_NAME=simple_bank

# Path to the main.go file
MAIN_FILE=./cmd/api/main.go

# Build the Go application
build:
	export APP_ENV=debug
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the built application
run: build
	./$(BINARY_NAME)

# Clean up the build artifacts
clean:
	rm $(BINARY_NAME)

# Optional: Add more targets here for testing, linting, etc.