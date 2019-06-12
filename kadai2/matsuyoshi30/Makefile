BIN := goconv
GODOCCMD := godoc
GODOCPORT := 6060
GODOCPATH := /github.com/matsuyoshi30/dojo5/kadai2/matsuyoshi30/
GOCOVPKG1 := conv
GOCOVPROFILE1 := $(GOCOVPKG1)profile
GOCOVPKG2 := opt
GOCOVPROFILE2 := $(GOCOVPKG2)profile

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
	rm $(GOCOVPROFILE1) $(GOCOVPROFILE2)
	go test -coverprofile=$(GOCOVPROFILE1) ./$(GOCOVPKG1)
	go test -coverprofile=$(GOCOVPROFILE2) ./$(GOCOVPKG2)
	go tool cover -html=$(GOCOVPROFILE1)
	go tool cover -html=$(GOCOVPROFILE2)

.PHONY: doc
doc:
	open http://localhost:$(GODOCPORT)/pkg$(GODOCPATH)
	$(GODOCCMD) -http :$(GODOCPORT)
