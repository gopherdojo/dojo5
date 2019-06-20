package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Range struct{}

func (r *Range) Download(ctx context.Context, url string, from, to int64, outFile string) error {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", from, to))

	ch := make(chan struct{})
	errCh := make(chan error)

	go func() {
		resp, err := client.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()
		output, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			errCh <- err
		}
		defer output.Close()

		io.Copy(output, resp.Body)

		ch <- struct{}{}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ch:
		return nil
	}
}

func (r *Range) GetContentLength(ctx context.Context, url string) (int64, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)

	ch := make(chan int64)
	errCh := make(chan error)

	go func() {
		resp, err := client.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()
		if resp.Header.Get("Accept-Ranges") != "bytes" {
			errCh <- fmt.Errorf("not supported range access: %s", url)
			return
		}

		if resp.ContentLength <= 0 {
			fmt.Printf("%v", resp.ContentLength)
			errCh <- fmt.Errorf("not supported range access")
			return
		}

		ch <- resp.ContentLength
	}()

	select {
	case err := <-errCh:
		return 0, err
	case size := <-ch:
		return size, nil
	}
}
