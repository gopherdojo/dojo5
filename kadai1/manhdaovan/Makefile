build: ./cmd/main.go
	- go build -a -o ./bin/imgconv ./cmd/main.go
test: ./pkg/imgconv/converter_test.go ./pkg/imgconv/encoder_test.go
	- go test ./pkg/...