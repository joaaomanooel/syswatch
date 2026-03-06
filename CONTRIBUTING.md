# Contributing to SysWatch

Thank you for your interest in contributing to SysWatch! This document outlines the process for contributing to this project.

## Code of Conduct

Please be respectful and considerate of others when contributing.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported
2. Create a detailed issue with:
   - Clear title
   - Steps to reproduce
   - Expected vs actual behavior
   - Environment details (OS, Go version)

### Suggesting Features

1. Open an issue with the `enhancement` label
2. Describe the feature you'd like to see
3. Explain why it would be beneficial

### Pull Requests

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Run tests: `make test`
5. Run linter: `make lint`
6. Commit with clear messages
7. Push to your fork
8. Submit a Pull Request

## Development Setup

### Prerequisites

- Go 1.24+
- Make
- golangci-lint (optional)

### Running Tests

```bash
make test
```

### Running Linter

```bash
make lint
```

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build
```

## Style Guidelines

- Follow standard Go conventions
- Use 2-space indentation
- Run `go fmt` before committing
- Add tests for new features
- Maintain 90% test coverage

## Commit Messages

Use clear, descriptive commit messages:

```
feat: add new monitoring feature
fix: resolve CPU usage calculation bug
docs: update installation instructions
test: add tests for metrics collection
```

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
