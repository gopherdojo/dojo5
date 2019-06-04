NAME := gocon

GO ?= go
BUILD_DIR=./build
BINARY ?= $(BUILD_DIR)/$(NAME)

.PHONY: all
all: clean test build

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINARY)

.PHONY: build
build:
	$(GO) build -o $(BINARY) -v
