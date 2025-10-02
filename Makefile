# StarCraft Concurrency War - Learning Edition
# Makefile for common development tasks

.PHONY: help build run test clean examples check-race fmt vet deps

# Default target
help:
	@echo "StarCraft Concurrency War - Learning Edition"
	@echo ""
	@echo "Available commands:"
	@echo "  make build          - Build the main application"
	@echo "  make run           - Run the main application"
	@echo "  make test          - Run all tests"
	@echo "  make test-race     - Run tests with race detector"
	@echo "  make examples      - Run example scenarios"
	@echo "  make fmt           - Format all Go code"
	@echo "  make vet           - Run go vet on all packages"
	@echo "  make check-race    - Check for race conditions"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make deps          - Download dependencies"
	@echo ""
	@echo "Learning targets:"
	@echo "  make basic-combat     - Run basic combat example"
	@echo "  make resource-demo    - Run resource management demo"
	@echo "  make stress-test      - Run performance stress test"

# Build the main application
build:
	@echo "Building StarCraft Concurrency War..."
	go build -o bin/starcraft-war ./cmd/starcraft-war

# Run the main application
run: build
	@echo "Running StarCraft Concurrency War..."
	./bin/starcraft-war

# Run with race detector
run-race:
	@echo "Running with race detector..."
	go run -race ./cmd/starcraft-war

# Run all tests
test:
	@echo "Running tests..."
	go test ./...

# Run tests with race detector
test-race:
	@echo "Running tests with race detector..."
	go test -race ./...

# Run example scenarios
examples:
	@echo "Running examples..."
	@echo "Note: Examples require implementation of TODO functions"
	@echo ""
	-go run ./examples/basic_combat.go
	@echo ""
	-go run ./examples/resource_management.go

# Run specific examples
basic-combat:
	@echo "Running Basic Combat Example..."
	go run -race ./examples/basic_combat.go

resource-demo:
	@echo "Running Resource Management Demo..."
	go run -race ./examples/resource_management.go

# Code quality checks
fmt:
	@echo "Formatting code..."
	go fmt ./...

vet:
	@echo "Running go vet..."
	go vet ./...

# Check for race conditions in examples
check-race:
	@echo "Checking for race conditions..."
	@echo "Building with race detector..."
	go build -race -o bin/starcraft-war-race ./cmd/starcraft-war
	@echo "Run './bin/starcraft-war-race' to test for races"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Development helpers
dev-setup: deps fmt vet
	@echo "Development environment ready!"
	@echo "Start implementing the TODO functions in:"
	@echo "  - internal/types/types.go"
	@echo "  - pkg/utils/utils.go"
	@echo "  - examples/basic_combat.go"

# Check project structure
check-structure:
	@echo "Project structure:"
	@tree -I 'bin|*.git' || find . -type d | grep -E '^[^.]' | sort

# Show implementation progress (count TODO items)
progress:
	@echo "Implementation Progress:"
	@echo "TODO items remaining:"
	@grep -r "TODO:" --include="*.go" . | wc -l
	@echo ""
	@echo "TODO breakdown by package:"
	@grep -r "TODO:" --include="*.go" . | cut -d: -f1 | sort | uniq -c | sort -nr

# Performance testing
stress-test:
	@echo "Running stress test..."
	@echo "Note: Requires implementation of stress test scenario"
	go run -race ./cmd/starcraft-war -scenario=stress-test

# Memory profiling
profile-mem:
	@echo "Running with memory profiling..."
	go run ./cmd/starcraft-war -cpuprofile=cpu.prof -memprofile=mem.prof
	@echo "Analyze with: go tool pprof mem.prof"

# CPU profiling
profile-cpu:
	@echo "Running with CPU profiling..."
	go run ./cmd/starcraft-war -cpuprofile=cpu.prof
	@echo "Analyze with: go tool pprof cpu.prof"

# Learning path helper
learning-path:
	@echo "Suggested Learning Path:"
	@echo ""
	@echo "Phase 1 - Core Concepts:"
	@echo "  1. Implement String() methods in internal/types/types.go"
	@echo "  2. Implement basic utilities in pkg/utils/utils.go"
	@echo "  3. Implement NewUnit and basic getters in internal/types/types.go"
	@echo ""
	@echo "Phase 2 - Unit Management:"
	@echo "  4. Implement unit manager in internal/units/manager.go"
	@echo "  5. Implement AI system in internal/units/ai.go"
	@echo ""
	@echo "Phase 3 - Resource Management:"
	@echo "  6. Implement resource manager in internal/resources/manager.go"
	@echo ""
	@echo "Phase 4 - Battle Simulation:"
	@echo "  7. Implement battle simulator in internal/battle/simulator.go"
	@echo ""
	@echo "Phase 5 - Coordination:"
	@echo "  8. Implement commander system in internal/coordination/commander.go"
	@echo ""
	@echo "Phase 6 - Integration:"
	@echo "  9. Implement main application in cmd/starcraft-war/main.go"
	@echo "  10. Complete examples in examples/"

# Validate implementation
validate:
	@echo "Validating implementation..."
	@echo "Checking if basic functions compile..."
	@go build ./... && echo "✅ All packages compile" || echo "❌ Compilation errors"
	@echo ""
	@echo "Checking for race conditions..."
	@go build -race ./... && echo "✅ No race conditions detected" || echo "❌ Race conditions found"
	@echo ""
	@echo "Running basic validation..."
	@go test ./... && echo "✅ All tests pass" || echo "❌ Test failures"

# Quick start helper
quickstart:
	@echo "🚀 Quick Start Guide"
	@echo ""
	@echo "1. First, implement basic String() methods:"
	@echo "   code internal/types/types.go"
	@echo ""
	@echo "2. Try building the project:"
	@echo "   make build"
	@echo ""
	@echo "3. Run an example (will show what needs implementation):"
	@echo "   make basic-combat"
	@echo ""
	@echo "4. Check your progress:"
	@echo "   make progress"
	@echo ""
	@echo "5. Follow the learning path:"
	@echo "   make learning-path"

# Show what needs to be implemented
todo:
	@echo "TODO items by file:"
	@echo ""
	@for file in $$(find . -name "*.go" -exec grep -l "TODO:" {} \;); do \
		echo "📁 $$file:"; \
		grep -n "TODO:" "$$file" | sed 's/^/  /'; \
		echo ""; \
	done