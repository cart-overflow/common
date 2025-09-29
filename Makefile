SHELL = /bin/bash

# Set up development environment
setup:
	@lefthook install
	@echo "✅ Development environment ready"

# Format Go code using golangci-lint
fmt:
	@echo "🔧 Formatting Go code..."
	@golangci-lint fmt
	@echo "✅ Code formatting complete"

# Run linter checks using gloangci-lint
lint:
	@echo "🔨 Running linter checks..."
	@golangci-lint run
	@echo "✅ Linting complete"

# Fix linting if possible and format the source code
fix: 
	@echo "🛠️ Fix linter issues and formatting the code..."
	@golangci-lint run --fix
	@echo "✅ Fixing complete"

# CI Build discarding artefacts
check-build:
	@echo ⏳ "Building..."
	@go build -o /dev/null ./...
	@echo "✅ Building complete"

.PHONY: setup fmt lint fix check-build