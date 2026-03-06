# SysWatch

<p align="center">
  <img src="https://img.shields.io/badge/version-0.0.1-blue" alt="Version">
  <img src="https://img.shields.io/badge/go-1.24-blue" alt="Go Version">
  <img src="https://img.shields.io/badge/coverage-90%25-green" alt="Coverage">
  <img src="https://img.shields.io/github/license/joaaomanooel/syswatch" alt="License">
</p>

SysWatch is a powerful CLI tool for real-time system monitoring. It provides instant visibility into CPU usage, memory consumption, disk usage, and process count directly in your terminal.

## Features

- **Real-time Monitoring**: Watch CPU, memory, disk, and process metrics update in real-time
- **Cross-Platform**: Works on macOS, Linux, and Windows
- **Lightweight**: Single binary, no external dependencies
- **Customizable Intervals**: Adjust update frequency to suit your needs

## Installation

### macOS

```bash
# Using Homebrew
brew install joaaomanooel/homebrew-syswatch/syswatch

# Or download manually
curl -L https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-darwin-amd64.zip -o syswatch.zip
unzip syswatch.zip
./syswatch
```

### Linux

```bash
# Using APT (Debian/Ubuntu)
# Add the repository (coming soon)

# Using Snap
sudo snap install syswatch --classic

# Using RPM (Fedora/RHEL)
sudo rpm -i https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-*.rpm

# Download manually
curl -L https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-linux-amd64.zip -o syswatch.zip
unzip syswatch.zip
./syswatch
```

### Windows

```bash
# Download from GitHub Releases
# Extract the ZIP file and run syswatch.exe
```

## Usage

### Basic Usage

```bash
syswatch monitor
```

### Custom Interval

```bash
# Update every 5 seconds
syswatch monitor -i 5s

# Update every 500 milliseconds
syswatch monitor -i 500ms
```

### Version

```bash
syswatch version
```

## Development

### Prerequisites

- Go 1.24 or later
- Make

### Build from Source

```bash
# Clone the repository
git clone https://github.com/joaaomanooel/syswatch.git
cd syswatch

# Build for your current platform
make build

# Or run directly
make run
```

### Running Tests

```bash
make test
```

### Running Linter

```bash
make lint
```

### Building for All Platforms

```bash
# Builds for macOS, Linux, Windows (amd64 + arm64)
make build
```

## Configuration

SysWatch uses sensible defaults that work well for most use cases. No configuration file is required.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [gopsutil](https://github.com/shirou/gopsutil) for system metrics
- [Cobra](https://github.com/spf13/cobra) for CLI framework
