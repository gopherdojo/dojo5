build: cmd/*.go typinggame/*.go
	GO111MODULE=on go build -o bin/typinggame cmd/main.go

fmt:
	go fmt ./...
	go vet ./...

check:
	GO111MODULE=on go test ./... -v

coverage:
	GO111MODULE=on go test ./... -cover
