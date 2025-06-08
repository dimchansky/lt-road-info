.PHONY: build run test clean install lint fmt

# Variables
BINARY_NAME=lt-road-info
MAIN_PATH=./cmd/lt-road-info
OUTPUT_DIR=output

# Build the binary
build:
	go build -v -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application
run: build
	./$(BINARY_NAME)

# Run with verbose output
run-verbose: build
	./$(BINARY_NAME) -verbose

# Download only restrictions
run-restrictions: build
	./$(BINARY_NAME) -type restrictions

# Download only speed control
run-speed: build
	./$(BINARY_NAME) -type speed-control

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -rf $(OUTPUT_DIR)

# Install the binary to $GOPATH/bin
install:
	go install $(MAIN_PATH)

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

# Verify coordinate transformations are correct
verify-coords:
	@echo "🔍 Verifying coordinate transformations..."
	go run ./cmd/verify-coords

# Run comprehensive tests including coordinate validation
test-all: test
	@echo "🧪 Running coordinate transformation tests..."
	go test ./internal/transform -v
	@echo "🧪 Running integration tests..."
	go test ./internal/eismoinfo -v -run TestCoordinateTransformationRegression
	go test ./internal/arcgis -v -run TestArcGISCoordinateValidation