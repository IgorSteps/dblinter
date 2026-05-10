CMD_DIR := ./cmd
DIST_DIR := ./dist
BINARY_NAME := dblinter

## build: build dblinter
build:
	go build -o $(DIST_DIR)/$(BINARY_NAME) $(CMD_DIR)/$(BINARY_NAME)

## run: run dblinter
run: build
	$(DIST_DIR)/$(BINARY_NAME) ./tests/...

## unit: run unit tests
unit:
	go test ./... -timeout 30s
