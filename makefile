# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=tasknova
MAIN_PACKAGE=.

# Output directories
BIN_DIR=./bin
MACOS_BIN=$(BIN_DIR)/$(BINARY_NAME)-darwin
LINUX_BIN=$(BIN_DIR)/$(BINARY_NAME)-linux
WINDOWS_BIN=$(BIN_DIR)/$(BINARY_NAME)-windows.exe
MACOS_BIN_ARM64=$(BIN_DIR)/$(BINARY_NAME)-darwin-arm64
RUN_BIN=$(BIN_DIR)/$(BINARY_NAME)

all: test build

build: clean build-macos build-macos-arm64 build-linux build-windows

# Test targets
test:
	gotestsum --format testdox --hide-summary output --format-hide-empty-pkg --packages=all -- -v ./... -coverprofile=coverage/coverage.out -covermode=atomic -coverpkg=./... && \
	go tool cover -html=coverage/coverage.out -o coverage/index.html && \
	go tool cover -func=coverage/coverage.out | grep total: | awk '{if (substr($$3, 1, length($$3)-1) < 90) { print "Coverage " $$3 " is below 90%"; exit 1 } else { print "Coverage " $$3 " meets minimum 90% requirement" }}'

test-coverage:
	gotestsum --format testdox --hide-summary output --format-hide-empty-pkg --packages=all -- -v ./... -coverprofile=coverage/coverage.out -covermode=atomic -coverpkg=./... && \
	go tool cover -html=coverage/coverage.out -o coverage/index.html && \
	go tool cover -func=coverage/coverage.out | tee coverage/coverage.txt
	@coverage=$$(go tool cover -func=coverage/coverage.out | grep total | awk '{print $$3}' | sed 's/%//'); \
	if [ $$(echo "$$coverage < 90" | bc -l) -eq 1 ]; then \
		echo "Code coverage is below 90% (current: $$coverage%)"; \
		exit 1; \
	fi


test-watch: ## Watch for changes and run tests
	gotestsum --watch --format testdox

test-verbose: ## Run tests in verbose mode
	gotestsum --format standard-verbose --hide-summary output --format-hide-empty-pkg --packages=all -- -v ./... -coverprofile=coverage/coverage.out -covermode=atomic -coverpkg=./...

clean:
	$(GOCMD) clean
	rm -rf $(BIN_DIR) dist coverage/coverage.out coverage/coverage.txt coverage/index.html

run:
	$(GOBUILD) -o $(RUN_BIN) -v
	./$(RUN_BIN)

deps:
	$(GOMOD) download && \
  go install gotest.tools/gotestsum@latest

tidy:
	$(GOMOD) tidy

# Create bin directory if it doesn't exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Cross compilation targets
build-macos: $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(MACOS_BIN) $(MAIN_PACKAGE)
	@echo "Built for macOS (amd64): $(MACOS_BIN)"

build-macos-arm64: $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(MACOS_BIN_ARM64) $(MAIN_PACKAGE)
	@echo "Built for macOS (arm64): $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64"

build-linux: $(BIN_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_BIN) $(MAIN_PACKAGE)
	@echo "Built for Linux: $(LINUX_BIN)"

build-windows: $(BIN_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_BIN) $(MAIN_PACKAGE)
	@echo "Built for Windows: $(WINDOWS_BIN)"

# Create a release with all binaries
release: build
	@echo "Release binaries created in $(BIN_DIR) directory"

.PHONY: all build test test-coverage test-watch test-verbose test-race test-bench test-nocache test-short test-timeout test-failed clean run deps tidy build-macos build-macos-arm64 build-linux build-windows release


.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run ./... --fix
