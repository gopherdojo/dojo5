build:
	GO111MODULE=on go build -o bin/mypget cmd/main.go

PHONY: fmt
fmt:
	go fmt ./...

PHONY: check
check:
	GO111MODULE=on go test ./... -v

PHONY: coverage
coverage:
	GO111MODULE=on go test ./... -cover
