bin/convert: cmd/convert/*.go imgconv/*.go
	go build -o bin/convert cmd/convert/main.go

fmt:
	go fmt ./...
	go vet ./...

check:
	go test ./... -v
