# Simple Makefile for a Go project

# Define the name of the output binary
BINARY_NAME=simple-bank

# Path to the main.go file
MAIN_FILE=/Users/vivmehta1/Documents/Projects/GoLang/simpleBank/cmd/api/main.go

# Build the Go application
build:
	export APP_ENV=debug
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the built application
run: build
	./$(BINARY_NAME)

# Added this text command because the task.json does not indetifies run
test: build
	./$(BINARY_NAME)

# Clean up the build artifacts
clean:
	rm $(BINARY_NAME)

# Optional: Add more targets here for testing, linting, etc.