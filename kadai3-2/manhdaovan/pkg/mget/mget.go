package mget

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
)

type MGet struct {
	WorkerNum int
	chunks    chan chunkData
	dstDir    string
}

func (m *MGet) Download(ctx context.Context, dst, url string) (savedFilePath string, err error) {
	m.dstDir = dst

	defer m.shutdown()

	res, err := ctxhttp.Head(ctx, http.DefaultClient, url)
	if err != nil {
		err = errors.Wrap(err, "failed to head request: "+url)
		return
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		fmt.Println("[WARN]: server not support concurrency downloading")
		m.WorkerNum = 1
	}

	fileSize, err := strconv.ParseUint(res.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		err = errors.Wrapf(err, "cannot get file size")
		return
	}

	chunkSize := fileSize / uint64(m.WorkerNum)
	if fileSize%uint64(m.WorkerNum) != 0 {
		m.WorkerNum++
	}
	m.chunks = make(chan chunkData, m.WorkerNum)

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < m.WorkerNum; i++ {
		fileChunk := newChunkInfo(i, url, fileSize, chunkSize)
		eg.Go(func() error {
			chunkData, err := m.downloadChunk(fileChunk)
			if err != nil {
				return err
			}
			m.chunks <- *chunkData
			return nil
		})
	}

	if egErr := eg.Wait(); err != nil {
		err = errors.Wrapf(egErr, "error on chunks downloading")
		return
	}

	savedFilePath, err = m.summary()
	return
}

func (m *MGet) downloadChunk(chunk *chunkInfo) (*chunkData, error) {
	// create get request
	req, err := http.NewRequest("GET", chunk.url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to split NewRequest for get: %d", chunk.idx)
	}

	// set download ranges
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", chunk.rangeLow, chunk.rangeHigh))

	// // set useragent
	// if p.useragent != "" {
	// 	req.Header.Set("User-Agent", p.useragent)
	// }

	// // set referer
	// if p.referer != "" {
	// 	req.Header.Set("Referer", p.referer)
	// }

	res, reqErr := http.DefaultClient.Do(req)
	if reqErr != nil {
		return nil, errors.Wrapf(reqErr, "cannot get data for chunk %d", chunk.idx)
	}

	chunkName := fmt.Sprintf("%s-%d.chunk", chunk.baseFileName, chunk.idx)
	chunkPath := filepath.Join(m.dstDir, chunkName)
	chunkFile, createErr := os.Create(chunkPath)
	if createErr != nil {
		return nil, errors.Wrapf(reqErr, "cannot create chunk %d at %s", chunk.idx, chunkPath)
	}
	if _, err := io.Copy(chunkFile, res.Body); err != nil {
		return nil, errors.Wrapf(reqErr, "cannot save data for chunk %d to %s", chunk.idx, chunkPath)
	}

	return &chunkData{
		chunkPath: chunkPath,
		idx:       chunk.idx,
	}, nil
}

func (m *MGet) summary() (string, error) {
	return "", nil
}

func (m *MGet) shutdown() {

}
