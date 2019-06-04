BIN := iconv

.PHONY: test
test:
	go test -cover -v ./...

.PHONY: build
build: test
	go build -o $(BIN)

.PHONY: clean
clean:
	rm $(BIN)
	go clean