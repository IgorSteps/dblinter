CMD_DIR := ./cmd
DIST_DIR := ./dist
BINARY_NAME := dblinter

GOLANGCI_LINT := golangci-lint
GOLANGCI_VERSION := v2.12.2 # TODO: make this aligned with CI.

## tools: installs necessary tools to run quality check
tools:
	@if ! command -v $(GOLANGCI_LINT) >/dev/null 2>&1; then \
		echo "Installing $(GOLANGCI_LINT)..."; \
		curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_VERSION); \
	else \
		echo "$(GOLANGCI_LINT) already installed"; \
	fi

## build: build dblinter
build:
	go build -o $(DIST_DIR)/$(BINARY_NAME) $(CMD_DIR)/$(BINARY_NAME)

## run: run dblinter
run: build
	$(DIST_DIR)/$(BINARY_NAME) ./tests/...

## unit: run unit tests
unit: mocks
	go test ./... -timeout 30s

lint: tools
	golangci-lint run --config .golangci.yaml

## mocks: generate mocks
mocks:
	mockery --config .mockery.yml

functional:
	 $(DIST_DIR)/$(BINARY_NAME) ./tests/...
