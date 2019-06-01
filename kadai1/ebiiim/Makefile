ROOT=github.com/gopherdojo/dojo5/kadai1/ebiiim
MAIN=cmd/imgconv.go
TEST=test/...

CMD_BUILD=build
CMD_TEST=test

.PHONY: ${CMD_BUILD}
${CMD_BUILD}: ${MAIN}
	go build ${GOPATH}/src/${ROOT}/$?

.PHONY: ${CMD_TEST}
${CMD_TEST}:
	go test ${ROOT}/${TEST}
