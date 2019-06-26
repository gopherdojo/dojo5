package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/gopherdojo/dojo5/kadai3-2/manhdaovan/pkg/mget"
)

const (
	exitErr = 1
)

func main() {
	ca := parseArgs()
	if err := ca.validate(); err != nil {
		fmt.Fprintln(os.Stderr, "given args error: ", err)
		printHelp()
		os.Exit(exitErr)
	}
	if _, err := os.Stat(ca.outPath); err != nil {
		fmt.Fprintf(os.Stderr, "invalid output path: %s, err: %+v\n", ca.outPath, err)
		printHelp()
		os.Exit(exitErr)
	}

	downloader := mget.NewMGet(http.DefaultClient, ca.numWorkers, mget.DefaultExitSigs, "", "")
	outputPath, err := downloader.Download(context.Background(), ca.outPath, ca.url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "download error: %+v\n", err)
		printHelp()
		os.Exit(exitErr)
	}

	fmt.Println("Download completed. Output: ", outputPath)
}

type cliArgs struct {
	numWorkers uint
	outPath    string
	url        string
}

func (ca *cliArgs) validate() error {
	if ca.numWorkers == 0 {
		return fmt.Errorf("number worker must be greater than 0")
	}
	if ca.url == "" {
		return fmt.Errorf("no download url given")
	}

	if ca.outPath == "" {
		ca.outPath = "./out/"
	}

	return nil
}

func parseArgs() *cliArgs {
	var ca cliArgs
	flag.UintVar(&ca.numWorkers, "w", uint(runtime.NumCPU()), "number of workers")
	flag.StringVar(&ca.outPath, "o", "", "output file path")
	flag.Parse()
	ca.url = flag.Arg(0) // last arg is download url

	return &ca
}

func printHelp() {
	helpStr := `
mget -- a simple concurrency download tool
usage: $./bin/mget [-w] [-o] download-url
options:
	-w unit
	    Number of workers that download in concurrent
	-o string
		Output of downloaded file. It can be a directory or a file path.
`
	fmt.Println(helpStr)
}
