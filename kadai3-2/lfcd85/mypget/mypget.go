package mypget

import (
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
	err := d.Download()
	return err
}

func (d *Downloader) Download() error {
	req, err := http.NewRequest("GET", d.url.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// FIXME: create proper directory for downloading
	f, err := os.Create("./output.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}
