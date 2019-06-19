package mget

type chunkInfo struct {
	idx          int
	url          string
	baseFileName string
	size         uint64
	rangeLow     uint64
	rangeHigh    uint64
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

type chunkData struct {
	chunkPath string
	idx       int
}
