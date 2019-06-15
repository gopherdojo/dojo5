BIN := gotp
GODOCCMD := godoc
GODOCPORT := 6060
GODOCPATH := /github.com/matsuyoshi30/dojo5/kadai3-1/matsuyoshi30/
GOCOV := coverageprofile

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: clean
clean:
	rm $(BIN)
	go clean

.PHONY: test
test: build
	go test -race ./...

.PHONY: cover
cover: build
	go test -race ./... -coverprofile=$(GOCOV)
	go tool cover -html=$(GOCOV)
	rm $(GOCOV)

.PHONY: doc
doc:
	open http://localhost:$(GODOCPORT)/pkg$(GODOCPATH)
	$(GODOCCMD) -http :$(GODOCPORT)
