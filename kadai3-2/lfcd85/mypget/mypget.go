package mypget

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

const (
	tempDirName    = "partials"
	tempFilePrefix = "partial"
)

type Downloader struct {
	url    *url.URL
	ranges []string
}

func New(url *url.URL) *Downloader {
	return &Downloader{
		url: url,
	}
}

func (d *Downloader) Execute() error {
	err := d.download()
	return err
}

func (d *Downloader) download() error {
	bc := context.Background()
	ctx, cancel := context.WithCancel(bc)
	defer cancel()

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
	d.splitToRanges(length)

	tempDir, err := ioutil.TempDir("", tempDirName)
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	err = d.downloadByRanges(ctx, tempDir)
	if err != nil {
		return err
	}

	err = d.combine(tempDir)
	if err != nil {
		return err
	}

	return err
}

func acceptBytesRanges(resp *http.Response) bool {
	return resp.Header.Get("Accept-Ranges") == "bytes"
}

func (d *Downloader) splitToRanges(length int) {
	rangeNum := 4 // FIXME: dynamically get range's number

	var ranges []string
	var rangeStart, rangeEnd int
	rangeLength := length / rangeNum

	for i := 0; i < rangeNum; i++ {
		if i != 0 {
			rangeStart = rangeEnd + 1
		}
		rangeEnd = rangeStart + rangeLength

		if i == rangeNum-1 && rangeEnd != length {
			rangeEnd = length
		}

		ranges = append(ranges, fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	}
	d.ranges = ranges
}

func (d *Downloader) downloadByRanges(ctx context.Context, tempDir string) error {
	var eg errgroup.Group

	for i, r := range d.ranges {
		i, r := i, r
		eg.Go(func() error {
			req, err := http.NewRequest("GET", d.url.String(), nil)
			if err != nil {
				return err
			}
			req = req.WithContext(ctx)
			req.Header.Set("Range", r)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			err = validateStatusPartialContent(resp)
			if err != nil {
				return err
			}

			// FIXME: create proper directory for downloading
			partialPath := generatePartialPath(tempDir, i)
			fmt.Printf("Downloading %v (%v) ...\n", partialPath, r)

			f, err := os.Create(partialPath)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, resp.Body)
			return err
		})
	}
	return eg.Wait()
}

func validateStatusPartialContent(resp *http.Response) error {
	validStatusCode := http.StatusPartialContent
	if resp.StatusCode != validStatusCode {
		return fmt.Errorf("status code must be %d: actually was %d", validStatusCode, resp.StatusCode)
	}
	return nil
}

func generatePartialPath(tempDir string, i int) string {
	base := strings.Join([]string{tempFilePrefix, strconv.Itoa(i)}, "_")
	return strings.Join([]string{tempDir, base}, "/")
}

func (d *Downloader) combine(tempDir string) error {
	outputPath := d.getOutputFileName()
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("Combining partials to %v ...\n", outputPath)

	for i, _ := range d.ranges {
		partialPath := generatePartialPath(tempDir, i)
		partial, err := os.Open(partialPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, partial)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Downloader) getOutputFileName() string {
	base := filepath.Base(d.url.EscapedPath())
	switch base {
	case "/", ".", "":
		return "output"
	default:
		return base
	}
}
