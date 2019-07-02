package godl

import (
	"errors"
	"net/http"
)

func (g *Godl) ValidateStatus(resp *http.Response) bool {
	return resp.StatusCode == http.StatusPartialContent
}

func (g *Godl) ValidateURL(url string) error {
	err := g.checkHeader(url)
	if err != nil {
		return err
	}

	return nil
}

func (g *Godl) checkHeader(url string) error {
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("error status code")
	}

	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return errors.New("does not correspond Accept-Ranges")
	}

	if resp.ContentLength < 1 {
		return errors.New("does not get content length")
	}

	g.contenttype = resp.Header.Get("Content-Type")
	g.contentlength = uint64(resp.ContentLength)

	return nil
}
