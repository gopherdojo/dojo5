/*
Provides the ability to perform image conversion with gocon.
*/
package gocon

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/pkg/convert"
	"github.com/gopherdojo/dojo5/kadai2/nagaa052/pkg/search"
)

const (
	ExitOK = iota
	ExitError
)

const (
	SuccessConvertFileMessageFmt string = "convert file : %s\n"
)

type gocon struct {
	SrcDir string
	Options
	outStream, errStream io.Writer
}

// Options is a specifiable option. Default is DefaultOptions.
type Options struct {
	FromFormat ImgFormat
	ToFormat   ImgFormat
	DestDir    string
}

// DefaultOptions is the default value of Options.
var DefaultOptions = Options{
	FromFormat: JPEG,
	ToFormat:   PNG,
	DestDir:    "out",
}

// New is Generate a new gocon.
func New(srcDir string, opt Options, outStream, errStream io.Writer) (*gocon, error) {
	if srcDir == "" {
		return nil, fmt.Errorf("target path is required")
	}

	srcDir, err := filepath.Abs(srcDir)
	if err != nil {
		return nil, fmt.Errorf("Invalid directory specification")
	}

	if _, err := os.Stat(srcDir); err != nil {
		return nil, err
	}

	if opt.FromFormat == "" {
		opt.FromFormat = DefaultOptions.FromFormat
	}

	if opt.ToFormat == "" {
		opt.ToFormat = DefaultOptions.ToFormat
	}

	if opt.DestDir == "" {
		opt.DestDir = DefaultOptions.DestDir
	}

	opt.DestDir, err = filepath.Abs(opt.DestDir)
	if err != nil {
		return nil, fmt.Errorf("Invalid directory specification")
	}

	return &gocon{
		SrcDir:    srcDir,
		Options:   opt,
		outStream: outStream,
		errStream: errStream,
	}, nil
}

// Run is executes gocon.
func (gc *gocon) Run() int {

	wg := &sync.WaitGroup{}
	chw := make(chan string)
	defer close(chw)

	ext, err := gc.FromFormat.GetExtentions()
	if err != nil {
		fmt.Fprintf(gc.errStream, "%+v\n", err)
		return ExitError
	}

	err = search.WalkWithExtHandle(gc.SrcDir, ext,
		func(srcImgPath string, info os.FileInfo, err error) {
			if err != nil {
				fmt.Fprintf(gc.errStream, "%+v\n", err)
				return
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				destFilePath, err := gc.convert(srcImgPath, info)
				if err != nil {
					fmt.Fprintf(gc.errStream, "%+v\n", err)
					return
				}

				fmt.Fprintf(gc.outStream, SuccessConvertFileMessageFmt, destFilePath)
			}()
		})
	if err != nil {
		fmt.Fprintf(gc.errStream, "%+v\n", err)
		return ExitError
	}

	wg.Wait()
	return ExitOK
}

func (gc *gocon) convert(srcImgPath string, info os.FileInfo) (string, error) {
	srcImgDir := strings.Replace(srcImgPath, "/"+info.Name(), "", -1)
	destDir := strings.Replace(srcImgDir, gc.SrcDir, gc.DestDir, -1)

	con, err := convert.New(srcImgPath, destDir)
	if err != nil {
		return "", err
	}

	switch gc.ToFormat {
	case JPEG:
		return con.ToJpeg(&convert.JpegOptions{})
	case PNG:
		return con.ToPng()
	default:
		return "", fmt.Errorf("There is no convertible format")
	}
}
