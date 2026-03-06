# AGENTS.md - SysWatch Development Guide

This file provides guidance for AI agents working on the SysWatch codebase.

## Project Overview

SysWatch is a CLI system monitoring tool built with Go (v1.24.0). It monitors CPU, memory, disk, and process count in real-time using the Cobra CLI framework and gopsutil library.

## Build Commands

### Full Build (all platforms)
```bash
make build
```

### Individual Platform Builds
```bash
make build-macos      # macOS x86_64
make build-macos-arm64 # macOS ARM64
make build-linux      # Linux x86_64
make build-linux-arm64 # Linux ARM64
make build-windows    # Windows x86_64
```

### Run Locally
```bash
make run              # Build and run locally
```

### Development Dependencies
```bash
make deps             # Install gotestsum
make tidy             # Run go mod tidy
```

## Test Commands

### Run All Tests
```bash
make test             # Runs tests with coverage (requires 90% minimum)
make test-coverage    # Generate HTML coverage report
make test-verbose     # Verbose test output
```

### Run Single Test
```bash
# Run a specific test function
go test -v ./... -run TestFunctionName

# Run tests in a specific package
go test -v ./internal/metrics

# Run tests with coverage for a specific package
go test -v -coverprofile=coverage.out ./internal/metrics -covermode=atomic
```

### Watch Mode
```bash
make test-watch       # Watch for changes and re-run tests
```

## Lint Commands

```bash
make lint             # Run golangci-lint
make lint-fix         # Run golangci-lint with auto-fix
```

## Code Style Guidelines

### Formatting
- Use 2-space indentation (defined in `.editorconfig`)
- Use `gofmt` and `goimports` for formatting
- Always use LF line endings
- Trim trailing whitespace
- Add final newline to files

### Imports
- Use standard Go import grouping:
  1. Standard library
  2. Third-party packages
  3. Internal packages
- Use explicit imports (no implicit package names)
- Run `go mod tidy` after adding dependencies

### Naming Conventions
- **Files**: lowercase, snake_case (e.g., `metrics_test.go`)
- **Packages**: lowercase, short names (e.g., `cmd`, `metrics`)
- **Functions/Variables**: MixedCase for exports, snake_case for unexported
- **Constants**: MixedCase for exports, snake_case for unexported
- **Interfaces**: Add `er` suffix for single-method interfaces (e.g., `MetricsProvider`)
- **Tests**: Use table-driven tests and suite-based testing with testify

### Types
- Use concrete types over interfaces unless polymorphism is needed
- Define interfaces close to where they are used
- Use structs for data models; pointer receivers for methods that modify state

### Error Handling
- Return errors explicitly; do not use `panic` for expected error paths
- Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Handle errors at the appropriate level (don't ignore with `_`)
- Use custom error types only when error-specific handling is needed

### Concurrency
- Use `sync.WaitGroup` for goroutine synchronization
- Always close channels or use proper signal handling for cleanup
- Use `context.Context` for cancellation propagation when appropriate

### Testing Patterns
- Use `testify/suite` for test suites with Setup/TearDown
- Use `testify/assert` for assertions
- Create mock providers for interface-based dependencies
- Use table-driven tests for multiple test cases
- Tests should be in `*_test.go` files in the same package or `_test` subpackage

### Code Organization
- `cmd/`: CLI commands (Cobra commands)
- `internal/`: Private application code
- `internal/metrics/`: System metrics collection
- Main entry point in `main.go`

### Linter Configuration
The project uses golangci-lint with these enabled linters:
- `copyloopvar`, `dupl`, `errcheck`, `goconst`, `gocyclo`, `govet`
- `ineffassign`, `lll`, `misspell`, `nakedret`, `prealloc`, `revive`
- `staticcheck`, `unconvert`, `unparam`, `unused`

### Key Dependencies
- `github.com/spf13/cobra`: CLI framework
- `github.com/shirou/gopsutil/v3`: System metrics
- `github.com/stretchr/testify`: Testing utilities

## Common Tasks

### Adding a New Command
1. Create new file in `cmd/` (e.g., `cmd/stats.go`)
2. Define command using Cobra patterns
3. Register in `cmd/root.go` via `init()` or explicit AddCommand

### Adding New Metrics
1. Add method to `MetricsProvider` interface in `internal/metrics/metrics.go`
2. Implement in `RealProvider`
3. Add to `Data` struct and `CollectAll` function

### Adding Tests
1. Use existing test patterns in `*_test.go` files
2. Follow naming: `TestSuiteName_MethodName` or `TestFunctionName`
3. Mock interfaces for isolated unit testing
