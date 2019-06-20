NAME := vget

GO ?= go
BUILD_DIR=./build
BINARY ?= $(BUILD_DIR)/$(NAME)

.PHONY: all
all: clean test build

.PHONY: test
test:
	$(GO) test -v -race ./...

.PHONY: test_integration
test_integration:
	$(GO) test -v -tags=integration ./...

.PHONY: test_cover
test_cover:
	$(GO) test -v -cover ./...

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINARY)

.PHONY: build
build:
	$(GO) build -o $(BINARY) -v
