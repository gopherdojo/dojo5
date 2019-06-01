ROOT=github.com/gopherdojo/dojo5/kadai1/ebiiim
MAIN=cmd/imgconv.go
TEST=pkg/...

CMD_ALL=all
CMD_BUILD=build
CMD_TEST=test

.PHONY: ${CMD_ALL}
${CMD_ALL}: ${CMD_TEST} ${CMD_BUILD}

.PHONY: ${CMD_BUILD}
${CMD_BUILD}: ${MAIN}
	go build ${GOPATH}/src/${ROOT}/$?

.PHONY: ${CMD_TEST}
${CMD_TEST}:
	go test -cover ${ROOT}/${TEST}
