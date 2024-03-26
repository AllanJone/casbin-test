# Makefile for Casbin Go application

# Project-specific settings
BINARY_NAME=myapp
SOURCE_FILE=main.go

# Build the application
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) $(SOURCE_FILE)

# Run the application
run: build
	@echo "Running application..."
	./$(BINARY_NAME)

# Clean up the build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Display this help
help:
	@echo "Makefile commands:"
	@echo "  build   - Compile the application."
	@echo "  run     - Build and run the application."
	@echo "  clean   - Remove binary files."
	@echo "  help    - Display this help."

# Prevent make from confusing the command names with file names
.PHONY: build run clean test help
