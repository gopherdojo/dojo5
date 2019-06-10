bin/convert: cmd/convert/*.go imgconv/*.go
	GO111MODULE=on go build -o bin/convert cmd/convert/main.go

fmt:
	go fmt ./...
	go vet ./...

check:
	GO111MODULE=on go test ./imgconv/... -v
