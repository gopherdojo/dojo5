package mget

import (
	"fmt"
	"path/filepath"
)

type chunkInfo struct {
	idx       int
	url       string
	size      uint64
	rangeLow  uint64
	rangeHigh uint64
}

func newChunkInfo(idx int, url string, fileSize, chunkSize uint64) *chunkInfo {
	offset := uint64(idx) * chunkSize
	if offset+chunkSize > fileSize { // last chunk
		chunkSize = fileSize - offset
	}

	return &chunkInfo{
		idx:       idx,
		url:       url,
		size:      chunkSize,
		rangeLow:  offset,
		rangeHigh: offset + chunkSize - 1,
	}
}

func chunkPath(dstDir, dstFile string, idx int) string {
	chunkName := fmt.Sprintf("%s-%d.chunk", dstFile, idx)
	return filepath.Join(dstDir, chunkName)
}
