package godl

import (
	"fmt"
	"net/http"
)

type Range struct {
	low  uint64
	high uint64
}

func (g *Godl) SetRoutine(rt uint) {
	g.rt = rt
}

func (g *Godl) SetURL(url string) {
	g.url = url
}

func (g *Godl) SetOutput(output string) {
	g.output = output
}

func (g *Godl) SetRange() {
	g.ranges = make([]Range, 0)

	all := g.contentlength
	base := all / uint64(g.rt)

	l := uint64(0)
	h := base
	for {
		rg := Range{low: l, high: h}
		g.ranges = append(g.ranges, rg)

		l = h + 1
		h = h + base
		if l > all {
			break
		}
	}
}

func (g *Godl) MakeRequest() ([]*http.Request, error) {
	reqs := make([]*http.Request, 0)
	for _, r := range g.ranges {
		req, err := http.NewRequest("GET", g.url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.low, r.high))

		reqs = append(reqs, req)
	}

	return reqs, nil
}
