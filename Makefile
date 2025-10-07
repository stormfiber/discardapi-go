.PHONY: test build examples install clean fmt vet lint

# Run tests
test:
	go test -v -cover ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run short tests only
test-short:
	go test -v -short ./...

# Build examples
build:
	go build -o bin/example examples/main.go

# Run examples
examples:
	go run examples/main.go

# Install dependencies
install:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Run all checks
check: fmt vet test

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/example-linux-amd64 examples/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/example-darwin-amd64 examples/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/example-windows-amd64.exe examples/main.go

help:
	@echo "Available targets:"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  test-short    - Run short tests only"
	@echo "  build         - Build examples"
	@echo "  examples      - Run examples"
	@echo "  install       - Install dependencies"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run golangci-lint"
	@echo "  clean         - Clean build artifacts"
	@echo "  check         - Run all checks (fmt, vet, test)"
	@echo "  build-all     - Build for multiple platforms"
