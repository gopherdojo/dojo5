package mypget

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Downloader struct {
	url *url.URL
}

func New(url *url.URL) *Downloader {
	return &Downloader{
		url: url,
	}
}

func (d *Downloader) Execute() error {
	fmt.Println("Hello, split downloader!")
	err := d.download()
	return err
}

func (d *Downloader) download() error {
	req, err := http.NewRequest("GET", d.url.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !acceptBytesRanges(resp) {
		return errors.New("split download is not supported in this response")
	}
	length := int(resp.ContentLength)
	ranges := splitToRanges(length)
	fmt.Println(ranges) // TODO: split downloading by ranges

	// FIXME: create proper directory for downloading
	f, err := os.Create("./output.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func acceptBytesRanges(resp *http.Response) bool {
	return resp.Header.Get("Accept-Ranges") == "bytes"
}

func splitToRanges(length int) []string {
	rangeNum := 4 // FIXME: dynamically get range's number

	var ranges []string
	var rangeStart, rangeEnd int
	rangeLength := length / rangeNum

	for i := 0; i < rangeNum; i++ {
		if i != 0 {
			rangeStart = rangeEnd + 1
		}
		rangeEnd = rangeStart + rangeLength

		if i == rangeNum - 1 && rangeEnd != length {
			rangeEnd = length
		}

		ranges = append(ranges, fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	}
	return ranges
}
