package gdown

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

type RangeRequest struct {
	URL      string
	FileName string
	Unit     int64
	Ranges   []*Range
}

type Range struct {
	start int64
	end   int64
}

func New(rawurl string, p int) (*RangeRequest, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	fn := getName(u.Path)

	res, err := http.Head(rawurl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return &RangeRequest{URL: rawurl, FileName: fn, Ranges: nil}, nil
	}

	cl := res.ContentLength
	unit := cl / int64(p)
	ranges := make([]*Range, p)

	for i := 0; i < p; i++ {
		start := int64(i) * unit
		end := start + unit - 1
		if i == p-1 {
			end = cl
		}

		ranges[i] = &Range{start: start, end: end}
	}

	return &RangeRequest{URL: rawurl, FileName: fn, Unit: unit, Ranges: ranges}, nil
}

func (r *RangeRequest) Run() error {
	if r.Ranges == nil {
		req, err := http.NewRequest(http.MethodGet, r.URL, nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		return save(r.FileName, res)

	} else {
		g, ctx := errgroup.WithContext(context.Background())

		for i, ran := range r.Ranges {
			i, ran := i, ran
			tmpn := fmt.Sprintf("%s-%d", r.FileName, i)

			if info, err := os.Stat(tmpn); err == nil {
				size := info.Size()
				if i == len(r.Ranges)-1 {
					if size == ran.end-ran.start {
						continue
					}
				} else if size == r.Unit {
					continue
				}
				ran.start += size
			}

			g.Go(func() error {
				req, err := makeRangeReqest(ctx, r.URL, ran.start, ran.end)
				if err != nil {
					return err
				}

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				return save(tmpn, res)
			})
		}

		if err := g.Wait(); err != nil {
			return err
		}

		return merge(r.FileName, len(r.Ranges))
	}
}

func makeRangeReqest(ctx context.Context, url string, start, end int64) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	return req, nil
}

func getName(url string) string {
	us := strings.Split(url, "/")
	return us[len(us)-1]
}

func save(fn string, res *http.Response) error {
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, res.Body); err != nil {
		return err
	}

	return nil
}

func merge(fn string, p int) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < p; i++ {
		tmpn := fmt.Sprintf("%s-%d", fn, i)
		tmp, err := os.Open(tmpn)
		if err != nil {
			return err
		}

		io.Copy(f, tmp)
		tmp.Close()

		if err := os.Remove(tmpn); err != nil {
			return err
		}
	}

	return nil
}
