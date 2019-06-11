BIN := goconv
GODOCCMD := godoc
GODOCPORT := 6060
GODOCPATH := /github.com/matsuyoshi30/dojo5/kadai2/matsuyoshi30/

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: clean
clean:
	rm $(BIN)
	go clean

.PHONY: test
test: build
	go test -v ./...

.PHONY: doc
doc:
	open http://localhost:$(GODOCPORT)/pkg$(GODOCPATH)
	$(GODOCCMD) -http :$(GODOCPORT)
