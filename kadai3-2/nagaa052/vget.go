package vget

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-2/nagaa052/pkg/executor"

	"github.com/gopherdojo/dojo5/kadai3-2/nagaa052/pkg/request"
)

const (
	ExitOK = iota
	ExitError
)

type Vget struct {
	URL string
	Options
	outStream, errStream io.Writer
}

// Options is a specifiable option. Default is DefaultOptions.
type Options struct {
	Procs    int
	TimeOut  time.Duration
	DestDir  string
	FileName string
}

var DefaultOptions = Options{
	Procs:   runtime.NumCPU(),
	TimeOut: time.Duration(60) * time.Second,
	DestDir: "download",
}

func New(url string, opt Options, outStream, errStream io.Writer) (*Vget, error) {
	if url == "" {
		return nil, fmt.Errorf("target url is required")
	}

	if opt.Procs == 0 {
		opt.Procs = DefaultOptions.Procs
	}

	if opt.TimeOut == 0 {
		opt.TimeOut = DefaultOptions.TimeOut
	}

	if opt.DestDir == "" {
		opt.DestDir = DefaultOptions.DestDir
	}

	destDir, err := filepath.Abs(opt.DestDir)
	if err != nil {
		return nil, fmt.Errorf("Invalid directory specification")
	}
	opt.DestDir = destDir

	if opt.FileName == "" {
		sURL := strings.Split(url, "/")
		opt.FileName = sURL[len(sURL)-1]
	}

	return &Vget{
		URL:       url,
		Options:   opt,
		outStream: outStream,
		errStream: errStream,
	}, nil
}

func (v *Vget) Download() int {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &request.Range{}
	size, err := r.GetContentLength(ctx, v.URL)
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	procs, err := v.getProcs(size)
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	if _, err := os.Stat(v.DestDir); err != nil {
		if err := os.MkdirAll(v.DestDir, 0755); err != nil {
			fmt.Fprintf(v.errStream, "%+v\n", err)
			return ExitError
		}
	}

	ex := executor.New(v.Options.Procs, v.Options.TimeOut)
	for _, proc := range procs {
		ex.AddPayload(proc)
	}

	fmt.Fprintf(v.outStream, "Start download.\n")
	err = ex.Start()
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	outfile := fmt.Sprintf("%s/%s", v.DestDir, v.FileName)
	err = MargeProcFiles(v.DestDir, v.FileName, outfile, v.Procs)
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	fmt.Fprintf(v.outStream, "Download finish\n")
	return ExitOK
}

func (v *Vget) getProcs(size int64) ([]*Proc, error) {
	procs := []*Proc{}

	procSize := size / int64(v.Procs)
	for i := 0; i < v.Procs; i++ {
		if !IxExistProcFile(v.DestDir, v.FileName, i, procSize) {
			from := procSize * int64(i)
			to := procSize * int64(i+1)
			proc := NewProc(v.URL, GetProcPath(v.DestDir, v.FileName, i), from, to, i)
			procs = append(procs, proc)
		}
	}
	return procs, nil
}
