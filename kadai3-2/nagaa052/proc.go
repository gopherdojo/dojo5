package vget

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gopherdojo/dojo5/kadai3-2/nagaa052/pkg/executor"
	"github.com/gopherdojo/dojo5/kadai3-2/nagaa052/pkg/request"
)

var _ executor.Payload = &Proc{}

type Proc struct {
	URL     string
	OutFile string
	Index   int
	From    int64
	To      int64
}

func NewProc(url, outFile string, from, to int64, index int) *Proc {
	return &Proc{
		URL:     url,
		OutFile: outFile,
		Index:   index,
		From:    from,
		To:      to,
	}
}

func (p *Proc) Execute(ctx context.Context) error {
	r := request.Range{}
	return r.Download(ctx, p.URL, p.From, p.To, p.OutFile)
}

func GetProcPath(dir, filename string, index int) string {
	return fmt.Sprintf("%s/%s.%d", dir, filename, index)
}

func IxExistProcFile(dir, filename string, index int, fileSize int64) bool {
	filePath := GetProcPath(dir, filename, index)
	if info, err := os.Stat(filePath); err == nil {
		if info.Size() == fileSize {
			return true
		}
	}

	return false
}

func MargeProcFiles(dir, filename string, outFile string, procsCount int) error {
	fh, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer fh.Close()

	for i := 0; i < procsCount; i++ {
		procFile := GetProcPath(dir, filename, i)
		subfp, err := os.Open(procFile)
		if err != nil {
			return err
		}
		defer subfp.Close()

		io.Copy(fh, subfp)
		if err := os.Remove(procFile); err != nil {
			return err
		}
	}
	return nil
}
