package godl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/sync/errgroup"
)

const (
	TMPDIR  = "tmpdir"
	TMPFILE = "tmpfile"
)

type Godl struct {
	url           string
	contenttype   string
	contentlength uint64
	rt            uint
	ranges        []Range
	reqs          []*http.Request
	tmpdir        string
	tmpfiles      []string
	output        string
}

func NewGodl() *Godl {
	return &Godl{}
}

func (g *Godl) Run() error {
	if err := g.Prepare(); err != nil {
		return err
	}
	defer g.Cleanup()

	if err := g.Download(); err != nil {
		return err
	}

	if err := g.Merge(); err != nil {
		return err
	}

	return nil
}

func (g *Godl) Prepare() error {
	o, err := ParseFlag(os.Args[1:]...)
	if err != nil {
		return err
	}

	err = g.ValidateURL(o.URL)
	if err != nil {
		return err
	}
	g.SetURL(o.URL)
	g.SetRoutine(o.Rt)
	g.SetOutput(o.Output)

	g.SetRange()

	reqs, err := g.MakeRequest()
	if err != nil {
		return err
	}
	g.reqs = reqs

	g.tmpdir, err = ioutil.TempDir("", TMPDIR)
	if err != nil {
		return err
	}

	return nil
}

func (g *Godl) Download() error {
	bc := context.Background()
	ctx, cancel := context.WithCancel(bc)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	for idx, _ := range g.ranges {
		idx := idx
		eg.Go(func() error {
			return g.download(ctx, idx)
		})
	}

	return eg.Wait()
}

func (g *Godl) download(ctx context.Context, idx int) error {
	req := g.reqs[idx]

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !g.ValidateStatus(resp) {
		return errors.New("unexpected status code")
	}

	tmpfilename := filepath.Join(g.tmpdir, fmt.Sprintf("%s.%d", TMPFILE, idx))
	tmpfile, err := os.Create(tmpfilename)
	if err != nil {
		return err
	}
	defer tmpfile.Close()

	g.tmpfiles = append(g.tmpfiles, tmpfilename)

LOOP:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_, err = io.Copy(tmpfile, resp.Body)
			if err != nil && err != io.EOF {
				return err
			}
			break LOOP
		}
	}

	return nil
}

func (g *Godl) Merge() error {
	file, err := os.Create(g.output)
	if err != nil {
		return err
	}
	defer file.Close()

	sort.SliceStable(g.tmpfiles, func(i, j int) bool {
		return g.tmpfiles[i] < g.tmpfiles[j]
	})

	for _, tmpfile := range g.tmpfiles {
		combine(file, tmpfile)
	}

	return nil
}

func combine(file *os.File, tmpfile string) error {
	source, err := os.Open(tmpfile)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = io.Copy(file, source)
	if err != nil {
		return err
	}

	return nil
}

func (g *Godl) Cleanup() {
	os.RemoveAll(g.tmpdir)
}
