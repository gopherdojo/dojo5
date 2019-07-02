BIN := godl

.PHONY: build
build:
	go build -o $(BIN) cmd/main.go

.PHONY: clean
clean:
	rm $(BIN)
	go clean

.PHONY: test
test: build
	go test -race -v
