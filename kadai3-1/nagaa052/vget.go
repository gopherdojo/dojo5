package vget

import (
	"fmt"
	"io"
	"runtime"
	"time"
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
	Procs   int
	TimeOut time.Duration
	DestDir string
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

	return &Vget{
		URL:       url,
		Options:   opt,
		outStream: outStream,
		errStream: errStream,
	}, nil
}

func (v *Vget) Download() int {

	// URLが正しいか
	// 分割ダウンロードが可能か
	// sizeの取得
	getLength := func() (int, error) {
		return 100, nil
	}
	size, err := getLength()
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	// 開始前のディレクトリ準備
	// 途中からの場合は進捗の取得
	makePayloads := func(size int) interface{} {
		return nil
	}
	jobs := makePayloads(size)

	/**
	- 分割数の決定,並列処理の開始
	*/
	publish := func(job interface{}) (interface{}, error) {
		return nil, nil
	}
	result, err := publish(jobs)
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	bindFiles := func(result interface{}) error {
		return nil
	}
	err = bindFiles(result)
	if err != nil {
		fmt.Fprintf(v.errStream, "%+v\n", err)
		return ExitError
	}

	return ExitOK
}
