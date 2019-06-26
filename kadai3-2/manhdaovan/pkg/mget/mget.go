package mget

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

// DefaultExitSigs is default signals that make this tool exit
var DefaultExitSigs = []os.Signal{syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT}

// MGet represents struct of downloader
type MGet struct {
	workerNum  uint
	exitSigs   []os.Signal
	UserAgent  string
	Referer    string
	httpClient *http.Client

	dstDir  string
	dstFile string

	sigChan  chan os.Signal
	errChan  chan error
	doneChan chan struct{}
}

// NewMGet inits new MGet
func NewMGet(httpClient *http.Client, workers uint, exitSigs []os.Signal, userAgent, referer string) *MGet {
	m := &MGet{
		workerNum:  workers,
		exitSigs:   exitSigs,
		UserAgent:  userAgent,
		Referer:    referer,
		httpClient: httpClient,
	}
	m.init()
	return m
}

// Download returns downloaded file path and error of downloading.
func (m *MGet) Download(ctx context.Context, dst, url string) (string, error) {
	m.setDirAndFileName(dst, url)
	defer m.shutdown()
	return m.download(ctx, url)
}

func (m *MGet) init() {
	if len(m.exitSigs) == 0 { // no signal given
		m.exitSigs = DefaultExitSigs
	}
	m.errChan = make(chan error, 2)     // cap for one error and closing
	m.doneChan = make(chan struct{}, 2) // cap for one struct{} and closing
	m.sigChan = make(chan os.Signal, 2) // cap for one sig and closing
	signal.Notify(m.sigChan, m.exitSigs...)
}

func (m *MGet) setDirAndFileName(dstPath, url string) {
	m.dstDir, m.dstFile = parseDirAndFileName(dstPath)
	if m.dstFile == "" { // dstPath is a directory, no custom file name given
		m.dstFile = parseFileName(url)
	}
}

func (m *MGet) download(ctx context.Context, url string) (savedFilePath string, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	res, err := ctxhttp.Head(ctx, m.httpClient, url)
	if err != nil {
		err = errors.Wrap(err, "failed to head request: "+url)
		return
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		fmt.Println("[WARN]: server not support concurrency downloading")
		m.workerNum = 1
	}

	fileSize, err := strconv.ParseUint(res.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		err = errors.Wrapf(err, "cannot get file size")
		return
	}

	chunkSize := fileSize / uint64(m.workerNum)
	if (fileSize % uint64(m.workerNum)) != 0 {
		m.workerNum++ // recalculate workers to fit with chunks
	}

	go func() {
		eg, ctx := errgroup.WithContext(ctx)
		for i := 0; i < int(m.workerNum); i++ {
			chunk := newChunkInfo(i, url, fileSize, chunkSize)
			eg.Go(func() error {
				return m.downloadChunk(ctx, chunk)
			})
		}

		if egErr := eg.Wait(); err != nil {
			m.errChan <- errors.Wrapf(egErr, "error on chunks downloading")
			return
		}

		m.doneChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		cancel()
		err = fmt.Errorf("timeout")
	case sig := <-m.sigChan:
		err = fmt.Errorf("got quit sig: %s", sig.String())
	case downloadErr := <-m.errChan:
		err = downloadErr
	case <-m.doneChan:
		err = nil
	}

	if err != nil {
		return
	}

	savedFilePath, err = m.mergeChunks()
	return
}

func (m *MGet) downloadChunk(ctx context.Context, chunk *chunkInfo) error {
	// create get request
	req, err := http.NewRequest(http.MethodGet, chunk.url, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to init NewRequest for chunk: %d", chunk.idx)
	}

	// set download ranges
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", chunk.rangeLow, chunk.rangeHigh))
	// set useragent
	if m.UserAgent != "" {
		req.Header.Set("User-Agent", m.UserAgent)
	}
	// set referer
	if m.Referer != "" {
		req.Header.Set("Referer", m.Referer)
	}

	res, reqErr := ctxhttp.Do(ctx, m.httpClient, req)
	if reqErr != nil {
		return errors.Wrapf(reqErr, "cannot get data for chunk %d", chunk.idx)
	}
	defer res.Body.Close()

	chunkPath := chunkPath(m.dstDir, m.dstFile, chunk.idx)
	chunkFile, createErr := os.Create(chunkPath)
	if createErr != nil {
		return errors.Wrapf(reqErr, "cannot create chunk %d at %s", chunk.idx, chunkPath)
	}
	defer chunkFile.Close()

	if _, err := io.Copy(chunkFile, res.Body); err != nil {
		return errors.Wrapf(reqErr, "cannot save data for chunk %d to %s", chunk.idx, chunkPath)
	}

	return nil
}

func (m *MGet) mergeChunks() (string, error) {
	savedPath := filepath.Join(m.dstDir, m.dstFile)
	resultFile, err := os.Create(savedPath)
	if err != nil {
		return "", errors.Wrapf(err, "cannot create result file")
	}
	defer resultFile.Close()

	mergeChunk := func(srcFile *os.File, chunkPath string) error {
		chunkData, err := os.Open(chunkPath)
		defer chunkData.Close()

		if err != nil {
			return err
		}
		if _, err := io.Copy(resultFile, chunkData); err != nil {
			return err
		}

		return nil
	}

	for i := 0; i < int(m.workerNum); i++ {
		chunkIdx := i
		chunkPath := chunkPath(m.dstDir, m.dstFile, chunkIdx)
		if err := mergeChunk(resultFile, chunkPath); err != nil {
			return "", errors.Wrapf(err, "cannot merge chunk: %d", chunkIdx)
		}
	}

	return savedPath, nil
}

func (m *MGet) shutdown() {
	close(m.sigChan)
	close(m.errChan)
	close(m.doneChan)
	m.cleanChunks()
}

func (m *MGet) cleanChunks() {
	for chunkIdx := 0; chunkIdx < int(m.workerNum); chunkIdx++ {
		chunkPath := chunkPath(m.dstDir, m.dstFile, chunkIdx)
		if err := os.Remove(chunkPath); err != nil {
			fmt.Fprintf(os.Stderr, "error on remove chunk %s data: %v\n", chunkPath, err)
		}
	}
}
