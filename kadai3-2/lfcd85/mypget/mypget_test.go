package mypget_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-2/lfcd85/mypget"
)

func TestDownloader_Execute(t *testing.T) {
	cases := []struct {
		urlStr   string
		splitNum int
	}{
		{"https://golang.org/doc/gopher/frontpage.png", 8}, // TODO: replace URL with local test server
	}

	for _, c := range cases {
		c := c
		t.Run(c.urlStr, func(t *testing.T) {
			d := initDownloader(t, c.urlStr, c.splitNum)
			if err := d.Execute(); err != nil {
				t.Errorf("failed to execute split downloader: %v", err)
			}
			deleteOutputFile(t, d)
		})
	}
}

func initDownloader(t *testing.T, urlStr string, splitNum int) *mypget.Downloader {
	t.Helper()

	url, err := url.Parse(urlStr)
	if err != nil {
		t.Errorf("failed to parse URL for testing: %v", err)
	}

	return mypget.New(url, splitNum)
}

func deleteOutputFile(t *testing.T, d *mypget.Downloader) {
	t.Helper()

	outputPath := d.ExportOutputPath()
	if outputPath != "" {
		err := os.Remove(outputPath)
		if err != nil {
			t.Errorf("failed to remove output file: %v", err)
		}
	}
}
