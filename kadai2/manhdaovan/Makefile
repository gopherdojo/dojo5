COVERAGE_VIEW_MODE=html# or "func"
GOMOD=on#or "off" or "auto"

install:
	- GO111MODULE=${GOMOD} go mod vendor
	- GO111MODULE=${GOMOD} go mod tidy
build: ./cmd/main.go
	- GO111MODULE=${GOMOD} go build -a -o ./bin/imgconv ./cmd/main.go
test: ./pkg/imgconv/converter_test.go ./pkg/imgconv/img_picker_test.go
	- GO111MODULE=${GOMOD} go test ./pkg/...
coverage:
	- GO111MODULE=${GOMOD} go test -coverprofile=cover.out ./...
	- GO111MODULE=${GOMOD} go tool cover -${COVERAGE_VIEW_MODE} cover.out
