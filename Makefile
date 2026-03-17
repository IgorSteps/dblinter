
CMD_DIR := ./cmd
DIST_DIR := ./dist
BINARY_NAME := dblinter

run:
	go build -o $(DIST_DIR)/$(BINARY_NAME) $(CMD_DIR)/$(BINARY_NAME)
	$(DIST_DIR)/$(BINARY_NAME) ./tests/...