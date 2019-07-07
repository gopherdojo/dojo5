BIN := godl
GOCOV := coverageprofile

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

.PHONY: cover
cover: build
	go test -race ./... -coverprofile=$(GOCOV)
	go tool cover -html=$(GOCOV)
	rm $(GOCOV)
