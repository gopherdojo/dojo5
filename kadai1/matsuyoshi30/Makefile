BIN := goconv

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: clean
clean:
	rm $(BIN)
	go clean

.PHONY: test
test: build
	go test -v ./conv
