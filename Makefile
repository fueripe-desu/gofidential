.DEFAULT_GOAL := help

# Build the Go application binary into the bin/ directory.
build:
	@mkdir -p bin
	@go build -o bin/myapp .

# Clear the contents of the bin/ directory without removing the folder.
clean:
	@rm -rf bin/*

# Run all tests.
test:
	@go test -v ./...

# Format code using gofumpt.
fmt:
	@gofumpt -l -w .

# Install gofumpt for go code formatting.
install-fmt:
	@go install mvdan.cc/gofumpt@latest

# Shows a help message explaining all available targets.
help:
	@echo "Available targets:"
	@echo "  test		Run Go tests"
	@echo "  fmt		Format the code with gofumpt"
	@echo "  install-fmt	Install gofumpt if not installed"
