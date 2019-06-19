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
