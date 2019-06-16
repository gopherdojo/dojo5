package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo5/kadai2/manhdaovan/pkg/imgconv"
)

const (
	exitSuccess = 0
	exitFailure = 1
)

func main() {
	ca := parseArgs()
	if err := ca.validate(); err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		printHelp()
		os.Exit(exitFailure)
	}

	if ca.srcImgType == ca.destImgType {
		fmt.Fprintln(os.Stderr, "source and destination image type are the same. Nothing to to.")
		printHelp()
		os.Exit(exitFailure)
	}

	conv, err := initConverter(ca)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		printHelp()
		os.Exit(exitFailure)
	}

	if err := conv.Convert(ca.path, imgconv.ImgType(ca.srcImgType)); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		printHelp()
		os.Exit(exitFailure)
	}
}

func initConverter(ca *cliArguments) (*imgconv.Converter, error) {
	var decoder imgconv.Decoder
	var encoder imgconv.Encoder
	var destImgExt imgconv.ImgExt
	var picker imgconv.DefaultImgPicker

	switch ca.srcImgType {
	case string(imgconv.ImgTypeGIF):
		decoder = imgconv.GIFDecoder{}
	case string(imgconv.ImgTypeJPEG):
		decoder = imgconv.JPEGDecoder{}
	case string(imgconv.ImgTypePNG):
		decoder = imgconv.PNGDecoder{}
	default:
		return nil, fmt.Errorf("no decoder for src img: %s", ca.srcImgType)
	}

	switch ca.destImgType {
	case string(imgconv.ImgTypeGIF):
		encoder = imgconv.GIFEncoder{}
		destImgExt = imgconv.ImgExtGIF
	case string(imgconv.ImgTypeJPEG):
		encoder = imgconv.JPEGEncoder{}
		destImgExt = imgconv.ImgExtJPEG
	case string(imgconv.ImgTypePNG):
		encoder = imgconv.PNGEncoder{}
		destImgExt = imgconv.ImgExtPNG
	default:
		return nil, fmt.Errorf("no encoder and extension for src img: %s", ca.destImgType)
	}

	return &imgconv.Converter{
		Enc:        encoder,
		Dec:        decoder,
		Picker:     picker,
		DestImgExt: destImgExt,
		SkipErr:    ca.skipErr,
		KeepSrcImg: ca.keepSrcImg,
	}, nil
}

type cliArguments struct {
	path        string
	srcImgType  string
	destImgType string
	skipErr     bool
	keepSrcImg  bool
}

func (ca *cliArguments) validate() error {
	if ca.path == "" {
		return fmt.Errorf("no given directory")
	}
	if _, err := os.Stat(ca.path); err != nil {
		return fmt.Errorf("invalid directory %s: \n%+v", ca.path, err)
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

func parseArgs() *cliArguments {
	ca := &cliArguments{}

	flag.StringVar(&ca.srcImgType, "s", string(imgconv.ImgTypePNG), "source image type")
	flag.StringVar(&ca.destImgType, "d", string(imgconv.ImgTypeJPEG), "destination image type")

	flag.BoolVar(&ca.skipErr, "skip-err", false, "Skip error and continue while converting")
	flag.BoolVar(&ca.keepSrcImg, "keep-src", false, "Keep source files after converted")
	flag.Parse()

	ca.path = flag.Arg(0)
	return ca
}

func printHelp() {
	helpStr := `
imgconv -- convert image to other image type
usage: $./bin/imgconv [-s] [-d] [-skip-err] [keep-src] dir
options:
  -s string
	Source image type. jpeg, png and gif are supported. Default png.
  -d string
  	Destination image type. jpeg, png and gif are supported. Default jpeg.
  -skip-err
	Skip error and continue while converting. Default false.
  -keep-src
    Keep source files after converted. Default false.
`
	fmt.Println(helpStr)
}
