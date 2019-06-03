package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo5/kadai1/manhdaovan/pkg/imgconv"
)

const (
	exitSuccess = 0
	exitFailure = 1
)

func main() {
	ca := parseArgs()
	if err := ca.validate(); err != nil {
		fmt.Printf("error: %v\n", err)
		printHelp()
		os.Exit(exitFailure)
	}

	if ca.srcImgType == ca.destImgType {
		fmt.Println("source and destination image type are the same. Nothing to to.")
		printHelp()
		os.Exit(exitSuccess)
	}

	err := imgconv.ConvertDir(ca.dirPath,
		imgconv.ImgType(ca.srcImgType),
		imgconv.ImgType(ca.destImgType),
		true)
	if err != nil {
		fmt.Println("error: ", err)
		printHelp()
	}
}

type cliArguments struct {
	dirPath     string
	srcImgType  string
	destImgType string
	skipErr     bool
}

func (ca cliArguments) validate() error {
	if ca.dirPath == "" {
		return fmt.Errorf("no given directory")
	}
	if _, err := os.Stat(ca.dirPath); err != nil {
		return fmt.Errorf("invalid directory %s: \n%+v", ca.dirPath, err)
	}

	if ca.srcImgType == "" || ca.destImgType == "" {
		return fmt.Errorf("both source and destination image type must be set")
	}

	at := imgconv.GetSupportSrcImgTypes()
	if !at.IsSupport(imgconv.ImgType(ca.srcImgType)) {
		return fmt.Errorf("not support source image type: %s", ca.srcImgType)
	}
	if !at.IsSupport(imgconv.ImgType(ca.destImgType)) {
		return fmt.Errorf("not support destination image type: %s", ca.destImgType)
	}

	return nil
}

func parseArgs() cliArguments {
	ca := cliArguments{}

	flag.StringVar(&ca.srcImgType, "s", string(imgconv.ImgTypePNG), "source image type")
	flag.StringVar(&ca.destImgType, "d", string(imgconv.ImgTypeJPEG), "destination image type")

	flag.BoolVar(&ca.skipErr, "k", false, "Skip error and continue while converting")
	flag.Parse()

	ca.dirPath = flag.Arg(0)
	return ca
}

func printHelp() {
	helpStr := `
imgconv -- convert image to other image type
usage: $./bin/imgconv [-s] [-d] [-k] dir
options:
  -s string
	Source image type. jpeg, png and gif are supported
  -d string
  	Destination image type. jpeg, png and gif are supported
  -k
	Skip error and continue while converting
`
	fmt.Println(helpStr)
}
